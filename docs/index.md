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
  # Paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Wildcard based searches are supported, including recursive searches
  # Local paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.dockerfile" matches all Dockerfiles in the CWD
  #  - "**/*.dockerfile" matches all Dockerfiles in the CWD and all sub-directories
  #  - "../*.dockerfile" matches all Dockerfiles in the CWD's parent directory
  #  - "Dockerfile.*" matches all Dockerfiles starting with "Dockerfile" in the CWD
  #  - "/path/to/dir/*.dockerfile" matches all Dockerfiles in a specific directory
  #  - "/path/to/dir/Dockerfile" matches a specific Dockerfile

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

- `host` - Location of the docker engine endpoint. Defaults to `DOCKER_HOST` env var.
- `api_version` - API version to use. Defaults to `DOCKER_API_VERSION` env var.
- `cert_path` - Path to a custom TLS certificate. Defaults to `DOCKER_CERT_PATH` env var.
- `tls_verify` - Flag to control TLS verification. Defaults to `DOCKER_TLS_VERIFY` env var.

### Supported Path Formats

The `paths` config argument is flexible and can search for Dockerfiles from several different sources, e.g., local directory paths, Git, S3.

The following sources are supported:

- [Local files](#configuring-local-file-paths)
- [Remote Git repositories](#configuring-remote-git-repository-urls)
- [S3](#configuring-s3-urls)

Paths may [include wildcards](https://pkg.go.dev/path/filepath#Match) and support `**` for recursive matching. For example:

```hcl
connection "docker" {
  plugin = "docker"

  paths = [
    "Dockerfile",
    "*.dockerfile",
    "~/*.dockerfile",
    "github.com/komljen/dockerfile-examples//*.dockerfile",
    "github.com/komljen/dockerfile-examples//**/Dockerfile",
    "gitlab.com/gitlab-examples/docker//Dockerfile",
    "bitbucket.org/fscm/docker-docker//Dockerfile",
    "s3::https://bucket.s3.ap-southeast-1.amazonaws.com/Dockerfile"
  ]
}
```

**Note**: If any path matches on `*` without `Dockerfile` or `*.dockerfile`, all files (including non-Dockerfiles) in the directory will be matched, which may cause errors if incompatible file types exist.

#### Configuring Local File Paths

You can define a list of local directory paths to search for Dockerfiles. Paths are resolved relative to the current working directory. For example:

- `*.dockerfile` matches all Dockerfiles in the CWD.
- `**/*.dockerfile` matches all Dockerfiles in the CWD and all sub-directories.
- `../*.dockerfile` matches all Dockerfiles in the CWD's parent directory.
- `Dockerfile.*` matches all Dockerfiles starting with `Dockerfile` in the CWD.
- `/path/to/dir/*.dockerfile` matches all Dockerfiles in a specific directory. For example:
  - `~/*.dockerfile` matches all Dockerfiles in the home directory.
  - `~/**/*.dockerfile` matches all Dockerfiles recursively in the home directory.
- `/path/to/dir/Dockerfile` matches a specific file.

```hcl
connection "docker" {
  plugin = "docker"

  paths = [ "*.dockerfile", "~/*.dockerfile", "/path/to/dir/Dockerfile" ]
}
```

#### Configuring Remote Git Repository URLs

You can also configure `paths` with any Git remote repository URLs, e.g., GitHub, BitBucket, GitLab. The plugin will then attempt to retrieve any Dockerfiles from the remote repositories.

For example:

- `github.com/komljen/dockerfile-examples//*.dockerfile` matches all top-level Dockerfiles in the specified github repository.
- `github.com/komljen/dockerfile-examples//**/Dockerfile` matches all Dockerfiles in the specified github repository and all sub-directories.
- `github.com/komljen/dockerfile-examples?ref=fix_7677//**/Dockerfile` matches all Dockerfiles in the specific tag of github repository.
- `github.com/komljen/dockerfile-examples//ghost//Dockerfile` matches all Dockerfiles in the specified folder path.

You can specify a subdirectory after a double-slash (`//`) if you want to download only a specific subdirectory from a downloaded directory.

```hcl
connection "docker" {
  plugin = "docker"

  paths = [ "github.com/komljen/dockerfile-examples//ghost//Dockerfile" ]
}
```

Similarly, you can define a list of GitLab and BitBucket URLs to search for Dockerfiles:

```hcl
connection "docker" {
  plugin = "docker"

  paths = [
    "github.com/komljen/dockerfile-examples//**/Dockerfile",
    "github.com/komljen/dockerfile-examples//ghost//Dockerfile",
    "gitlab.com/gitlab-examples/docker//Dockerfile",
    "gitlab.com/gitlab-examples/docker//**/Dockerfile"
    "bitbucket.org/fscm/docker-docker//Dockerfile",
    "bitbucket.org/fscm/docker-docker//**/Dockerfile"
  ]
}
```

#### Configuring S3 URLs

You can also query all Dockerfiles stored inside an S3 bucket (public or private) using the bucket URL.

##### Accessing a Private Bucket

In order to access your files in a private S3 bucket, you will need to configure your credentials. You can use your configured AWS profile from local `~/.aws/config`, or pass the credentials using the standard AWS environment variables, e.g., `AWS_PROFILE`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`.

We recommend using AWS profiles for authentication.

**Note:** Make sure that `region` is configured in the config. If not set in the config, `region` will be fetched from the standard environment variable `AWS_REGION`.

You can also authenticate your request by setting the AWS profile and region in `paths`. For example:

```hcl
connection "docker" {
  plugin = "docker"

  paths = [
    "s3::https://bucket-2.s3.us-east-1.amazonaws.com//Dockerfile?aws_profile=<AWS_PROFILE>",
    "s3::https://bucket-2.s3.us-east-1.amazonaws.com/test_folder//*.dockerfile?aws_profile=<AWS_PROFILE>"
  ]
}
```

**Note:**

In order to access the bucket, the IAM user or role will require the following IAM permissions:

- `s3:ListBucket`
- `s3:GetObject`
- `s3:GetObjectVersion`

If the bucket is in another AWS account, the bucket policy will need to grant access to your user or role. For example:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ReadBucketObject",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:user/YOUR_USER"
      },
      "Action": ["s3:ListBucket", "s3:GetObject", "s3:GetObjectVersion"],
      "Resource": ["arn:aws:s3:::test-bucket1", "arn:aws:s3:::test-bucket1/*"]
    }
  ]
}
```

##### Accessing a Public Bucket

Public access granted to buckets and objects through ACLs and bucket policies allows any user access to data in the bucket. We do not recommend making S3 buckets public, but if there are specific objects you'd like to make public, please see [How can I grant public read access to some objects in my Amazon S3 bucket?](https://aws.amazon.com/premiumsupport/knowledge-center/read-access-objects-s3-bucket/).

You can query any public S3 bucket directly using the URL without passing credentials. For example:

```hcl
connection "docker" {
  plugin = "docker"

  paths = [
    "s3::https://bucket-1.s3.us-east-1.amazonaws.com/test_folder//Dockerfile",
    "s3::https://bucket-2.s3.us-east-1.amazonaws.com/test_folder//**/*.dockerfile"
  ]
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-docker
- Community: [Slack Channel](https://steampipe.io/community/join)
