package docker

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/compose-spec/compose-go/loader"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDockerComposeVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_compose_volume",
		Description: "List all volumes from the Docker compose files.",
		List: &plugin.ListConfig{
			Hydrate: listComposeVolumes,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the volume.",
			},
			{
				Name:        "driver",
				Type:        proto.ColumnType_STRING,
				Description: "Driver used for the volume.",
			},
			{
				Name:        "driver_opts",
				Type:        proto.ColumnType_JSON,
				Description: "Driver options for the volume.",
			},
			{
				Name:        "external",
				Type:        proto.ColumnType_JSON,
				Description: "External volume configuration.",
			},
			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "Labels for the volume.",
			},
			{
				Name:        "extensions",
				Type:        proto.ColumnType_JSON,
				Description: "Extensions for the volume configuration.",
			},
		},
	}
}

func listComposeVolumes(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	composeFilePath := "docker-compose.yml"

	// docker compose config renders the actual data model to be applied on the Docker engine. It merges the Compose files set by -f flags, resolves variables in the Compose file, and expands short-notation into the canonical format.
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "config")

	// Redirect the command output to a buffer
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Run the command
	err := cmd.Run()
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_volume.listComposeVolumes", "cmd_error", err)
		return nil, err
	}

	parsedCompose, err := loader.ParseYAML(stdout.Bytes())
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_volume.listComposeVolumes", "parse_error", err)
		return nil, err
	}

	section, ok := parsedCompose["volumes"]
	if !ok {
		return nil, err
	}
	volumes, err := loader.LoadVolumes(section.(map[string]interface{}))
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_service.listComposeVolumes", "load_error", err)
		return nil, err
	}

	for _, volume := range volumes {
		d.StreamListItem(ctx, volume)
	}
	return nil, nil
}
