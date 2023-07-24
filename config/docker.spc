connection "docker" {
  plugin = "docker"

  # dockerfile_paths is a list of locations to search for Dockerfiles
  # dockerfile_paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Wildcard based searches are supported, including recursive searches
  # Local dockerfile_paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.dockerfile" matches all Dockerfiles in the CWD
  #  - "**/*.dockerfile" matches all Dockerfiles in the CWD and all sub-directories
  #  - "../*.dockerfile" matches all Dockerfiles in the CWD's parent directory
  #  - "Dockerfile.*" matches all Dockerfiles starting with "Dockerfile" in the CWD
  #  - "/path/to/dir/*.dockerfile" matches all Dockerfiles in a specific directory
  #  - "/path/to/dir/Dockerfile" matches a specific Dockerfile

  # If dockerfile_paths includes "*", all files (including non-Dockerfiles) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # Defaults to CWD
  dockerfile_paths = ["Dockerfile", "*.dockerfile"]

  # docker_compose_file_paths is a list of locations to search for Docker Compose files
  # docker_compose_file_paths can be configured with a local directory, a remote Git repository URL, or an S3 bucket URL
  # Wildcard based searches are supported, including recursive searches
  # Local docker_compose_file_paths are resolved relative to the current working directory (CWD)

  # For example:
  #  - "*.yml" or "*.yaml" matches all Docker Compose files in the CWD
  #  - "**/*.yml" matches all Docker Compose files in the CWD and all sub-directories
  #  - "../*.yml" matches all Docker Compose files in the CWD's parent directory
  #  - "/path/to/dir/*.yml" matches all Docker Compose files in a specific directory
  #  - "/path/to/dir/compose.yml" matches a specific Docker Compose file

  # If docker_compose_file_paths includes "*", all files (including non-DockerComposefiles) in
  # the CWD will be matched, which may cause errors if incompatible file types exist

  # If docker_compose_file_paths is not set, the plugin will proceed with the below default compose files if they are in CWD:
  # - compose.yaml, compose.yml, docker-compose.yml, docker-compose.yaml

  # Defaults to CWD
  # docker_compose_file_paths = ["compose.yml", "*compose.yml"]

  # Optional docker engine configuration.
  # host        = "tcp://192.168.59.103:2376"
  # cert_path   = "/path/to/my-cert"
  # api_version = "1.41"
  # tls_verify  = true
}
