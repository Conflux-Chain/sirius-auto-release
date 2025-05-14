package config

import (
	_ "embed"
)

//go:embed nginx.conf
var NginxTemplate string
