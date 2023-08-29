# Table: docker_compose_secret

List all secrets from the Docker compose files.

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
  docker_compose_secret;
```

### List the external configuration of secrets

```sql
select
  name,
  file,
  driver,
  external ->> 'Name' as external_name,
  external ->> 'External' as external,
  external -> 'Extensions' as external_extensions
from
  docker_compose_secret;
```

### List secrets with local driver

```sql
select
  name,
  file,
  environment,
  driver,
  template_driver
from
  docker_compose_secret
where
  driver = 'local';
```

### List secrets without environment vars

```sql
select
  name,
  file,
  driver,
  template_driver
from
  docker_compose_secret
where
  environment is null;
```
