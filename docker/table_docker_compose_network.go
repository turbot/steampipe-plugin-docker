package docker

import (
	"context"

	"github.com/compose-spec/compose-go/loader"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDockerComposeNetwork(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_compose_network",
		Description: "List all networks from the Docker compose files.",
		List: &plugin.ListConfig{
			Hydrate: listComposeNetworks,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the network.",
			},
			{
				Name:        "driver",
				Type:        proto.ColumnType_STRING,
				Description: "Driver used for the network.",
			},
			{
				Name:        "driver_opts",
				Type:        proto.ColumnType_JSON,
				Description: "Driver options for the network.",
			},
			{
				Name:        "ipam",
				Type:        proto.ColumnType_JSON,
				Description: "IPAM (IP Address Management) configuration for the network.",
			},
			{
				Name:        "external",
				Type:        proto.ColumnType_JSON,
				Description: "External network configuration.",
			},
			{
				Name:        "internal",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if the network is internal.",
			},
			{
				Name:        "attachable",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if containers can be attached to the network.",
			},
			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "Labels for the network.",
			},
			{
				Name:        "enable_ipv6",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if IPv6 is enabled for the network.",
				Transform:   transform.FromField("EnableIPv6"),
			},
			{
				Name:        "extensions",
				Type:        proto.ColumnType_JSON,
				Description: "Extensions for the network configuration.",
			},
		},
	}
}

func listComposeNetworks(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	parsedComposeData, err := getParsedComposeData(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_network.listComposeNetworks", "parse_error", err)
		return nil, err
	}

	for _, parsedData := range parsedComposeData {
		section, ok := parsedData["networks"]
		if !ok {
			return nil, err
		}
		networks, err := loader.LoadNetworks(section.(map[string]interface{}))
		if err != nil {
			plugin.Logger(ctx).Error("docker_compose_service.listComposeServices", "load_error", err)
			return nil, err
		}

		for _, network := range networks {
			d.StreamListItem(ctx, network)
		}
	}
	return nil, nil
}
