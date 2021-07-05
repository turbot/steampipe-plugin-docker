# Table: docker_image

List images from the docker engine.

## Examples

### List all images

```sql
select
  *
from
  docker_image
```

### Find an image by tag

```sql
select
  *
from
  docker_image
where
  repo_tags ? 'postgres:latest'
```
