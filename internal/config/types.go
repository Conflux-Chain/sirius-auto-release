package config

type WalletConfig struct {
	ChainID           int      `mapstructure:"chainId"`
	ChainName         string   `mapstructure:"chainName"`
	RpcUrls           []string `mapstructure:"rpcUrls"`
	BlockExplorerUrls []string `mapstructure:"blockExplorerUrls"`
	NativeCurrency    struct {
		Name     string `mapstructure:"name"`
		Symbol   string `mapstructure:"symbol"`
		Decimals int    `mapstructure:"decimals"`
	} `mapstructure:"nativeCurrency"`
}

type ThemeConfig struct {
	Primary              string   `mapstructure:"primary,omitempty"`
	AntdPrimaryButtonBg  string   `mapstructure:"antdPrimaryButtonBg,omitempty"`
	ButtonBg             string   `mapstructure:"buttonBg,omitempty"`
	OutlineColor         string   `mapstructure:"outlineColor,omitempty"`
	ShadowColor          string   `mapstructure:"shadowColor,omitempty"`
	SearchButtonBg       string   `mapstructure:"searchButtonBg,omitempty"`
	SearchButtonHoverBg  string   `mapstructure:"searchButtonHoverBg,omitempty"`
	GasPriceLineBg       string   `mapstructure:"gasPriceLineBg,omitempty"`
	FooterBg             string   `mapstructure:"footerBg,omitempty"`
	FooterHighLightColor string   `mapstructure:"footerHighLightColor,omitempty"`
	LinkColor            string   `mapstructure:"linkColor,omitempty"`
	LinkHoverColor       string   `mapstructure:"linkHoverColor,omitempty"`
	ChartColors          []string `mapstructure:"chartColors,omitempty"`
	MixedChartColors     []string `mapstructure:"mixedChartColors,omitempty"`
	PieChartColors       []string `mapstructure:"pieChartColors,omitempty"`
	ChartTitleColor      string   `mapstructure:"chartTitleColor,omitempty"`
	ChartDetailLinkColor string   `mapstructure:"chartDetailLinkColor,omitempty"`
}

type IconsConfig struct {
	ImgArrow string `mapstructure:"imgArrow,omitempty"`
}

type CoreSpaceSettings struct {
	Enabled                      bool   `mapstructure:"enabled,omitempty" json:"-"`
	EnvOpenApiHost               string `mapstructure:"ENV_OPEN_API_HOST,omitempty" json:"ENV_OPEN_API_HOST"`
	EnvRpcServer                 string `mapstructure:"ENV_RPC_SERVER,omitempty" json:"ENV_RPC_SERVER"`
	EnvNetworkID                 int    `mapstructure:"ENV_NETWORK_ID,omitempty" json:"ENV_NETWORK_ID"`
	EnvNetworkType               string `mapstructure:"ENV_NETWORK_TYPE,omitempty" json:"ENV_NETWORK_TYPE"`
	EnvChainType                 string `mapstructure:"ENV_CHAIN_TYPE,omitempty" json:"ENV_CHAIN_TYPE"`
	EnvFcAddress                 string `mapstructure:"ENV_FC_ADDRESS,omitempty" json:"ENV_FC_ADDRESS"`
	EnvFcExchangeAddress         string `mapstructure:"ENV_FC_EXCHANGE_ADDRESS,omitempty" json:"ENV_FC_EXCHANGE_ADDRESS"`
	EnvFcExchangeInterestAddress string `mapstructure:"ENV_FC_EXCHANGE_INTEREST_ADDRESS,omitempty" json:"ENV_FC_EXCHANGE_INTEREST_ADDRESS"`
	EnvEnsRegistryAddress        string `mapstructure:"ENV_ENS_REGISTRY_ADDRESS,omitempty" json:"ENV_ENS_REGISTRY_ADDRESS"`
	EnvEnsPublicResolverAddress  string `mapstructure:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS,omitempty" json:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS"`
	EnvLogo                      string `mapstructure:"ENV_LOGO,omitempty" json:"ENV_LOGO"`
}

type ESpaceSettings struct {
	Enabled         bool              `mapstructure:"enabled,omitempty" json:"-"`
	EnvApiHost      string            `mapstructure:"ENV_API_HOST,omitempty" json:"ENV_API_HOST"`
	EnvCoreApiHost  string            `mapstructure:"ENV_CORE_API_HOST,omitempty" json:"ENV_CORE_API_HOST"`
	EnvCoreScanHost string            `mapstructure:"ENV_CORE_SCAN_HOST,omitempty" json:"ENV_CORE_SCAN_HOST"`
	EnvRpcServer    string            `mapstructure:"ENV_RPC_SERVER,omitempty" json:"ENV_RPC_SERVER"`
	EnvNetworkID    int               `mapstructure:"ENV_NETWORK_ID,omitempty" json:"ENV_NETWORK_ID"`
	EnvNetworkType  string            `mapstructure:"ENV_NETWORK_TYPE,omitempty" json:"ENV_NETWORK_TYPE"`
	EnvChainType    string            `mapstructure:"ENV_CHAIN_TYPE,omitempty" json:"ENV_CHAIN_TYPE"`
	EnvWalletConfig WalletConfig      `mapstructure:"ENV_WALLET_CONFIG,omitempty" json:"ENV_WALLET_CONFIG"`
	EnvLogo         string            `mapstructure:"ENV_LOGO,omitempty" json:"ENV_LOGO"`
	EnvTheme        ThemeConfig       `mapstructure:"ENV_THEME,omitempty" json:"ENV_THEME"`
	EnvIcons        IconsConfig       `mapstructure:"ENV_ICONS,omitempty" json:"ENV_ICONS"`
	EnvLocalesEn    map[string]string `mapstructure:"ENV_LOCALES_EN,omitempty" json:"ENV_LOCALES_EN"`
	EnvLocalesCn    map[string]string `mapstructure:"ENV_LOCALES_CN,omitempty" json:"ENV_LOCALES_CN"`
}

type Frontend struct {
	Type              string            `mapstructure:"type"`
	SiriusRepo        string            `mapstructure:"sirius_repo,omitempty"`
	SiriusEthRepo     string            `mapstructure:"sirius_eth_repo,omitempty"`
	PrebuiltRepo      string            `mapstructure:"prebuilt_repo,omitempty"`
	CoreSpaceSettings CoreSpaceSettings `mapstructure:"core_space_settings,omitempty"`
	ESpaceSettings    ESpaceSettings    `mapstructure:"e_space_settings,omitempty"`
}

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
