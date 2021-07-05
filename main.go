package main

import (
	"github.com/turbot/steampipe-plugin-docker/docker"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: docker.Plugin})
}
