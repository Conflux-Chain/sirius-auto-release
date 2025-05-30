package config

type WalletConfig struct {
	ChainID           int      `toml:"chainId" json:"chainId"`
	ChainName         string   `toml:"chainName" json:"chainName"`
	RpcUrls           []string `toml:"rpcUrls" json:"rpcUrls"`
	BlockExplorerUrls []string `toml:"blockExplorerUrls" json:"blockExplorerUrls"`
	NativeCurrency    struct {
		Name     string `toml:"name" json:"name"`
		Symbol   string `toml:"symbol" json:"symbol"`
		Decimals int    `toml:"decimals" json:"decimals"`
	} `toml:"nativeCurrency" json:"nativeCurrency"`
}
type ThemeConfig struct {
	Primary              string   `toml:"primary,omitempty" json:"primary,omitempty"`
	AntdPrimaryButtonBg  string   `toml:"antdPrimaryButtonBg,omitempty" json:"antdPrimaryButtonBg,omitempty"`
	ButtonBg             string   `toml:"buttonBg,omitempty" json:"buttonBg,omitempty"`
	OutlineColor         string   `toml:"outlineColor,omitempty" json:"outlineColor,omitempty"`
	ShadowColor          string   `toml:"shadowColor,omitempty" json:"shadowColor,omitempty"`
	SearchButtonBg       string   `toml:"searchButtonBg,omitempty" json:"searchButtonBg,omitempty"`
	SearchButtonHoverBg  string   `toml:"searchButtonHoverBg,omitempty" json:"searchButtonHoverBg,omitempty"`
	GasPriceLineBg       string   `toml:"gasPriceLineBg,omitempty" json:"gasPriceLineBg,omitempty"`
	FooterBg             string   `toml:"footerBg,omitempty" json:"footerBg,omitempty"`
	FooterHighLightColor string   `toml:"footerHighLightColor,omitempty" json:"footerHighLightColor,omitempty"`
	LinkColor            string   `toml:"linkColor,omitempty" json:"linkColor,omitempty"`
	LinkHoverColor       string   `toml:"linkHoverColor,omitempty" json:"linkHoverColor,omitempty"`
	ChartColors          []string `toml:"chartColors,omitempty" json:"chartColors,omitempty"`
	MixedChartColors     []string `toml:"mixedChartColors,omitempty" json:"mixedChartColors,omitempty"`
	PieChartColors       []string `toml:"pieChartColors,omitempty" json:"pieChartColors,omitempty"`
	ChartTitleColor      string   `toml:"chartTitleColor,omitempty" json:"chartTitleColor,omitempty"`
	ChartDetailLinkColor string   `toml:"chartDetailLinkColor,omitempty" json:"chartDetailLinkColor,omitempty"`
}
type IconsConfig struct {
	ImgArrow string `toml:"imgArrow,omitempty" json:"imgArrow,omitempty"`
}
type CoreSpaceSettings struct {
	API_URL                      string `toml:"api_url,omitempty" json:"-"`
	EnvOpenApiHost               string `toml:"ENV_OPEN_API_HOST,omitempty" json:"ENV_OPEN_API_HOST,omitempty"`
	EnvRpcServer                 string `toml:"ENV_RPC_SERVER,omitempty" json:"ENV_RPC_SERVER,omitempty"`
	EnvNetworkID                 int    `toml:"ENV_NETWORK_ID,omitempty" json:"ENV_NETWORK_ID,omitempty"`
	EnvNetworkType               string `toml:"ENV_NETWORK_TYPE,omitempty" json:"ENV_NETWORK_TYPE,omitempty"`
	EnvChainType                 string `toml:"ENV_CHAIN_TYPE,omitempty" json:"ENV_CHAIN_TYPE,omitempty"`
	EnvFcAddress                 string `toml:"ENV_FC_ADDRESS,omitempty" json:"ENV_FC_ADDRESS,omitempty"`
	EnvFcExchangeAddress         string `toml:"ENV_FC_EXCHANGE_ADDRESS,omitempty" json:"ENV_FC_EXCHANGE_ADDRESS,omitempty"`
	EnvFcExchangeInterestAddress string `toml:"ENV_FC_EXCHANGE_INTEREST_ADDRESS,omitempty" json:"ENV_FC_EXCHANGE_INTEREST_ADDRESS,omitempty"`
	EnvEnsRegistryAddress        string `toml:"ENV_ENS_REGISTRY_ADDRESS,omitempty" json:"ENV_ENS_REGISTRY_ADDRESS,omitempty"`
	EnvEnsPublicResolverAddress  string `toml:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS,omitempty" json:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS,omitempty"`
	EnvLogo                      string `toml:"ENV_LOGO,omitempty" json:"ENV_LOGO,omitempty"`
}
type ESpaceSettings struct {
	API_URL         string            `toml:"api_url,omitempty" json:"-"`
	EnvApiHost      string            `toml:"ENV_API_HOST,omitempty" json:"ENV_API_HOST,omitempty"`
	EnvCoreApiHost  string            `toml:"ENV_CORE_API_HOST,omitempty" json:"ENV_CORE_API_HOST,omitempty"`
	EnvCoreScanHost string            `toml:"ENV_CORE_SCAN_HOST,omitempty" json:"ENV_CORE_SCAN_HOST,omitempty"`
	EnvRpcServer    string            `toml:"ENV_RPC_SERVER,omitempty" json:"ENV_RPC_SERVER,omitempty"`
	EnvNetworkID    int               `toml:"ENV_NETWORK_ID,omitempty" json:"ENV_NETWORK_ID,omitempty"`
	EnvNetworkType  string            `toml:"ENV_NETWORK_TYPE,omitempty" json:"ENV_NETWORK_TYPE,omitempty"`
	EnvChainType    string            `toml:"ENV_CHAIN_TYPE,omitempty" json:"ENV_CHAIN_TYPE,omitempty"`
	EnvWalletConfig WalletConfig      `toml:"ENV_WALLET_CONFIG,omitempty" json:"ENV_WALLET_CONFIG"`
	EnvLogo         string            `toml:"ENV_LOGO,omitempty" json:"ENV_LOGO,omitempty"`
	EnvTheme        ThemeConfig       `toml:"ENV_THEME,omitempty" json:"ENV_THEME"`
	EnvIcons        IconsConfig       `toml:"ENV_ICONS,omitempty" json:"ENV_ICONS"`
	EnvLocalesEn    map[string]string `toml:"ENV_LOCALES_EN,omitempty" json:"ENV_LOCALES_EN,omitempty"`
	EnvLocalesCn    map[string]string `toml:"ENV_LOCALES_CN,omitempty" json:"ENV_LOCALES_CN,omitempty"`
}
type Frontend struct {
	Type string `toml:"type"`
	// SiriusRepo        string            `toml:"sirius_repo,omitempty"`
	// SiriusEthRepo     string            `toml:"sirius_eth_repo,omitempty"`
	PrebuiltRepo      string            `toml:"prebuilt_repo"`
	CoreSpaceSettings CoreSpaceSettings `toml:"core_space_settings,omitempty"`
	ESpaceSettings    ESpaceSettings    `toml:"e_space_settings,omitempty"`
}

type Global struct {
	Version int    `toml:"version"`
	Space   string `toml:"space"`
	Workdir string `toml:"workdir"`
}

type ProxySpace struct {
	Port int `toml:"port,omitempty"`
}

type Proxy struct {
	Type      string     `toml:"type"`
	CoreSpace ProxySpace `toml:"core_space,omitempty"`
	ESpace    ProxySpace `toml:"e_space,omitempty"`
}

type ContainerSpace struct {
	Port int `toml:"port,omitempty"`
}

type Container struct {
	Enabled bool   `toml:"enabled"`
	Type    string `toml:"type"`
	Name    string `toml:"name"`
	Tag     string `toml:"tag"`

	CoreSpace ContainerSpace `toml:"core_space,omitempty"`
	ESpace    ContainerSpace `toml:"e_space,omitempty"`
}

type Config struct {
	Global    Global    `toml:"global"`
	Frontend  Frontend  `toml:"frontend"`
	Container Container `toml:"container"`
	Proxy     Proxy     `toml:"proxy"`
}
