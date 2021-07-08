package docker

import (
	"context"
	"os"
	"time"

	"github.com/docker/docker/client"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
	if &dockerConfig != nil {
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
