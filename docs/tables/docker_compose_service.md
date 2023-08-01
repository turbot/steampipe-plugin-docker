# Table: docker_compose_service

List all services from the Docker compose files.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

### List services which do not have health check configuration

```sql
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

### List services which do not have logging configuration

```sql
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

```sql
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

### List services with read_only mode enabled for containers

```sql
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

```sql
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
