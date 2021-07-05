---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/docker.svg"
brand_color: "#0db7ed"
display_name: "Docker"
short_name: "docker"
description: "Steampipe plugin to query Dockerfile commands and more from Docker."
og_description: "Query Docker with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/docker-social-graphic.png"
---

# Docker + Steampipe

[Docker](https://docker.com) provides OS-level virtualization to deliver software in packages called containers.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Query commands from a Dockerfile:

```sql
select
  cmd,
  args
from
  dockerfile_cmd
where
  path = '/my/Dockerfile'
```

```
+---------+--------------------------------------------------------------------------+
| cmd     | data                                                                     |
+---------+--------------------------------------------------------------------------+
| from    | {"image":"node","tag":"12-alpine"}                                       |
| run     | {"commands":["apk add --no-cache python g++ make"],"prepend_shell":true} |
| workdir | {"path":"/app"}                                                          |
| copy    | {"dest":".","sources":["."]}                                             |
| run     | {"commands":["yarn install --production"],"prepend_shell":true}          |
| cmd     | {"commands":["node","src/index.js"]}                                     |
+---------+--------------------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/docker/tables)**

## Get started

### Install

Download and install the latest Docker plugin:

```bash
steampipe plugin install docker
```

### Credentials

No credentials are required.

### Configuration

Installing the latest docker plugin will create a config file (`~/.steampipe/config/docker.spc`) with a single connection named `docker`:

```hcl
connection "docker" {
  plugin = "docker"
  paths = [ "/path/to/files" ]
}
```

- `paths` - A list of directory paths to search for Dockerfiles. Paths may [include wildcards](https://pkg.go.dev/path/filepath#Match). File matches must start with `Dockerfile` or have an extension of `.dockerfile`.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-docker
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
