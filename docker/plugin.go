package docker

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-docker",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"docker_dockerfile_instruction": tableDockerDockerfileInstruction(ctx),
			"docker_container":              tableDockerContainer(ctx),
			"docker_image":                  tableDockerImage(ctx),
			"docker_info":                   tableDockerInfo(ctx),
			"docker_network":                tableDockerNetwork(ctx),
			"docker_volume":                 tableDockerVolume(ctx),
		},
	}
	return p
}
