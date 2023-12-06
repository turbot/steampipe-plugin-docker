---
title: "Steampipe Table: dockerfile_instruction - Query Docker Dockerfile Instructions using SQL"
description: "Allows users to query Dockerfile Instructions in Docker, specifically the different instructions that can be used in a Dockerfile, providing insights into the building blocks of Docker images."
---

# Table: dockerfile_instruction - Query Docker Dockerfile Instructions using SQL

Dockerfile Instructions are the commands that are used in a Dockerfile to build a Docker image. These instructions specify what to include in the Docker image and how it should behave when it is run. They are the building blocks of Docker images, specifying everything from the base image to use, to the commands to run, to the metadata to include.

## Table Usage Guide

The `dockerfile_instruction` table provides insights into the instructions used within Dockerfiles in Docker. As a DevOps engineer, explore instruction-specific details through this table, including the instruction type, arguments, and associated metadata. Utilize it to uncover information about instructions, such as what base images are used, the commands that are run, and the metadata that is included.

## Examples

### List instructions in a specific Dockerfile
Explore the sequence of instructions within a specific Dockerfile to better understand the build process and dependencies. This could be beneficial in assessing the complexity of the build, identifying potential areas for optimization, or troubleshooting issues.
Set the `path` column to query a specific Dockerfile. A full path must be provided.


```sql+postgres
select
  *
from
  dockerfile_instruction
where
  path = '/full/path/to/Dockerfile'
order by
  start_line;
```

```sql+sqlite
select
  *
from
  dockerfile_instruction
where
  path = '/full/path/to/Dockerfile'
order by
  start_line;
```

### List all Dockerfiles matched in the paths config
Explore all Dockerfile instructions sorted by their respective paths and starting lines. This can help you understand the structure and organization of your Dockerfiles, making it easier to manage and troubleshoot your Docker environment.

The `paths` config parameter sets directories (including wildcards) to search
for Dockerfiles. To match, either the filename is `Dockerfile` (e.g.
`Dockerfile`, `Dockerfile.example`), or the extension is `.dockerfile` (e.g.
`nginx.dockerfile`).

```sql+postgres
select
  *
from
  dockerfile_instruction
order by
  path,
  start_line;
```

```sql+sqlite
select
  *
from
  dockerfile_instruction
order by
  path,
  start_line;
```

### List base images
Explore the foundational elements of your Docker environment by identifying the base images used in your Dockerfiles. This can aid in understanding dependencies, ensuring consistency, and managing potential security vulnerabilities across your projects.

```sql+postgres
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
  tag;
```

```sql+sqlite
select
  path,
  start_line,
  json_extract(data, '$.image') as image,
  json_extract(data, '$.tag') as tag
from
  dockerfile_instruction as cmd
where
  cmd.cmd = 'from'
order by
  path,
  start_line,
  image,
  tag;
```

### Find all exposed ports
Identify instances where certain ports are exposed in your Dockerfile instructions. This can help you manage and secure your network traffic by understanding which ports are open.

```sql+postgres
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
  protocol;
```

```sql+sqlite
select
  path,
  start_line,
  cast(json_extract(p.value, '$.port') as integer) as port,
  json_extract(p.value, '$.protocol') as protocol
from
  dockerfile_instruction as cmd,
  json_each(cmd.data) as p
where
  cmd.cmd = 'expose'
order by
  path,
  start_line,
  port,
  protocol;
```