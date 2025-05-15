package config

import (
	_ "embed"
)

//go:embed nginx.base.conf
var NginxBaseTemplate string

//go:embed nginx.server.conf
var NginxServerTemplate string

//go:embed dockerfile
var DockerTemplate string
