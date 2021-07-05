package docker

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type dockerConfig struct {
	Paths []string `cty:"paths"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
}

func ConfigInstance() interface{} {
	return &dockerConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) dockerConfig {
	if connection == nil || connection.Config == nil {
		return dockerConfig{}
	}
	config, _ := connection.Config.(dockerConfig)
	return config
}
