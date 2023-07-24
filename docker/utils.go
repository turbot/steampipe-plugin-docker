package docker

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"time"

	"github.com/compose-spec/compose-go/loader"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(_ context.Context, d *plugin.QueryData) (*client.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "docker"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*client.Client), nil
	}

	// Docker configuration is always loaded from environment variables. If a setting
	// is given in the steampipe configuration file, then set the env locally for the
	// process.
	dockerConfig := GetConfig(d.Connection)
	if dockerConfig.Host != nil {
		os.Setenv("DOCKER_HOST", *dockerConfig.Host)
	}
	if dockerConfig.APIVersion != nil {
		os.Setenv("DOCKER_API_VERSION", *dockerConfig.APIVersion)
	}
	if dockerConfig.CertPath != nil {
		os.Setenv("DOCKER_CERT_PATH", *dockerConfig.CertPath)
	}
	if dockerConfig.TLSVerify != nil {
		if *dockerConfig.TLSVerify {
			os.Setenv("DOCKER_TLS_VERIFY", "1")
		} else {
			os.Setenv("DOCKER_TLS_VERIFY", "0")
		}
	}

	// Always load the docker config from ENV vars
	conn, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func timeToTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	var ts time.Time
	switch t := d.Value.(type) {
	case time.Time:
		ts = t
	case *time.Time:
		ts = *t
	}
	tsBytes, err := ts.MarshalText()
	if err != nil {
		return nil, nil
	}
	return string(tsBytes), nil
}

func getParsedComposeData(ctx context.Context, d *plugin.QueryData) ([]map[string]interface{}, error) {
	conn, err := getParsedComposeDataCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.([]map[string]interface{}), nil
}

var getParsedComposeDataCached = plugin.HydrateFunc(getParsedComposeDataUncached).Memoize()

func getParsedComposeDataUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	dockerConfig := GetConfig(d.Connection)

	// Gather file path matches for the glob
	var matches []string

	if dockerConfig.DockerComposeFilePaths == nil {

		// add the default available files in CWD
		defaultFiles := []string{"compose.yaml", "compose.yml", "docker-compose.yml", "docker-compose.yaml"}
		for _, file := range defaultFiles {
			_, err := os.Stat(file)
			if err == nil {
				matches = append(matches, file)
			}
		}
	} else {
		for _, i := range dockerConfig.DockerComposeFilePaths {

			// List the files in the given source directory
			files, err := d.GetSourceFiles(i)
			if err != nil {
				return nil, err
			}
			plugin.Logger(ctx).Warn("getParsedComposeDataUncached", "source", i, "files", files)
			matches = append(matches, files...)
		}
	}

	if len(matches) == 0 {
		return nil, errors.New("docker_compose_file_paths must be configured")
	}

	var parsedComposeContent []map[string]interface{}

	// fetch compose data from the files
	for _, composeFilePath := range matches {

		// docker compose config renders the actual data model to be applied on the Docker engine. It merges the Compose files set by -f flags, resolves variables in the Compose file, and expands short-notation into the canonical format.
		cmd := exec.Command("docker-compose", "-f", composeFilePath, "config")

		// Redirect the command output to a buffer
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		// Run the command
		err := cmd.Run()
		if err != nil {
			plugin.Logger(ctx).Error("getParsedComposeDataUncached", "cmd_error", err)
			return nil, err
		}

		parsedCompose, err := loader.ParseYAML(stdout.Bytes())
		if err != nil {
			plugin.Logger(ctx).Error("getParsedComposeDataUncached", "parse_error", err)
			return nil, err
		}
		parsedComposeContent = append(parsedComposeContent, parsedCompose)
	}

	return parsedComposeContent, nil
}
