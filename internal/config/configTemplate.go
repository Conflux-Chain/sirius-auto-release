package config

type WalletConfig struct {
	ChainID           int      `toml:"chainId"`
	ChainName         string   `toml:"chainName"`
	RpcUrls           []string `toml:"rpcUrls"`
	BlockExplorerUrls []string `toml:"blockExplorerUrls"`
	NativeCurrency    struct {
		Name     string `toml:"name"`
		Symbol   string `toml:"symbol"`
		Decimals int    `toml:"decimals"`
	} `toml:"nativeCurrency"`
}

type ThemeConfig struct {
	Primary              string   `toml:"primary,omitempty"`
	AntdPrimaryButtonBg  string   `toml:"antdPrimaryButtonBg,omitempty"`
	ButtonBg             string   `toml:"buttonBg,omitempty"`
	OutlineColor         string   `toml:"outlineColor,omitempty"`
	ShadowColor          string   `toml:"shadowColor,omitempty"`
	SearchButtonBg       string   `toml:"searchButtonBg,omitempty"`
	SearchButtonHoverBg  string   `toml:"searchButtonHoverBg,omitempty"`
	GasPriceLineBg       string   `toml:"gasPriceLineBg,omitempty"`
	FooterBg             string   `toml:"footerBg,omitempty"`
	FooterHighLightColor string   `toml:"footerHighLightColor,omitempty"`
	LinkColor            string   `toml:"linkColor,omitempty"`
	LinkHoverColor       string   `toml:"linkHoverColor,omitempty"`
	ChartColors          []string `toml:"chartColors,omitempty"`
	MixedChartColors     []string `toml:"mixedChartColors,omitempty"`
	PieChartColors       []string `toml:"pieChartColors,omitempty"`
	ChartTitleColor      string   `toml:"chartTitleColor,omitempty"`
	ChartDetailLinkColor string   `toml:"chartDetailLinkColor,omitempty"`
}

type IconsConfig struct {
	ImgArrow string `toml:"imgArrow,omitempty"`
}

type CoreSpaceSettings struct {
	API_URL                      string `toml:"api_url" `
	EnvOpenApiHost               string `toml:"ENV_OPEN_API_HOST,omitempty" `
	EnvRpcServer                 string `toml:"ENV_RPC_SERVER,omitempty"`
	EnvNetworkID                 int    `toml:"ENV_NETWORK_ID,omitempty"`
	EnvNetworkType               string `toml:"ENV_NETWORK_TYPE,omitempty"`
	EnvChainType                 string `toml:"ENV_CHAIN_TYPE,omitempty"`
	EnvFcAddress                 string `toml:"ENV_FC_ADDRESS,omitempty" `
	EnvFcExchangeAddress         string `toml:"ENV_FC_EXCHANGE_ADDRESS,omitempty"`
	EnvFcExchangeInterestAddress string `toml:"ENV_FC_EXCHANGE_INTEREST_ADDRESS,omitempty"`
	EnvEnsRegistryAddress        string `toml:"ENV_ENS_REGISTRY_ADDRESS,omitempty"`
	EnvEnsPublicResolverAddress  string `toml:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS,omitempty"`
	EnvLogo                      string `toml:"ENV_LOGO,omitempty" toml:"ENV_LOGO"`
}

type ESpaceSettings struct {
	API_URL         string            `toml:"api_url" huh_en:"Please input a proxy API URL"`
	EnvApiHost      string            `toml:"ENV_API_HOST,omitempty"`
	EnvCoreApiHost  string            `toml:"ENV_CORE_API_HOST,omitempty"`
	EnvCoreScanHost string            `toml:"ENV_CORE_SCAN_HOST,omitempty"`
	EnvRpcServer    string            `toml:"ENV_RPC_SERVER,omitempty"`
	EnvNetworkID    int               `toml:"ENV_NETWORK_ID,omitempty"`
	EnvNetworkType  string            `toml:"ENV_NETWORK_TYPE,omitempty"`
	EnvChainType    string            `toml:"ENV_CHAIN_TYPE,omitempty"`
	EnvWalletConfig WalletConfig      `toml:"ENV_WALLET_CONFIG,omitempty"`
	EnvLogo         string            `toml:"ENV_LOGO,omitempty"`
	EnvTheme        ThemeConfig       `toml:"ENV_THEME,omitempty"`
	EnvIcons        IconsConfig       `toml:"ENV_ICONS,omitempty"`
	EnvLocalesEn    map[string]string `toml:"ENV_LOCALES_EN,omitempty"`
	EnvLocalesCn    map[string]string `toml:"ENV_LOCALES_CN,omitempty"`
}

type Frontend struct {
	Type string `toml:"type"`
	// SiriusRepo        string            `toml:"sirius_repo,omitempty"`
	// SiriusEthRepo     string            `toml:"sirius_eth_repo,omitempty"`
	PrebuiltRepo      string            `toml:"prebuilt_repo"`
	CoreSpaceSettings CoreSpaceSettings `toml:"core_space_settings"`
	ESpaceSettings    ESpaceSettings    `toml:"e_space_settings"`
}

type Global struct {
	Version int    `toml:"version"`
	Space   string `toml:"space"`
	Workdir string ` toml:"workdir"`
}

type ProxySpace struct {
	Port int `toml:"port"`
}

type Proxy struct {
	Type      string     `toml:"type"`
	CoreSpace ProxySpace `toml:"core_space"`
	ESpace    ProxySpace `toml:"e_space"`
}

type ContainerSpace struct {
	Port int `toml:"port"`
}

type Container struct {
	Enabled bool   `toml:"enabled"`
	Type    string `toml:"type"`
	Name    string `toml:"name"`
	Tag     string `toml:"tag"`

	CoreSpace ContainerSpace `toml:"core_space"`
	ESpace    ContainerSpace `toml:"e_space"`
}

type Config struct {
	Global    Global    `toml:"global"`
	Frontend  Frontend  `toml:"frontend"`
	Container Container `toml:"container"`
	Proxy     Proxy     `toml:"proxy"`
}
