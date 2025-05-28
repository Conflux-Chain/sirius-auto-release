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

//go:embed docker-compose.template.yml
var DockerComposeTemplate string

//go:embed defaultESpaceTheme.toml
var DefaultESpaceTheme string

//go:embed configData.toml
var ConfigDataEmbed string
