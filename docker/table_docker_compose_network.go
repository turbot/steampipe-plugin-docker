package docker

import (
	"bytes"
	"context"
	"os/exec"

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
	composeFilePath := "docker-compose.yml"

	// docker compose config renders the actual data model to be applied on the Docker engine. It merges the Compose files set by -f flags, resolves variables in the Compose file, and expands short-notation into the canonical format.
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "config")

	// Redirect the command output to a buffer
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Run the command
	err := cmd.Run()
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_network.listComposeNetworks", "cmd_error", err)
		return nil, err
	}

	parsedCompose, err := loader.ParseYAML(stdout.Bytes())
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_network.listComposeNetworks", "parse_error", err)
		return nil, err
	}

	// configFile := types.ConfigFile{}
	// configFile.Config = parsedCompose
	// configDetails := types.ConfigDetails{
	// 	ConfigFiles: []types.ConfigFile{configFile},
	// }

	// project, err := loader.Load(configDetails)
	// if err != nil {
	// 	plugin.Logger(ctx).Error("docker_compose_network.listComposeNetworks", "load_error", err)
	// 	return nil, err
	// }

	section, ok := parsedCompose["networks"]
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
	return nil, nil
}
