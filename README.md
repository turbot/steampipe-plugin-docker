![image](https://hub.steampipe.io/images/plugins/turbot/docker-social-graphic.png)

# Docker Plugin for Steampipe

Use SQL to query Dockerfile configuration and more from Docker.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/docker)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/docker/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-docker/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install docker
```

Run a query:

```sql
select
  stage,
  cmd,
  args
from
  dockerfile_cmd
where
  path = '/my/Dockerfile'
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone git@github.com:turbot/steampipe-plugin-docker
cd steampipe-plugin-docker
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/docker.spc
```

Try it!

```
steampipe query
> .inspect docker
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-docker/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Docker Plugin](https://github.com/turbot/steampipe-plugin-docker/labels/help%20wanted)
