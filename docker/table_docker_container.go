package docker

import (
	"context"

	"github.com/docker/docker/api/types"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDockerContainer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_container",
		Description: "List all containers from the Docker engine.",
		List: &plugin.ListConfig{
			Hydrate: listContainer,
			//KeyColumns:    plugin.OptionalColumns([]string{"path"}),
		},
		//GetMatrixItem: dockerfileList,
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the container."},
			{Name: "names", Type: proto.ColumnType_JSON, Description: "Names assigned to the container."},
			// Other columns
			{Name: "image", Type: proto.ColumnType_STRING, Description: "Name of the image for the container."},
			{Name: "image_id", Type: proto.ColumnType_STRING, Description: "ID of the image for the container."},
			{Name: "command", Type: proto.ColumnType_STRING, Description: "Main command running in the container."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time when the container was created."},
			{Name: "ports", Type: proto.ColumnType_JSON, Description: "Ports open for the container."},
			{Name: "size_rw", Type: proto.ColumnType_INT, Description: ""},
			{Name: "size_root_fs", Type: proto.ColumnType_INT, Description: ""},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels for the container."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the container: running, restarting, etc."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "Status message from the container."},
			{Name: "host_config", Type: proto.ColumnType_JSON, Description: "Host configuration for the container."},
			{Name: "network_settings", Type: proto.ColumnType_JSON, Description: "Network settings for the container."},
			{Name: "mounts", Type: proto.ColumnType_JSON, Description: "Volume mounts for the container."},
		},
	}
}

func listContainer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_container.listContainer", "connection_error", err)
		return nil, err
	}
	params := types.ContainerListOptions{
		All:  true,
		Size: true,
	}
	items, err := conn.ContainerList(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("docker_container.listContainer", "query_error", err, "params", params)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
