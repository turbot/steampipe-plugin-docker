---
title: "Steampipe Table: docker_info - Query Docker Information using SQL"
description: "Allows users to query Docker Information, specifically the configuration details of the Docker engine running on the host machine."
---

# Table: docker_info - Query Docker Information using SQL

Docker is an open-source platform that automates the deployment, scaling, and management of applications inside lightweight, portable containers. It provides an additional layer of abstraction and automation of operating-system-level virtualization on Windows and Linux. Docker uses the resource isolation features of the Linux kernel to allow independent containers to run within a single Linux instance.

## Table Usage Guide

The `docker_info` table provides insights into the Docker engine running on the host machine. As a system administrator or a DevOps engineer, you can explore details about the Docker engine through this table, including its configuration, version, and runtime information. Utilize it to uncover information about the Docker engine, such as its operating system, architecture, number of containers, and images.

## Examples

### Get info
Explore the comprehensive details of your Docker environment to better understand its current state and configuration. This can aid in troubleshooting, optimizing resource usage, and enhancing your overall Docker management strategy.

```sql
select
  *
from
  docker_info
```