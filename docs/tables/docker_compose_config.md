---
title: "Steampipe Table: docker_compose_config - Query Docker Compose Configs using SQL"
description: "Allows users to query Docker Compose Configs, specifically the configurations of Docker Compose files, providing insights into the services, networks, volumes, and secrets defined in the Compose files."
---

# Table: docker_compose_config - Query Docker Compose Configs using SQL

Docker Compose is a tool for defining and managing multi-container Docker applications. It allows users to configure application services in a YAML file, which can then be managed (start, stop, build, etc.) with a single command. Docker Compose is primarily used in the development, testing, and staging environments.

## Table Usage Guide

The `docker_compose_config` table provides insights into Docker Compose configurations. As a DevOps engineer, you can explore details of your Docker Compose files through this table, including the services, networks, volumes, and secrets defined within. Use it to uncover information about your Docker Compose configurations, such as those related to network and volume configurations, service dependencies, and secret management.

## Examples

### Basic info
Explore the fundamental details of a Docker Compose configuration to understand its environment, associated files, and the drivers in use. This can help in managing and debugging your Docker Compose setup effectively.

```sql+postgres
select
  name,
  file,
  environment,
  driver,
  template_driver
from
  docker_compose_config;
```

```sql+sqlite
select
  name,
  file,
  environment,
  driver,
  template_driver
from
  docker_compose_config;
```

### List the external configuration of configs
Explore the external settings of your configurations to understand how they interact with external systems or resources. This can be particularly useful when troubleshooting or optimizing your system's interaction with external elements.

```sql+postgres
select
  name,
  file,
  driver,
  external ->> 'Name' as external_name,
  external ->> 'External' as external,
  external -> 'Extensions' as external_extensions
from
  docker_compose_config;
```

```sql+sqlite
select
  name,
  file,
  driver,
  json_extract(external, '$.Name') as external_name,
  json_extract(external, '$.External') as external,
  external as external_extensions
from
  docker_compose_config;
```

### List configs with local driver
Explore which configurations within your Docker Compose setup are utilizing a local driver. This can be beneficial for assessing elements within your system that may be operating independently of network drivers, aiding in system management and optimization.

```sql+postgres
select
  name,
  file,
  environment,
  driver,
  template_driver
from
  docker_compose_config
where
  driver = 'local';
```

```sql+sqlite
select
  name,
  file,
  environment,
  driver,
  template_driver
from
  docker_compose_config
where
  driver = 'local';
```

### List configs without environment vars
Explore configurations that lack environment variables. This is useful for identifying potential areas in your Docker Compose setup that may need additional context or settings.

```sql+postgres
select
  name,
  file,
  driver,
  template_driver
from
  docker_compose_config
where
  environment is null;
```

```sql+sqlite
select
  name,
  file,
  driver,
  template_driver
from
  docker_compose_config
where
  environment is null;
```