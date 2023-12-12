---
title: "Steampipe Table: docker_container - Query Docker Containers using SQL"
description: "Allows users to query Docker Containers, specifically the container's ID, image, command, created time, status, and more, providing insights into container configurations and status."
---

# Table: docker_container - Query Docker Containers using SQL

Docker Containers are a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another. With Docker, you can manage your infrastructure in the same ways you manage your applications. It provides a consistent and reproducible environment isolated from other applications.

## Table Usage Guide

The `docker_container` table provides insights into Docker Containers within Docker. As a DevOps engineer, explore container-specific details through this table, including the container's ID, image, command, created time, status, and more. Utilize it to uncover information about containers, such as those with specific configurations, the status of the containers, and the verification of container isolation.

## Examples

### List all containers
Explore all active containers in your Docker environment to manage and monitor your applications more effectively. This helps in identifying potential issues and understanding the overall status of your applications.

```sql+postgres
select
  *
from
  docker_container;
```

```sql+sqlite
select
  *
from
  docker_container;
```

### List running containers
Discover the segments that are actively running within your Docker environment. This can help you manage resources and troubleshoot issues more effectively.

```sql+postgres
select
  id,
  names
from
  docker_container
where
  state = 'running';
```

```sql+sqlite
select
  id,
  names
from
  docker_container
where
  state = 'running';
```

### Find a container by name
Discover the segments that correspond to a specific container name within your Docker environment. This allows you to quickly locate and analyze the details of a particular container, enhancing your overall management and oversight of your Docker resources.

```sql+postgres
select
  *
from
  docker_container
where
  names ? '/practical_austin';
```

```sql+sqlite
Error: SQLite does not support the '?' operator for JSON objects.
```

### List containers which do not have a health check configured
Identify instances where Docker containers may lack a health check configuration. This is useful to ensure all containers are functioning correctly and to maintain optimal system health.

```sql+postgres
select
  id,
  names,
  image,
  command,
  created
from
  docker_container
where
  config -> 'Healthcheck' is null;
```

```sql+sqlite
select
  id,
  names,
  image,
  command,
  created
from
  docker_container
where
  json_extract(config, '$.Healthcheck') is null;
```

### List containers with host network namespace shared
Explore which Docker containers share the host's network namespace. This is useful for understanding potential security risks, as such containers have access to all network interfaces and services running on the host machine.

```sql+postgres
select
  id,
  names,
  image,
  command,
  created
from
  docker_container
where
  inspect -> 'HostConfig' ->> 'NetworkMode' = 'host';
```

```sql+sqlite
select
  id,
  names,
  image,
  command,
  created
from
  docker_container
where
  json_extract(inspect, '$.HostConfig.NetworkMode') = 'host';
```