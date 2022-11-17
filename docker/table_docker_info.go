package docker

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDockerInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "docker_info",
		Description: "List all information from the Docker engine.",
		List: &plugin.ListConfig{
			Hydrate: listInfo,
		},
		Columns: []*plugin.Column{
			{Name: "architecture", Type: proto.ColumnType_STRING, Description: "Hardware architecture of the host, as returned by the Go runtime (GOARCH), e.g. amd64."},
			{Name: "bridge_nf_ip6_tables", Type: proto.ColumnType_BOOL, Transform: transform.FromField("BridgeNfIP6tables"), Description: "Indicates if bridge-nf-call-ip6tables is available on the host."},
			{Name: "bridge_nf_ip_tables", Type: proto.ColumnType_BOOL, Transform: transform.FromField("BridgeNfIptables"), Description: "Indicates if bridge-nf-call-iptables is available on the host."},
			{Name: "cgroup_driver", Type: proto.ColumnType_STRING, Description: "The driver to use for managing cgroups: cgroupfs, systemd, none."},
			{Name: "cgroup_version", Type: proto.ColumnType_STRING, Description: "The version of the cgroup: 1, 2"},
			// Deprecated - {Name: "cluster_advertise", Type: proto.ColumnType_STRING, Description: ""},
			// Deprecated - {Name: "cluster_store", Type: proto.ColumnType_STRING, Description: "URL of the distributed storage backend. The storage backend is used for multihost networking (to store network and endpoint information) and by the node discovery mechanism."},
			{Name: "containerd_commit", Type: proto.ColumnType_STRING, Transform: transform.FromField("ContainerdCommit.ID"), Description: "Commit holds the Git-commit (SHA1) that a binary was built from, as reported in the version-string of external tools, such as containerd, or runC."},
			{Name: "containers", Type: proto.ColumnType_INT, Description: "Total number of containers on the host."},
			{Name: "containers_paused", Type: proto.ColumnType_INT, Description: "Number of containers with status paused."},
			{Name: "containers_running", Type: proto.ColumnType_INT, Description: "Number of containers with status running."},
			{Name: "containers_stopped", Type: proto.ColumnType_INT, Description: "Number of containers with status stopped."},
			{Name: "cpu_cfs_period", Type: proto.ColumnType_BOOL, Description: "Indicates if CPU CFS(Completely Fair Scheduler) period is supported by the host."},
			{Name: "cpu_cfs_quota", Type: proto.ColumnType_BOOL, Description: "Indicates if CPU CFS(Completely Fair Scheduler) quota is supported by the host."},
			{Name: "cpu_set", Type: proto.ColumnType_BOOL, Description: "Indicates if CPUsets (cpuset.cpus, cpuset.mems) are supported by the host."},
			{Name: "cpu_shares", Type: proto.ColumnType_BOOL, Description: "Indicates if CPU Shares limiting is supported by the host."},
			{Name: "debug", Type: proto.ColumnType_BOOL, Description: "Indicates if the daemon is running in debug-mode / with debug-level logging enabled."},
			{Name: "default_address_pools", Type: proto.ColumnType_JSON, Description: "List of custom default address pools for local networks, which can be specified in the daemon.json file or dockerd option. Example: a Base 10.10.0.0/16 with Size 24 will define the set of 256 10.10.[0-255].0/24 address pools."},
			{Name: "default_runtime", Type: proto.ColumnType_STRING, Description: "Name of the default OCI runtime that is used when starting containers. The default can be overridden per-container at create time. Default: runc."},
			{Name: "docker_root_dir", Type: proto.ColumnType_STRING, Description: "Root directory of persistent Docker state. Defaults to /var/lib/docker on Linux, and C:\\ProgramData\\docker on Windows."},
			{Name: "driver", Type: proto.ColumnType_STRING, Description: "Name of the storage driver in use."},
			{Name: "driver_status", Type: proto.ColumnType_JSON, Transform: transform.FromField("DriverStatus").Transform(labelValuePairsToMap), Description: "Information specific to the storage driver, provided as label / value pairs."},
			{Name: "experimental_build", Type: proto.ColumnType_BOOL, Description: "Indicates if experimental features are enabled on the daemon."},
			{Name: "generic_resources", Type: proto.ColumnType_JSON, Description: "User-defined resources can be either Integer resources (e.g, SSD=3) or String resources (e.g, GPU=UUID1)."},
			{Name: "http_proxy", Type: proto.ColumnType_STRING, Description: "HTTP-proxy configured for the daemon. This value is obtained from the HTTP_PROXY environment variable. Credentials (user info component) in the proxy URL are masked in the API response. Containers do not automatically inherit this configuration."},
			{Name: "https_proxy", Type: proto.ColumnType_STRING, Description: "HTTPS-proxy configured for the daemon. This value is obtained from the HTTPS_PROXY environment variable. Credentials (user info component) in the proxy URL are masked in the API response. Containers do not automatically inherit this configuration."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the daemon."},
			{Name: "images", Type: proto.ColumnType_INT, Description: "Total number of images on the host. Both tagged and untagged (dangling) images are counted."},
			{Name: "index_server_address", Type: proto.ColumnType_STRING, Description: "Address / URL of the index server that is used for image search, and as a default for user authentication for Docker Hub and Docker Cloud."},
			{Name: "init_binary", Type: proto.ColumnType_STRING, Description: "Name and, optional, path of the docker-init binary. If the path is omitted, the daemon searches the host's $PATH for the binary and uses the first result."},
			{Name: "init_commit", Type: proto.ColumnType_STRING, Transform: transform.FromField("InitCommit.ID"), Description: "Commit holds the Git-commit (SHA1) that a binary was built from, as reported in the version-string of external tools, such as containerd, or runC."},
			{Name: "ipv4_forwarding", Type: proto.ColumnType_BOOL, Transform: transform.FromField("IPv4Forwarding"), Description: "Indicates IPv4 forwarding is enabled."},
			{Name: "isolation", Type: proto.ColumnType_STRING, Description: "Represents the isolation technology to use as a default for containers. The supported values are platform-specific, e.g. default, hyperv, process."},
			{Name: "kernel_memory", Type: proto.ColumnType_BOOL, Description: "Indicates if the host has kernel memory limit support enabled."},
			{Name: "kernel_version", Type: proto.ColumnType_STRING, Description: "Kernel version."},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "User-defined labels (key/value metadata) as set on the daemon."},
			{Name: "live_restore_enabled", Type: proto.ColumnType_BOOL, Description: "Indicates if live restore is enabled. If enabled, containers are kept running when the daemon is shutdown or upon daemon start if running containers are detected."},
			{Name: "logging_driver", Type: proto.ColumnType_STRING, Description: "The logging driver to use as a default for new containers."},
			{Name: "mem_total", Type: proto.ColumnType_INT, Description: "Total amount of physical memory available on the host, in bytes."},
			{Name: "memory_limit", Type: proto.ColumnType_BOOL, Description: "Indicates if the host has memory limit support enabled."},
			{Name: "n_cpu", Type: proto.ColumnType_STRING, Transform: transform.FromField("NCPU"), Description: "The number of logical CPUs usable by the daemon."},
			{Name: "n_events_listener", Type: proto.ColumnType_STRING, Description: "Number of event listeners subscribed."},
			{Name: "n_fd", Type: proto.ColumnType_INT, Description: "The total number of file Descriptors in use by the daemon process. This information is only returned if debug-mode is enabled."},
			{Name: "n_goroutines", Type: proto.ColumnType_INT, Description: "The number of goroutines that currently exist. This information is only returned if debug-mode is enabled."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Hostname of the host."},
			{Name: "no_proxy", Type: proto.ColumnType_STRING, Description: "Comma-separated list of domain extensions for which no proxy should be used. This value is obtained from the NO_PROXY environment variable. Containers do not automatically inherit this configuration."},
			{Name: "oom_kill_disable", Type: proto.ColumnType_BOOL, Description: "Indicates if OOM killer disable is supported on the host."},
			{Name: "operating_system", Type: proto.ColumnType_STRING, Description: "Name of the host's operating system, for example: 'Ubuntu 16.04.2 LTS' or 'Windows Server 2016 Datacenter'."},
			{Name: "os_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("OSType"), Description: "Generic type of the operating system of the host, as returned by the Go runtime (GOOS), e.g. linux, windows."},
			{Name: "os_version", Type: proto.ColumnType_STRING, Transform: transform.FromField("OSVersion"), Description: "Version of the host's operating system."},
			{Name: "pids_limit", Type: proto.ColumnType_BOOL, Description: "Indicates if the host kernel has PID limit support enabled."},
			{Name: "plugins", Type: proto.ColumnType_JSON, Description: "Available plugins per type."},
			{Name: "product_license", Type: proto.ColumnType_STRING, Description: "Reports a summary of the product license on the daemon. If a commercial license has been applied to the daemon, information such as number of nodes, and expiration are included."},
			{Name: "registry_config", Type: proto.ColumnType_JSON, Description: "RegistryServiceConfig stores daemon registry services configuration."},
			{Name: "runc_commit", Type: proto.ColumnType_STRING, Transform: transform.FromField("RuncCommit.ID"), Description: "Commit holds the Git-commit (SHA1) that a binary was built from, as reported in the version-string of external tools, such as containerd, or runC."},
			{Name: "runtimes", Type: proto.ColumnType_JSON, Description: "List of OCI compliant runtimes configured on the daemon. Keys hold the name used to reference the runtime."},
			{Name: "security_options", Type: proto.ColumnType_JSON, Description: "List of security features that are enabled on the daemon, such as apparmor, seccomp, SELinux, user-namespaces (userns), and rootless. Additional configuration options for each security feature may be present, and are included as a comma-separated list of key/value pairs."},
			{Name: "server_version", Type: proto.ColumnType_STRING, Description: "Version string of the daemon."},
			{Name: "swap_limit", Type: proto.ColumnType_BOOL, Description: "Indicates if the host has memory swap limit support enabled."},
			{Name: "swarm", Type: proto.ColumnType_JSON, Description: "Represents generic information about swarm."},
			{Name: "system_status", Type: proto.ColumnType_JSON, Transform: transform.FromField("SystemStatus").Transform(labelValuePairsToMap), Description: "Information specific to the storage driver, provided as label / value pairs."},
			{Name: "system_time", Type: proto.ColumnType_STRING, Description: "Current system-time in RFC 3339 format with nano-seconds."},
			{Name: "warnings", Type: proto.ColumnType_JSON, Description: "List of warnings / informational messages about missing features, or issues related to the daemon configuration. These messages can be printed by the client as information to the user."},
		},
	}
}

func labelValuePairsToMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	result := map[string]string{}
	data, ok := d.Value.([][2]string)
	if ok {
		for _, i := range data {
			result[i[0]] = i[1]
		}
	}
	return result, nil
}

func listInfo(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("docker_info.listInfo", "connection_error", err)
		return nil, err
	}
	result, err := conn.Info(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("docker_info.listInfo", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, result)
	return nil, nil
}
