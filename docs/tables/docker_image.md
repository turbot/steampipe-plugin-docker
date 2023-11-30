---
title: "Steampipe Table: docker_image - Query Docker Images using SQL"
description: "Allows users to query Docker Images, providing details about the images such as image ID, parent ID, and image size."
---

# Table: docker_image - Query Docker Images using SQL

Docker Images are read-only templates that you use to create Docker containers. They are the building blocks of a Docker container and contain the necessary components to run an application. Docker Images are created from a set of instructions, known as a Dockerfile, and can be stored and shared across different Docker hosts.

## Table Usage Guide

The `docker_image` table provides insights into Docker Images within the Docker system. As a DevOps engineer, you can explore image-specific details through this table, including image ID, parent ID, and image size. Utilize it to uncover information about images, such as those with specific tags, the layers within each image, and the verification of image metadata.

## Examples

### List all images
Explore all available images within your Docker system using this query. It's useful for quickly assessing the overall state of your Docker images, helping you manage resources and plan for future needs.

```sql
select
  *
from
  docker_image
```

### Find an image by tag
Explore which Docker images are associated with a specific tag, such as 'postgres:latest'. This can be useful in managing and organizing your Docker images based on their tags.

```sql
select
  *
from
  docker_image
where
  repo_tags ? 'postgres:latest'
```