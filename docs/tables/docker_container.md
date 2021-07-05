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
