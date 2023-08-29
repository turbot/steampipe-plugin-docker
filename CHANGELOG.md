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
