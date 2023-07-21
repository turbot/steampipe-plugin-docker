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

func tableDockerComposeService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_compose_service",
		Description: "List all services from the Docker compose files.",
		List: &plugin.ListConfig{
			Hydrate: listComposeServices,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the service.",
			},
			{
				Name:        "profiles",
				Type:        proto.ColumnType_JSON,
				Description: "List of profiles associated with the service.",
			},
			{
				Name:        "annotations",
				Type:        proto.ColumnType_JSON,
				Description: "Annotations for the service.",
			},
			{
				Name:        "attach",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if containers should be attached to the terminal.",
			},
			{
				Name:        "build",
				Type:        proto.ColumnType_JSON,
				Description: "Build configuration for the service.",
			},
			{
				Name:        "blkio_config",
				Type:        proto.ColumnType_JSON,
				Description: "Block I/O (BLKIO) configuration for the service.",
			},
			{
				Name:        "cap_add",
				Type:        proto.ColumnType_JSON,
				Description: "List of capabilities to add to the container.",
			},
			{
				Name:        "cap_drop",
				Type:        proto.ColumnType_JSON,
				Description: "List of capabilities to drop from the container.",
			},
			{
				Name:        "cgroup_parent",
				Type:        proto.ColumnType_STRING,
				Description: "Parent cgroup for the container.",
			},
			{
				Name:        "cgroup",
				Type:        proto.ColumnType_STRING,
				Description: "Cgroup rule to apply to the container.",
			},
			{
				Name:        "cpu_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of CPUs to allocate to the container.",
				Transform:   transform.FromField("CPUCount"),
			},
			{
				Name:        "cpu_percent",
				Type:        proto.ColumnType_DOUBLE,
				Description: "CPU utilization limit as a percentage.",
				Transform:   transform.FromField("CPUPercent"),
			},
			{
				Name:        "cpu_period",
				Type:        proto.ColumnType_INT,
				Description: "CPU CFS (Completely Fair Scheduler) period.",
				Transform:   transform.FromField("CPUPeriod"),
			},
			{
				Name:        "cpu_quota",
				Type:        proto.ColumnType_INT,
				Description: "CPU CFS (Completely Fair Scheduler) quota.",
				Transform:   transform.FromField("CPUQuota"),
			},
			{
				Name:        "cpu_rt_period",
				Type:        proto.ColumnType_INT,
				Description: "CPU real-time period.",
				Transform:   transform.FromField("CPURTPeriod"),
			},
			{
				Name:        "cpu_rt_runtime",
				Type:        proto.ColumnType_INT,
				Description: "CPU real-time runtime.",
				Transform:   transform.FromField("CPURTRuntime"),
			},
			{
				Name:        "cpus",
				Type:        proto.ColumnType_DOUBLE,
				Description: "Number of CPUs to allocate to the container (in fractional form).",
				Transform:   transform.FromField("CPUS"),
			},
			{
				Name:        "cpuset",
				Type:        proto.ColumnType_STRING,
				Description: "CPUs in which to allow execution (comma-separated list or ranges).",
				Transform:   transform.FromField("CPUSet"),
			},
			{
				Name:        "cpu_shares",
				Type:        proto.ColumnType_INT,
				Description: "CPU shares (relative weight) for the container.",
				Transform:   transform.FromField("CPUShares"),
			},
			{
				Name:        "command",
				Type:        proto.ColumnType_STRING,
				Description: "Command for the service containers.",
			},
			{
				Name:        "configs",
				Type:        proto.ColumnType_JSON,
				Description: "List of configurations for the service.",
			},
			{
				Name:        "container_name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the container.",
			},
			{
				Name:        "credential_spec",
				Type:        proto.ColumnType_JSON,
				Description: "Credential specification for the container.",
			},
			{
				Name:        "depends_on",
				Type:        proto.ColumnType_JSON,
				Description: "Dependencies for the service.",
			},
			{
				Name:        "deploy",
				Type:        proto.ColumnType_JSON,
				Description: "Deployment configuration for the service.",
			},
			{
				Name:        "device_cgroup_rules",
				Type:        proto.ColumnType_JSON,
				Description: "List of device cgroup rules for the container.",
			},
			{
				Name:        "devices",
				Type:        proto.ColumnType_JSON,
				Description: "List of devices to add to the container.",
			},
			{
				Name:        "dns",
				Type:        proto.ColumnType_JSON,
				Description: "List of DNS servers for the container.",
				Transform:   transform.FromField("DNS"),
			},
			{
				Name:        "dns_opt",
				Type:        proto.ColumnType_JSON,
				Description: "List of DNS options.",
				Transform:   transform.FromField("DNSOpts"),
			},
			{
				Name:        "dns_search",
				Type:        proto.ColumnType_JSON,
				Description: "List of DNS search domains.",
				Transform:   transform.FromField("DNSSearch"),
			},
			{
				Name:        "dockerfile",
				Type:        proto.ColumnType_STRING,
				Description: "Path to the Dockerfile to use for the container.",
			},
			{
				Name:        "domain_name",
				Type:        proto.ColumnType_STRING,
				Description: "Domain name of the container.",
			},
			{
				Name:        "entrypoint",
				Type:        proto.ColumnType_JSON,
				Description: "Entrypoint for the service containers.",
			},
			{
				Name:        "environment",
				Type:        proto.ColumnType_JSON,
				Description: "Environment variables for the container.",
			},
			{
				Name:        "env_file",
				Type:        proto.ColumnType_JSON,
				Description: "List of environment files to read.",
			},
			{
				Name:        "expose",
				Type:        proto.ColumnType_JSON,
				Description: "List of ports to expose from the container.",
			},
			{
				Name:        "extends",
				Type:        proto.ColumnType_JSON,
				Description: "Configuration that the service extends.",
			},
			{
				Name:        "external_links",
				Type:        proto.ColumnType_JSON,
				Description: "List of external links to other services.",
			},
			{
				Name:        "extra_hosts",
				Type:        proto.ColumnType_JSON,
				Description: "Additional hostnames to resolve inside the container.",
			},
			{
				Name:        "group_add",
				Type:        proto.ColumnType_JSON,
				Description: "List of additional groups for the container.",
			},
			{
				Name:        "hostname",
				Type:        proto.ColumnType_STRING,
				Description: "Hostname of the container.",
			},
			{
				Name:        "health_check",
				Type:        proto.ColumnType_JSON,
				Description: "Health check configuration for the service.",
			},
			{
				Name:        "image",
				Type:        proto.ColumnType_STRING,
				Description: "Docker image for the container.",
			},
			{
				Name:        "init",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if the container should run as an init process.",
			},
			{
				Name:        "ipc",
				Type:        proto.ColumnType_STRING,
				Description: "IPC (Inter-Process Communication) mode for the container.",
			},
			{
				Name:        "isolation",
				Type:        proto.ColumnType_STRING,
				Description: "Isolation technology used for the container.",
			},
			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "Labels for the service.",
			},
			{
				Name:        "custom_labels",
				Type:        proto.ColumnType_JSON,
				Description: "Custom labels for the service.",
			},
			{
				Name:        "links",
				Type:        proto.ColumnType_JSON,
				Description: "List of links to other services.",
			},
			{
				Name:        "logging",
				Type:        proto.ColumnType_JSON,
				Description: "Logging configuration for the service.",
			},
			{
				Name:        "log_driver",
				Type:        proto.ColumnType_STRING,
				Description: "Logging driver for the container.",
			},
			{
				Name:        "log_opt",
				Type:        proto.ColumnType_JSON,
				Description: "Options for the logging driver.",
			},
			{
				Name:        "mem_limit",
				Type:        proto.ColumnType_INT,
				Description: "Memory limit for the container.",
			},
			{
				Name:        "mem_reservation",
				Type:        proto.ColumnType_INT,
				Description: "Memory reservation for the container.",
			},
			{
				Name:        "memswap_limit",
				Type:        proto.ColumnType_INT,
				Description: "Swap limit for the container.",
			},
			{
				Name:        "mem_swappiness",
				Type:        proto.ColumnType_INT,
				Description: "Swappiness value for the container.",
			},
			{
				Name:        "mac_address",
				Type:        proto.ColumnType_STRING,
				Description: "MAC address for the container.",
			},
			{
				Name:        "net",
				Type:        proto.ColumnType_STRING,
				Description: "Network mode for the container.",
			},
			{
				Name:        "network_mode",
				Type:        proto.ColumnType_STRING,
				Description: "Network mode for the service.",
			},
			{
				Name:        "networks",
				Type:        proto.ColumnType_JSON,
				Description: "Network configurations for the service.",
			},
			{
				Name:        "oom_kill_disable",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if OOM (Out-Of-Memory) killer is disabled.",
			},
			{
				Name:        "oom_score_adj",
				Type:        proto.ColumnType_INT,
				Description: "OOM score adjustment for the container.",
			},
			{
				Name:        "pid",
				Type:        proto.ColumnType_STRING,
				Description: "PID (Process ID) namespace for the container.",
			},
			{
				Name:        "pids_limit",
				Type:        proto.ColumnType_INT,
				Description: "PIDS limit for the container.",
			},
			{
				Name:        "platform",
				Type:        proto.ColumnType_STRING,
				Description: "Platform to use for the container.",
			},
			{
				Name:        "ports",
				Type:        proto.ColumnType_JSON,
				Description: "Ports to publish from the container.",
			},
			{
				Name:        "privileged",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if the container should run in privileged mode.",
			},
			{
				Name:        "pull_policy",
				Type:        proto.ColumnType_STRING,
				Description: "Pull policy for the container image.",
			},
			{
				Name:        "read_only",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if the container's root filesystem should be read-only.",
			},
			{
				Name:        "restart",
				Type:        proto.ColumnType_STRING,
				Description: "Restart policy for the service.",
			},
			{
				Name:        "runtime",
				Type:        proto.ColumnType_STRING,
				Description: "Runtime to use for the container.",
			},
			{
				Name:        "scale",
				Type:        proto.ColumnType_INT,
				Description: "Number of replicas to run for the service.",
			},
			{
				Name:        "secrets",
				Type:        proto.ColumnType_JSON,
				Description: "Secrets configuration for the service.",
			},
			{
				Name:        "security_opt",
				Type:        proto.ColumnType_JSON,
				Description: "List of security options for the container.",
			},
			{
				Name:        "shm_size",
				Type:        proto.ColumnType_INT,
				Description: "Size of /dev/shm for the container.",
			},
			{
				Name:        "stdin_open",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if stdin should be kept open for the container.",
			},
			{
				Name:        "stop_grace_period",
				Type:        proto.ColumnType_STRING,
				Description: "Grace period for the container to stop.",
			},
			{
				Name:        "stop_signal",
				Type:        proto.ColumnType_STRING,
				Description: "Signal to stop the container.",
			},
			{
				Name:        "sysctls",
				Type:        proto.ColumnType_JSON,
				Description: "Sysctls configuration for the container.",
			},
			{
				Name:        "tmpfs",
				Type:        proto.ColumnType_JSON,
				Description: "List of tmpfs mounts for the container.",
			},
			{
				Name:        "tty",
				Type:        proto.ColumnType_BOOL,
				Description: "Specifies if the container should allocate a pseudo-TTY.",
			},
			{
				Name:        "ulimits",
				Type:        proto.ColumnType_JSON,
				Description: "ULimits (resource limits) for the container.",
			},
			{
				Name:        "user",
				Type:        proto.ColumnType_STRING,
				Description: "User to run commands inside the container.",
			},
			{
				Name:        "user_ns_mode",
				Type:        proto.ColumnType_STRING,
				Description: "User namespace mode for the container.",
				Transform:   transform.FromField("UserNSMode"),
			},
			{
				Name:        "uts",
				Type:        proto.ColumnType_STRING,
				Description: "UTS namespace for the container.",
			},
			{
				Name:        "volume_driver",
				Type:        proto.ColumnType_STRING,
				Description: "Volume driver to use for the container.",
			},
			{
				Name:        "volumes",
				Type:        proto.ColumnType_JSON,
				Description: "Volumes to mount in the container.",
			},
			{
				Name:        "volumes_from",
				Type:        proto.ColumnType_JSON,
				Description: "List of volumes to mount from other containers.",
			},
			{
				Name:        "working_dir",
				Type:        proto.ColumnType_STRING,
				Description: "Working directory for commands inside the container.",
			},
			{
				Name:        "extensions",
				Type:        proto.ColumnType_JSON,
				Description: "Extensions for the service configuration.",
			},
		},
	}
}

func listComposeServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	composeFilePath := "docker-compose.yml"

	// docker compose config renders the actual data model to be applied on the Docker engine. It merges the Compose files set by -f flags, resolves variables in the Compose file, and expands short-notation into the canonical format.
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "config")

	// Redirect the command output to a buffer
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// Run the command
	err := cmd.Run()
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_service.listComposeServices", "cmd_error", err)
		return nil, err
	}

	parsedCompose, err := loader.ParseYAML(stdout.Bytes())
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_service.listComposeServices", "parse_error", err)
		return nil, err
	}

	// configFile := types.ConfigFile{}
	// configFile.Config = parsedCompose
	// configDetails := types.ConfigDetails{
	// 	ConfigFiles: []types.ConfigFile{configFile},
	// }

	// project, err := loader.Load(configDetails)
	// if err != nil {
	// 	plugin.Logger(ctx).Error("docker_compose_service.listComposeServices", "load_error", err)
	// 	return nil, err
	// }

	section, ok := parsedCompose["services"]
	if !ok {
		return nil, err
	}
	services, err := loader.LoadServices("", section.(map[string]interface{}), "", nil, nil)
	if err != nil {
		plugin.Logger(ctx).Error("docker_compose_service.listComposeServices", "load_error", err)
		return nil, err
	}

	for _, service := range services {
		plugin.Logger(ctx).Error("docker_compose_service.service", service.Name)
		d.StreamListItem(ctx, service)
	}
	return nil, nil
}
