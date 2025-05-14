package config

type Config struct {
	Global   Global   `mapstructure:"global"`
	Frontend Frontend `mapstructure:"frontend"`
	Docker   Docker   `mapstructure:"docker"`
	Proxy    Proxy    `mapstructure:"proxy"`
}

type Global struct {
	Version int    `mapstructure:"version"`
	Workdir string `mapstructure:"workdir"`
}

type Frontend struct {
	Type          string `mapstructure:"type"`
	SiriusRepo    string `mapstructure:"sirius_repo"`
	SiriusEthRepo string `mapstructure:"sirius_eth_repo"`
	PrebuiltRepo  string `mapstructure:"prebuilt_repo"`
}

type Proxy struct {
	Enabled bool   `mapstructure:"enabled"`
	Type    string `mapstructure:"type"`
	Port    int    `mapstructure:"port"`
	API_URL string `mapstructure:"api_url"`
}

type Docker struct {
	Enabled bool `mapstructure:"enabled"`
}
