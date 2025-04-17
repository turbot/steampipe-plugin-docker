## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#86](https://github.com/turbot/steampipe-plugin-docker/pull/86))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#86](https://github.com/turbot/steampipe-plugin-docker/pull/86))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#79](https://github.com/turbot/steampipe-plugin-docker/pull/79))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#79](https://github.com/turbot/steampipe-plugin-docker/pull/79))

## v0.11.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#62](https://github.com/turbot/steampipe-plugin-docker/pull/62))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#62](https://github.com/turbot/steampipe-plugin-docker/pull/62))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-docker/blob/main/docs/LICENSE). ([#62](https://github.com/turbot/steampipe-plugin-docker/pull/62))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#61](https://github.com/turbot/steampipe-plugin-docker/pull/61))

## v0.10.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#48](https://github.com/turbot/steampipe-plugin-docker/pull/48))

## v0.10.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#45](https://github.com/turbot/steampipe-plugin-docker/pull/45))
- Recompiled plugin with Go version `1.21`. ([#45](https://github.com/turbot/steampipe-plugin-docker/pull/45))

## v0.9.0 [2023-09-15]

_Bug fixes_

- Fixed the plugin to return `nil` instead of an `error` when the file/path specified in `dockerfile_paths` or `docker_compose_file_paths` config arguments does not exist. ([#38](https://github.com/turbot/steampipe-plugin-docker/pull/38))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v551-2023-07-26). ([#34](https://github.com/turbot/steampipe-plugin-docker/pull/34))
- Recompiled plugin with `github.com/turbot/go-kit v0.7.0`.

## v0.8.0 [2023-08-29]

_What's new?_

- New tables added
  - [docker_compose_config](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_compose_config)
  - [docker_compose_network](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_compose_network)
  - [docker_compose_secret](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_compose_secret)
  - [docker_compose_service](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_compose_service)
  - [docker_compose_volume](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_compose_volume)
- Added support to query Docker compose files. This can be set using the `docker_compose_file_paths` config argument in the `docker.spc` file.  ([#30](https://github.com/turbot/steampipe-plugin-docker/pull/30))

_Deprecated_

- The `paths` config argument has been deprecated and will be removed in a future release, please use `dockerfile_paths` instead. ([#30](https://github.com/turbot/steampipe-plugin-docker/pull/30))

## v0.7.0 [2023-08-02]

_Enhancements_

- Added `config` and `inspect` columns to the `docker_container` table. ([#31](https://github.com/turbot/steampipe-plugin-docker/pull/31))

## v0.6.0 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#27](https://github.com/turbot/steampipe-plugin-docker/pull/27))

## v0.5.1 [2022-11-28]

_Bug fixes_

- Fixed example queries in README and docs/index.md.

## v0.5.0 [2022-11-17]

_What's new?_

- Added support for retrieving Docker files from remote Git repositories and S3 buckets. For more information, please see [Supported Path Formats](https://hub.steampipe.io/plugins/turbot/docker#supported-path-formats). ([#21](https://github.com/turbot/steampipe-plugin-docker/pull/21))
- Added file watching support for files included in the `paths` config argument. ([#21](https://github.com/turbot/steampipe-plugin-docker/pull/21))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.0.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v500-2022-11-16) which includes support for fetching remote files with go-getter and file watching. ([#21](https://github.com/turbot/steampipe-plugin-docker/pull/21))

## v0.4.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#18](https://github.com/turbot/steampipe-plugin-docker/pull/18))
- Recompiled plugin with Go version `1.19`. ([#18](https://github.com/turbot/steampipe-plugin-docker/pull/18))

## v0.3.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#13](https://github.com/turbot/steampipe-plugin-docker/pull/13))

## v0.3.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#10](https://github.com/turbot/steampipe-plugin-docker/pull/10))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#11](https://github.com/turbot/steampipe-plugin-docker/pull/11))

## v0.2.0 [2022-02-14]

_What's new?_

- File loading and matching through the `paths` argument has been updated to make the plugin easier to use:
  - The `paths` argument is no longer commented out by default for new plugin installations and now defaults to the current working directory
  - Home directory expansion (`~`) is now supported
  - Recursive directory searching (`**`) is now supported
- Previously, when using wildcard matching (`*`), non-Dockerfiles were automatically excluded to prevent parsing errors. These files are no longer automatically excluded to allow for a wider range of matches. If your current configuration uses wildcard matching, e.g., `paths = [ "/path/to/my/files/*" ]`, please update it to include the default Dockerfile name and file extension, e.g., `paths = [ "/path/to/my/files/Dockerfile", "/path/to/my/files/*.dockerfile" ]`.

## v0.1.0 [2021-12-15]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#5](https://github.com/turbot/steampipe-plugin-docker/pull/5))
- Recompiled plugin with Go version 1.17 ([#5](https://github.com/turbot/steampipe-plugin-docker/pull/5))

## v0.0.2 [2021-09-22]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#2](https://github.com/turbot/steampipe-plugin-docker/pull/2))

_Bug fixes_

- Fixed the broken links in docs/index.md file

## v0.0.1 [2021-07-12]

_What's new?_

- New tables added
  - [dockerfile_instruction](https://hub.steampipe.io/plugins/turbot/docker/tables/dockerfile_instruction)
  - [docker_container](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_container)
  - [docker_image](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_image)
  - [docker_info](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_info)
  - [docker_network](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_network)
  - [docker_volume](https://hub.steampipe.io/plugins/turbot/docker/tables/docker_volume)
