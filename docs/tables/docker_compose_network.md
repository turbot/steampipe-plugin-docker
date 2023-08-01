# Table: docker_compose_network

List all networks from the Docker compose files.

## Examples

### Basic info

```sql
select
  name,
  driver,
  internal,
  attachable,
  enable_ipv6
from
  docker_compose_network;
```

### List internal networks

```sql
select
  name,
  driver,
  internal,
  attachable,
  enable_ipv6
from
  docker_compose_network
where
  internal;
```

### List attachable networks

```sql
select
  name,
  driver,
  internal,
  attachable,
  enable_ipv6
from
  docker_compose_network
where
  attachable;
```

### List networks where IPv6 is enabled

```sql
select
  name,
  driver,
  internal,
  attachable,
  enable_ipv6
from
  docker_compose_network
where
  enable_ipv6;
```
