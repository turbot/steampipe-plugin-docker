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
