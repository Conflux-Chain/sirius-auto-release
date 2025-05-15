package config

type Config struct {
	Global    Global    `mapstructure:"global"`
	Frontend  Frontend  `mapstructure:"frontend"`
	Container Container `mapstructure:"container"`
	Proxy     Proxy     `mapstructure:"proxy"`
}

type Global struct {
	Version int    `mapstructure:"version"`
	Space   string `mapstructure:"space"`
	Workdir string `mapstructure:"workdir"`
}

type Frontend struct {
	Type          string `mapstructure:"type"`
	SiriusRepo    string `mapstructure:"sirius_repo"`
	SiriusEthRepo string `mapstructure:"sirius_eth_repo"`
	PrebuiltRepo  string `mapstructure:"prebuilt_repo"`
}

type ProxySpace struct {
	Port    int    `mapstructure:"port"`
	API_URL string `mapstructure:"api_url"`
}

type Proxy struct {
	Enabled   bool       `mapstructure:"enabled"`
	Type      string     `mapstructure:"type"`
	CoreSpace ProxySpace `mapstructure:"core_space"`
	ESpace    ProxySpace `mapstructure:"e_space"`
}

type ContainerSpace struct {
	Port int `mapstructure:"port"`
}

type Container struct {
	Enabled bool   `mapstructure:"enabled"`
	Type    string `mapstructure:"type"`
	Name    string `mapstructure:"name"`
	Tag     string `mapstructure:"tag"`

	CoreSpace ContainerSpace `mapstructure:"core_space"`
	ESpace    ContainerSpace `mapstructure:"e_space"`
}
