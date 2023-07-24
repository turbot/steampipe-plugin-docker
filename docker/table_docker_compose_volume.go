package docker

import (
	"context"

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
	parsedComposeData, err := getParsedComposeData(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_volume.listComposeVolumes", "parse_error", err)
		return nil, err
	}

	for _, parsedData := range parsedComposeData {
		section, ok := parsedData["volumes"]
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
	}
	return nil, nil
}
