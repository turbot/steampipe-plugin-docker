# Table: docker_volume

List volumes from the docker engine.

## Examples

### List all volumes

```sql
select
  *
from
  docker_volume
```

### Find all volumes with a given label

```sql
select
  *
from
  docker_volume
where
  labels ->> 'com.docker.compose.volume' = 'postgres'
```
