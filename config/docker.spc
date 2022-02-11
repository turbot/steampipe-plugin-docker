connection "docker" {
  plugin = "docker"

  # Paths is a list of locations to search for Dockerfiles
  # All paths are resolved relative to the current working directory (CWD)
  # Wildcard based searches are supported, including recursive searches

  # For example:
  #  - "*.dockerfile" matches all Docker files in the CWD
  #  - "**/*.dockerfile" matches all Docker files in the CWD and all sub-directories
  #  - "../*.dockerfile" matches all Docker files in the CWD's parent directory
  #  - "Dockerfile" matches all Docker files named "Dockerfile" in the current CWD
  #  - "/path/to/dir/*.dockerfile" matches all Docker files in a specific directory
  #  - "/path/to/dir/Dockerfile" matches a specific file

  # If paths includes "*", all files (including non-Docker files) in
  # the current CWD will be matched, which may cause errors if incompatible filetypes exist

  # Defaults to CWD
  paths = [ "Dockerfile", "*.dockerfile" ]

  # Optional docker engine configuration.
  # host        = "tcp://192.168.59.103:2376"
  # cert_path   = "/path/to/my-cert"
  # api_version = "1.41"
  # tls_verify  = true
}