package docker

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/compose-spec/compose-go/loader"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDockerComposeSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_compose_secret",
		Description: "List all secrets from the Docker compose files.",
		List: &plugin.ListConfig{
			Hydrate: listComposeSecrets,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the secret.",
			},
			{
				Name:        "file",
				Type:        proto.ColumnType_STRING,
				Description: "File to use for the secret.",
			},
			{
				Name:        "environment",
				Type:        proto.ColumnType_STRING,
				Description: "Environment variable to use for the secret.",
			},
			{
				Name:        "external",
				Type:        proto.ColumnType_JSON,
				Description: "External secret configuration.",
			},
			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "Labels for the secret.",
			},
			{
				Name:        "driver",
				Type:        proto.ColumnType_STRING,
				Description: "Driver used for the secret.",
			},
			{
				Name:        "driver_opts",
				Type:        proto.ColumnType_JSON,
				Description: "Driver options for the secret.",
			},
			{
				Name:        "template_driver",
				Type:        proto.ColumnType_STRING,
				Description: "Template driver used for the secret.",
			},
			{
				Name:        "extensions",
				Type:        proto.ColumnType_JSON,
				Description: "Extensions for the secret configuration.",
			},
		},
	}
}

func listComposeSecrets(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	composeFilePath := "docker-compose.yml"

	// docker compose config renders the actual data model to be applied on the Docker engine. It merges the Compose files set by -f flags, resolves variables in the Compose file, and expands short-notation into the canonical format.
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "config")

	// Redirect the command output to a buffer
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Run the command
	err := cmd.Run()
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_secret.listComposeSecrets", "cmd_error", err)
		return nil, err
	}

	parsedCompose, err := loader.ParseYAML(stdout.Bytes())
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_secret.listComposeSecrets", "parse_error", err)
		return nil, err
	}

	section, ok := parsedCompose["secrets"]
	if !ok {
		return nil, err
	}
	secrets, err := loader.LoadSecrets(section.(map[string]interface{}))
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_service.listComposeSecrets", "load_error", err)
		return nil, err
	}

	for _, secret := range secrets {
		d.StreamListItem(ctx, secret)
	}
	return nil, nil
}
