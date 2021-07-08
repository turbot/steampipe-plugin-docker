package docker

import (
	"context"

	"github.com/docker/docker/api/types"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDockerImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_image",
		Description: "List all images from the Docker engine.",
		List: &plugin.ListConfig{
			Hydrate: listImage,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the image."},
			// Other columns
			// TODO what is this, it's always -1? {Name: "containers", Type: proto.ColumnType_INT, Description: ""},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixToTimestamp), Description: "Time when the image was created."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels for the image."},
			{Name: "parent_id", Type: proto.ColumnType_STRING, Description: "Parent ID of the image."},
			{Name: "repo_digests", Type: proto.ColumnType_JSON, Description: "Repository digests for the image."},
			{Name: "repo_tags", Type: proto.ColumnType_JSON, Description: "Repository tags for the image."},
			// TODO what is this, it's always -1? {Name: "shared_size", Type: proto.ColumnType_INT, Description: "Shared size of the image in bytes."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size of the image in bytes."},
			{Name: "virtual_size", Type: proto.ColumnType_INT, Description: "Virtual size of the image in bytes."},
		},
	}
}

func listImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_image.listImage", "connection_error", err)
		return nil, err
	}
	params := types.ImageListOptions{}
	items, err := conn.ImageList(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("docker_image.listImage", "query_error", err, "params", params)
		return nil, err
	}
	for _, i := range items {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
