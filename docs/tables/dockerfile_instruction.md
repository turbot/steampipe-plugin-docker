# Table: dockerfile_instruction

List and query instructions from Dockerfiles.

## Examples

### List instructions in a specific Dockerfile

Set the `path` column to query a specific Dockerfile. A full path must be provided.

```sql
select
  *
from
  dockerfile_instruction
where
  path = '/full/path/to/Dockerfile'
order by
  start_line
```

### List all Dockerfiles matched in the paths config

The `paths` config parameter sets directories (including wildcards) to search
for Dockerfiles. To match, either the filename is `Dockerfile` (e.g.
`Dockerfile`, `Dockerfile.example`), or the extension is `.dockerfile` (e.g.
`nginx.dockerfile`).

```sql
select
  *
from
  dockerfile_instruction
order by
  path,
  start_line
```

### List base images

```sql
select
  path,
  start_line,
  data ->> 'image' as image,
  data ->> 'tag' as tag
from
  dockerfile_instruction as cmd
where
  cmd.cmd = 'from'
order by
  path,
  start_line,
  image,
  tag
```

### Find all exposed ports

```sql
select
  path,
  start_line,
  (p ->> 'port')::int as port,
  p ->> 'protocol' as protocol
from
  dockerfile_instruction as cmd,
  jsonb_array_elements(data) as p
where
  cmd.cmd = 'expose'
order by
  path,
  start_line,
  port,
  protocol
```
