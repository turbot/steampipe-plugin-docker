# Table: docker_compose_config

List all networks from the Docker compose files.

## Examples

### Basic info

```sql
select
  name,
  file,
  environment,
  driver,
  template_driver
from
  docker_compose_config;
```

### List configs external configuration

```sql
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

### List configs with local driver

```sql
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

```sql
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
