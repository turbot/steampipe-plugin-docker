package docker

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type dockerConfig struct {
	Paths      []string `cty:"paths" steampipe:"watch"`
	Host       *string  `cty:"host"`
	APIVersion *string  `cty:"api_version"`
	CertPath   *string  `cty:"cert_path"`
	TLSVerify  *bool    `cty:"tls_verify"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"paths": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"host": {
		Type: schema.TypeString,
	},
	"api_version": {
		Type: schema.TypeString,
	},
	"cert_path": {
		Type: schema.TypeString,
	},
	"tls_verify": {
		Type: schema.TypeBool,
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
