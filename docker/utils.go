package docker

import (
	"context"

	"github.com/docker/docker/client"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*client.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "docker"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*client.Client), nil
	}

	/*
		// Default to using env vars
		apiKey := os.Getenv("STRIPE_API_KEY")

		// But prefer the config
		dockerConfig := GetConfig(d.Connection)
		if &dockerConfig != nil {
			if dockerConfig.APIKey != nil {
				apiKey = *dockerConfig.APIKey
			}
		}

		if apiKey == "" {
			// Credentials not set
			return nil, errors.New("api_key must be configured")
		}

		config := &docker.BackendConfig{
			MaxNetworkRetries: 10,
		}
	*/

	conn, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}
