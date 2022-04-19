package docker

import (
	"context"

	"github.com/docker/docker/api/types/filters"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableDockerVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_volume",
		Description: "List all volumes from the Docker engine.",
		List: &plugin.ListConfig{
			Hydrate: listVolume,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the volume."},
			{Name: "driver", Type: proto.ColumnType_STRING, Description: "Name of the volume driver used by the volume."},
			// Other columns
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt"), Description: "Time when the volume was created."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels for the volume."},
			{Name: "mountpoint", Type: proto.ColumnType_STRING, Description: "Mount path of the volume on the host."},
			{Name: "options", Type: proto.ColumnType_JSON, Description: "The driver specific options used when creating the volume."},
			{Name: "scope", Type: proto.ColumnType_STRING, Description: "The level at which the volume exists. Either global for cluster-wide, or local for machine level."},
			{Name: "status", Type: proto.ColumnType_JSON, Description: "Low-level details about the volume, provided by the volume driver."},
			{Name: "usage_data", Type: proto.ColumnType_JSON, Description: "Usage data for the volume."},
		},
	}
}

func listVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_volume.listVolume", "connection_error", err)
		return nil, err
	}
	params := filters.Args{}
	result, err := conn.VolumeList(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("docker_volume.listVolume", "query_error", err, "params", params)
		return nil, err
	}
	for _, i := range result.Volumes {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
