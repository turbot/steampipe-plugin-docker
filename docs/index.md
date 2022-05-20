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
  path = '/my/Dockerfile';
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

  # Paths is a list of locations to search for Dockerfiles
  # All paths are resolved relative to the current working directory (CWD)
  # Wildcard based searches are supported, including recursive searches

  # For example:
  #  - "*.dockerfile" matches all Dockerfiles in the CWD
  #  - "**/*.dockerfile" matches all Dockerfiles in the CWD and all sub-directories
  #  - "../*.dockerfile" matches all Dockerfiles in the CWD's parent directory
  #  - "Dockerfile.*" matches all Dockerfiles starting with "Dockerfile" in the CWD
  #  - "/path/to/dir/*.dockerfile" matches all Dockerfiles in a specific directory
  #  - "/path/to/dir/Dockerfile" matches a specific file

  # If paths includes "*", all files (including non-Dockerfiles) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # Defaults to CWD
  paths = [ "Dockerfile", "*.dockerfile" ]

  # Optional docker engine configuration.
  # host        = "tcp://192.168.59.103:2376"
  # cert_path   = "/path/to/my-cert"
  # api_version = "1.41"
  # tls_verify  = true
}
```

- `paths` - A list of directory paths to search for Dockerfiles. Paths are resolved relative to the current working directory. Paths may [include wildcards](https://pkg.go.dev/path/filepath#Match) and also supports `**` for recursive matching. Defaults to the current working directory.
- `host` - Location of the docker engine endpoint. Defaults to `DOCKER_HOST` env var.
- `api_version` - API version to use. Defaults to `DOCKER_API_VERSION` env var.
- `cert_path` - Path to a custom TLS certificate. Defaults to `DOCKER_CERT_PATH` env var.
- `tls_verify` - Flag to control TLS verification. Defaults to `DOCKER_TLS_VERIFY` env var.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-docker
- Community: [Slack Channel](https://steampipe.io/community/join)
