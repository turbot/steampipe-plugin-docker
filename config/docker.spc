connection "docker" {
  plugin = "docker"

  # Paths is a list of locations to search for Dockerfiles
  # All paths are resolved relative to the current working directory (CWD)
  # Wildcard based searches are supported, including recursive searches
  # Defaults to CWD
  paths = [ "Dockerfile", "*.dockerfile" ]

  # Optional docker engine configuration.
  # host        = "tcp://192.168.59.103:2376"
  # cert_path   = "/path/to/my-cert"
  # api_version = "1.41"
  # tls_verify  = true
}
