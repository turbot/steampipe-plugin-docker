![image](https://hub.steampipe.io/images/plugins/turbot/docker-social-graphic.png)

# Docker Plugin for Steampipe

Use SQL to query Dockerfile configuration and more from Docker.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/docker)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/docker/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-docker/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install docker
```

Configure your [config file](https://hub.steampipe.io/plugins/turbot/docker#configuration) to include directories with Dockerfiles. If no directory is specified, the current working directory will be used.

Run steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  instruction,
  data
from
  dockerfile_instruction
order by
  start_line;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-docker.git
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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-docker/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-docker/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Docker Plugin](https://github.com/turbot/steampipe-plugin-docker/labels/help%20wanted)
