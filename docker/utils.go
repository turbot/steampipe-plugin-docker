package docker

import (
	"bytes"
	"context"
	"os/exec"
	"path/filepath"
	"strings"
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

	clientOpts := []client.Opt{}

	if dockerConfig.Host != nil {
		clientOpts = append(clientOpts, client.WithHost(*dockerConfig.Host))
	}

	if dockerConfig.APIVersion != nil {
		clientOpts = append(clientOpts, client.WithVersion(*dockerConfig.APIVersion))
	}

	if dockerConfig.CertPath != nil {
		clientOpts = append(clientOpts, client.WithTLSClientConfig(
			filepath.Join(*dockerConfig.CertPath, "ca.pem"),
			filepath.Join(*dockerConfig.CertPath, "cert.pem"),
			filepath.Join(*dockerConfig.CertPath, "key.pem"),
		))
	}

	conn, err := client.NewClientWithOpts(clientOpts...)
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

	if dockerConfig.DockerComposeFilePaths == nil || len(dockerConfig.DockerComposeFilePaths) == 0 {
		return nil, nil
	} else {
		for _, i := range dockerConfig.DockerComposeFilePaths {

			// List the files in the given source directory
			files, err := d.GetSourceFiles(i)
			if err != nil {
				plugin.Logger(ctx).Error("getParsedComposeDataUncached.DockerComposeFilePaths", "get_source_files_error", err)

				// If the specified path is unavailable, then an empty row should populate
				if strings.Contains(err.Error(), "failed to get directory specified by the source") {
					return nil, nil
				}
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
