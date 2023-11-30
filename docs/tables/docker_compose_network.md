---
title: "Steampipe Table: docker_compose_network - Query Docker Compose Networks using SQL"
description: "Allows users to query Docker Compose Networks, specifically the network configurations in a Docker Compose project."
---

# Table: docker_compose_network - Query Docker Compose Networks using SQL

Docker Compose is a tool that allows developers to define and manage multi-container Docker applications. It utilizes YAML files to configure application services and performs the creation and start-up process of all the containers with a single command. Docker Compose Networks are part of this configuration, allowing for the interconnection of containers within the application.

## Table Usage Guide

The `docker_compose_network` table provides insights into the network configurations within Docker Compose. As a DevOps engineer, explore network-specific details through this table, including network names, drivers, and attached services. Utilize it to uncover information about network configurations, such as those with custom drivers, the interconnection between services, and the verification of network scopes.

## Examples

### Basic info
Explore which Docker networks are attachable and whether they have IPv6 enabled. This allows you to assess your network settings and make necessary adjustments for optimal performance and security.

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
Analyze the settings to understand the internal networks within your Docker Compose setup, particularly focusing on those that are attachable and have IPv6 enabled. This can be useful for managing network configurations and ensuring optimal performance and security.

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
Discover the segments that have attachable networks within your Docker Compose setup. This can help you understand which networks allow dynamic attachment of services, aiding in efficient resource allocation and network management.

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
Discover the segments of your network where IPv6 is activated. This can be useful to understand where potential compatibility or security issues may arise due to the use of this protocol.

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