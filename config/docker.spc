connection "docker" {
  plugin = "docker"

  # Paths is a list of locations to search for Dockerfiles by default.
  # Wildcards are supported per https://golang.org/pkg/path/filepath/#Match
  # Exact file paths can have any name. Wildcard based matches must either
  # have a name of Dockerfile (e.g. Dockerfile, Dockerfile.example) or an
  # .dockerfile extension (e.g. nginx.dockerfile).
  # paths = [ "/path/to/dir/*", "/path/to/exact/custom-dockerfile-name" ]
}
