package docker

import (
	"context"

	"github.com/docker/docker/api/types"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableDockerNetwork(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_network",
		Description: "List all networks from the Docker engine.",
		List: &plugin.ListConfig{
			Hydrate: listNetwork,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the network."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the network."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(timeToTimestamp), Description: "Time when the network was created."},
			{Name: "scope", Type: proto.ColumnType_STRING, Description: "Scope describes the level at which the network exists (e.g. 'swarm' for cluster-wide or 'local' for machine level)."},
			{Name: "driver", Type: proto.ColumnType_STRING, Description: "Driver is the Driver name used to create the network (e.g. 'bridge', 'overlay')."},
			{Name: "enable_ipv6", Type: proto.ColumnType_BOOL, Transform: transform.FromField("EnableIPv6"), Description: "True if IPv6 is enabled."},
			{Name: "ipam", Type: proto.ColumnType_JSON, Transform: transform.FromField("IPAM"), Description: "Network IP address management."},
			{Name: "internal", Type: proto.ColumnType_BOOL, Description: "True if if the network is used internal only."},
			{Name: "attachable", Type: proto.ColumnType_BOOL, Description: "True if if the global scope is manually attachable by regular containers from workers in swarm mode."},
			{Name: "ingress", Type: proto.ColumnType_BOOL, Description: "True if the network is providing the routing-mesh for the swarm cluster."},
			{Name: "config_from", Type: proto.ColumnType_JSON, Description: "Source which will provide the configuration for this network."},
			{Name: "config_only", Type: proto.ColumnType_BOOL, Description: "True if the network is a place-holder for configuration of other networks only."},
			{Name: "containers", Type: proto.ColumnType_JSON, Description: "Endpoints belonging to the network."},
			{Name: "options", Type: proto.ColumnType_JSON, Description: "Network specific options to use for when creating the network."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Metadata specific to the network."},
			{Name: "peers", Type: proto.ColumnType_JSON, Description: "Peer nodes for an overlay network."},
			{Name: "services", Type: proto.ColumnType_JSON, Description: "Services in the network."},
		},
	}
}

func listNetwork(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_network.listNetwork", "connection_error", err)
		return nil, err
	}
	params := types.NetworkListOptions{}
	items, err := conn.NetworkList(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("docker_network.listNetwork", "query_error", err, "params", params)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
