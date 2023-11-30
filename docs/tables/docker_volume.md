---
title: "Steampipe Table: docker_volume - Query Docker Volumes using SQL"
description: "Allows users to query Docker Volumes, specifically retrieving detailed information about each volume's configuration and status."
---

# Table: docker_volume - Query Docker Volumes using SQL

Docker Volumes are a type of resource provided by the Docker service that allows persistent data storage and sharing among containers. They are designed to decouple the life of the data being stored from the life of the container, enabling data to persist even after a container is deleted. Docker Volumes are managed directly by Docker and are stored in a part of the host filesystem which is managed by Docker.

## Table Usage Guide

The `docker_volume` table provides insights into Docker Volumes within Docker. As a DevOps engineer, explore volume-specific details through this table, including driver, labels, options, and scope. Utilize it to uncover information about volumes, such as those with specific labels, the options set for each volume, and the scope of each volume.

## Examples

### List all volumes
Explore all the storage volumes within your Docker environment to understand the usage and allocation. This can be beneficial for managing storage resources and optimizing performance.

```sql
select
  *
from
  docker_volume
```

### Find all volumes with a given label
Identify all Docker volumes that are labeled as 'postgres'. This is useful for managing and organizing your Docker volumes, especially in large systems where volumes may not be consistently labeled.

```sql
select
  *
from
  docker_volume
where
  labels ->> 'com.docker.compose.volume' = 'postgres'
```