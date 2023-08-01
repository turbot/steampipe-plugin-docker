# Table: docker_container

List containers from the docker engine.

## Examples

### List all containers

```sql
select
  *
from
  docker_container
```

### List running containers

```sql
select
  id,
  names
from
  docker_container
where
  state = 'running'
```

### Find a container by name

```sql
select
  *
from
  docker_container
where
  names ? '/practical_austin'
```

### List containers which do not have a health check configured

```sql
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

### List containers with host network namespace shared

```sql
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
