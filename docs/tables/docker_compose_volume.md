# Table: docker_compose_volume

List all volumes from the Docker compose files.

## Examples

### Basic info

```sql
select
  name,
  driver,
  jsonb_pretty(driver_opts) as driver_opts,
  jsonb_pretty(external) as external
from
  docker_compose_volume;
```

### List volumes external configuration

```sql
select
  name,
  driver,
  external ->> 'Name' as external_name,
  external ->> 'External' as external,
  external -> 'Extensions' as external_extensions
from
  docker_compose_volume;
```

### List volumes with local driver

```sql
select
  name,
  driver,
  jsonb_pretty(driver_opts) as driver_opts,
  jsonb_pretty(external) as external
from
  docker_compose_volume
where
  driver = 'local';
```
