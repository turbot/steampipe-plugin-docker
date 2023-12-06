---
title: "Steampipe Table: docker_compose_service - Query Docker Compose Services using SQL"
description: "Allows users to query Docker Compose Services, specifically the configuration and status of services defined in a Docker Compose file."
---

# Table: docker_compose_service - Query Docker Compose Services using SQL

Docker Compose is a tool for defining and managing multi-container Docker applications. It uses YAML files to configure application services and performs the creation and start-up process of all the containers with a single command. Docker Compose Services are the different services defined in the Docker Compose file, which can include settings like build, command, image, and volumes.

## Table Usage Guide

The `docker_compose_service` table provides insights into Docker Compose Services within Docker. As a DevOps engineer, explore service-specific details through this table, including configuration, status, and associated metadata. Utilize it to uncover information about services, such as their current status, the Docker image they're using, and the commands they're running.

## Examples

### Basic info
Discover the segments that are using the maximum CPU resources in your Docker Compose service. This allows for efficient resource management and aids in identifying potential bottlenecks.

```sql+postgres
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service;
```

```sql+sqlite
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service;
```

### List CPU configurations of the services
Explore the CPU setup of your services to understand how they are configured and assess whether any adjustments are needed to optimize performance. This can provide valuable insights into potential bottlenecks and areas for improvement in your system.

```sql+postgres
select
  name,
  cpu_count,
  cpu_percent,
  cpu_period,
  cpu_quota,
  cpu_rt_period,
  cpu_rt_runtime,
  cpus,
  cpu_shares
from
  docker_compose_service;
```

```sql+sqlite
select
  name,
  cpu_count,
  cpu_percent,
  cpu_period,
  cpu_quota,
  cpu_rt_period,
  cpu_rt_runtime,
  cpus,
  cpu_shares
from
  docker_compose_service;
```

### List services running under default cgroup
Determine the areas in which Docker services are running under the default cgroup. This is useful for understanding resource allocation and identifying potential areas of optimization.

```sql+postgres
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  cgroup_parent is null;
```

```sql+sqlite
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  cgroup_parent is null;
```

### List services that do not have health check configured
Analyze the settings to understand which services are potentially vulnerable due to the absence of a configured health check. This can help in identifying areas that require immediate attention to ensure optimal system health and performance.

```sql+postgres
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  health_check is null;
```

```sql+sqlite
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  health_check is null;
```

### List services that do not have logging configured
Discover the segments that lack logging configurations to enhance system transparency and troubleshooting capabilities. This is beneficial in pinpointing areas for potential system improvement and ensuring optimal performance.

```sql+postgres
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  logging is null;
```

```sql+sqlite
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  logging is null;
```

### List services with privileged mode enabled for containers
Identify instances where services are running in privileged mode within Docker containers. This enables a comprehensive review of security practices, as running containers in privileged mode may expose them to potential risks.

```sql+postgres
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  privileged;
```

```sql+sqlite
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  privileged = 1;
```

### List services with READ ONLY mode enabled for containers
Explore services that have the READ ONLY mode enabled in their container settings. This can be useful to identify potential security measures or limitations within your Docker Compose services.

```sql+postgres
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  read_only;
```

```sql+sqlite
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  read_only;
```

### List services where user namespace is unused for containers
Discover the segments that consist of services where the user namespace remains unused for containers. This can be beneficial in identifying potential areas for optimization or troubleshooting in your Docker environment.

```sql+postgres
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  user_ns_mode is null;
```

```sql+sqlite
select
  name,
  attach,
  cgroup_parent,
  cpu_count,
  cpu_percent,
  cpus
from
  docker_compose_service
where
  user_ns_mode is null;
```