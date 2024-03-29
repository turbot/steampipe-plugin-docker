package docker

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-docker",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"docker_compose_config":  tableDockerComposeConfig(ctx),
			"docker_compose_network": tableDockerComposeNetwork(ctx),
			"docker_compose_secret":  tableDockerComposeSecret(ctx),
			"docker_compose_service": tableDockerComposeService(ctx),
			"docker_compose_volume":  tableDockerComposeVolume(ctx),
			"docker_container":       tableDockerContainer(ctx),
			"docker_image":           tableDockerImage(ctx),
			"docker_info":            tableDockerInfo(ctx),
			"docker_network":         tableDockerNetwork(ctx),
			"docker_volume":          tableDockerVolume(ctx),
			"dockerfile_instruction": tableDockerfileInstruction(ctx),
		},
	}
	return p
}
