package docker

import (
	"context"

	"github.com/compose-spec/compose-go/loader"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDockerComposeConfig(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_compose_config",
		Description: "List all configs from the Docker compose files.",
		List: &plugin.ListConfig{
			Hydrate: listComposeConfigs,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the config.",
			},
			{
				Name:        "file",
				Type:        proto.ColumnType_STRING,
				Description: "File to use for the config.",
			},
			{
				Name:        "environment",
				Type:        proto.ColumnType_STRING,
				Description: "Environment variable to use for the config.",
			},
			{
				Name:        "external",
				Type:        proto.ColumnType_JSON,
				Description: "External config configuration.",
			},
			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "Labels for the config.",
			},
			{
				Name:        "driver",
				Type:        proto.ColumnType_STRING,
				Description: "Driver used for the config.",
			},
			{
				Name:        "driver_opts",
				Type:        proto.ColumnType_JSON,
				Description: "Driver options for the config.",
			},
			{
				Name:        "template_driver",
				Type:        proto.ColumnType_STRING,
				Description: "Template driver used for the config.",
			},
			{
				Name:        "extensions",
				Type:        proto.ColumnType_JSON,
				Description: "Extensions for the config configuration.",
			},
		},
	}
}

func listComposeConfigs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	parsedComposeData, err := getParsedComposeData(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_config.listComposeConfigs", "parse_error", err)
		return nil, err
	}

	for _, parsedData := range parsedComposeData {
		section, ok := parsedData["configs"]
		if !ok {
			return nil, err
		}
		configs, err := loader.LoadConfigObjs(section.(map[string]interface{}))
		if err != nil {
			plugin.Logger(ctx).Error("docker_compose_service.listComposeConfigs", "load_error", err)
			return nil, err
		}

		for _, config := range configs {
			d.StreamListItem(ctx, config)
		}
	}

	return nil, nil
}
