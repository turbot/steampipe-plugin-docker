package docker

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type dockerConfig struct {
	Paths                  []string `hcl:"paths,optional" steampipe:"watch"`
	DockerfilePaths        []string `hcl:"dockerfile_paths,optional" steampipe:"watch"`
	DockerComposeFilePaths []string `hcl:"docker_compose_file_paths,optional" steampipe:"watch"`
	Host                   *string  `hcl:"host"`
	APIVersion             *string  `hcl:"api_version"`
	CertPath               *string  `hcl:"cert_path"`
	TLSVerify              *bool    `hcl:"tls_verify"`
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
