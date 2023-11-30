---
title: "Steampipe Table: docker_network - Query Docker Networks using SQL"
description: "Allows users to query Docker Networks, specifically providing details about network configuration, network options, and network services."
---

# Table: docker_network - Query Docker Networks using SQL

Docker Networks are a crucial component in Docker's architecture, enabling containers to communicate with each other and with outside networks. They are designed to facilitate network isolation and service discovery, which are essential for microservice architectures. Docker Networks also provide the flexibility to design and implement networking that fits a more complex application architecture.

## Table Usage Guide

The `docker_network` table offers insights into Docker Networks. Network administrators and DevOps engineers can utilize this table to obtain detailed information about Docker Networks, including their configuration, options, and the services running on them. This information can be valuable for troubleshooting network issues, optimizing network performance, and ensuring network security.

## Examples

### List all networks
Explore all the networks within your Docker environment. This can help in understanding the connectivity and isolation of your containers, aiding in better network management and security.

```sql
select
  *
from
  docker_network
```