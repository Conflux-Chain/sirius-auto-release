package config

import (
	_ "embed"
)

//go:embed nginx.conf
var NginxTemplate string


//go:embed dockerfile
var DockerTemplate string
