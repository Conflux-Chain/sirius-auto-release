package config

type Config struct {
	Version  int      `mapstructure:"version"`
	Frontend Frontend `mapstructure:"frontend"`
	Docker   Docker   `mapstructure:"docker"`
	Proxy    Proxy    `mapstructure:"proxy"`
}

type Frontend struct {
	Type    string `mapstructure:"type"`
	Source  string `mapstructure:"source"`
	Workdir string `mapstructure:"workdir"`
}

type Docker struct {
	Enabled bool `mapstructure:"enabled"`
}

type Proxy struct {
	Enabled bool   `mapstructure:"enabled"`
	Type    string `mapstructure:"type"`
	Port    int    `mapstructure:"port"`
}
