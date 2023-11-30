---
title: "Steampipe Table: docker_compose_secret - Query Docker Compose Secrets using SQL"
description: "Allows users to query Docker Compose Secrets, providing insights into the secrets used by services in a Docker Compose project."
---

# Table: docker_compose_secret - Query Docker Compose Secrets using SQL

Docker Compose is a tool for defining and running multi-container Docker applications. It uses YAML files to configure the application's services and performs the creation and start-up process of all the containers with a single command. Docker Compose Secrets are a resource that allows you to securely store sensitive information used by services in a Docker Compose project.

## Table Usage Guide

The `docker_compose_secret` table provides insights into the secrets used in a Docker Compose project. As a developer or system administrator, you can explore secret-specific details through this table, including the secret name, service using the secret, and the file path of the secret. Utilize this table to manage and monitor the use of secrets across your Docker Compose projects, ensuring secure and efficient use of sensitive information.

## Examples

### Basic info
Explore which Docker Compose secrets are being used in your environment. This can help you manage and understand the configuration of your secrets, providing insights into your Docker Compose setup.

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
Explore the external configuration of secrets in your Docker Compose setup to understand how they are managed and where they are stored. This is beneficial for assessing security measures and ensuring best practices are in place.

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
Discover the segments that utilize the local driver within the Docker Compose Secret. This is particularly beneficial to identify and manage secrets that are locally stored, aiding in security and configuration management.

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
Determine the areas in which Docker Compose secrets are not associated with any environment variables. This can be useful to identify potential misconfigurations or security risks within your Docker setup.

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