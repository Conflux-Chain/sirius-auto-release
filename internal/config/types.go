package config

import "Conflux-Chain/sirius-auto-release/internal/i18n"

type WalletConfig struct {
	ChainID           int      `mapstructure:"chainId"`
	ChainName         string   `mapstructure:"chainName" huh_en:"Please input a chain name" huh_zh:"请输入链名称"`
	RpcUrls           []string `mapstructure:"rpcUrls" huh_en:"Please input a rpc urls" huh_zh:"请输入 rpc 地址"`
	BlockExplorerUrls []string `mapstructure:"blockExplorerUrls" huh_en:"Please input a block explorer urls" huh_zh:"请输入区块浏览器地址"`
	NativeCurrency    struct {
		Name     string `mapstructure:"name" huh_en:"Please input a native currency name" huh_zh:"请输入原生货币名称"`
		Symbol   string `mapstructure:"symbol" huh_en:"Please input a native currency symbol" huh_zh:"请输入本地货币符号"`
		Decimals int    `mapstructure:"decimals" huh_en:"Please input a native currency decimals" huh_zh:"请输入本地货币小数位数"`
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
	Enabled                      bool   `mapstructure:"enabled,omitempty" json:"-" huh_en:"Enable Core Space Settings" huh_zh:"是否启用 Core Space 设置"`
	EnvOpenApiHost               string `mapstructure:"ENV_OPEN_API_HOST,omitempty" json:"ENV_OPEN_API_HOST" huh_en:"Please input a Open API host" huh_zh:"请输入 Open API 地址"`
	EnvRpcServer                 string `mapstructure:"ENV_RPC_SERVER,omitempty" json:"ENV_RPC_SERVER" huh_en:"Please input a RPC server" huh_zh:"请输入 RPC 服务器"`
	EnvNetworkID                 int    `mapstructure:"ENV_NETWORK_ID,omitempty" json:"ENV_NETWORK_ID" huh_en:"Please input a Network ID" huh_zh:"请输入网络 ID"`
	EnvNetworkType               string `mapstructure:"ENV_NETWORK_TYPE,omitempty" json:"ENV_NETWORK_TYPE" huh_en:"Please input a Network Type" huh_zh:"请输入网络类型"`
	EnvChainType                 string `mapstructure:"ENV_CHAIN_TYPE,omitempty" json:"ENV_CHAIN_TYPE" huh_en:"Please select a Chain Type" huh_zh:"请选择链类型"`
	EnvFcAddress                 string `mapstructure:"ENV_FC_ADDRESS,omitempty" json:"ENV_FC_ADDRESS" huh_en:"Please input a FC address" huh_zh:"请输入 FC 地址"`
	EnvFcExchangeAddress         string `mapstructure:"ENV_FC_EXCHANGE_ADDRESS,omitempty" json:"ENV_FC_EXCHANGE_ADDRESS" huh_en:"Please input a FC exchange address" huh_zh:"请输入 FC 交易所地址"`
	EnvFcExchangeInterestAddress string `mapstructure:"ENV_FC_EXCHANGE_INTEREST_ADDRESS,omitempty" json:"ENV_FC_EXCHANGE_INTEREST_ADDRESS" huh_en:"Please input a FC exchange interest address" huh_zh:"请输入 FC 交易所利息地址"`
	EnvEnsRegistryAddress        string `mapstructure:"ENV_ENS_REGISTRY_ADDRESS,omitempty" json:"ENV_ENS_REGISTRY_ADDRESS" huh_en:"Please input a ENS registry address" huh_zh:"请输入 ENS 注册表地址"`
	EnvEnsPublicResolverAddress  string `mapstructure:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS,omitempty" json:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS" huh_en:"Please input a ENS public resolver address" huh_zh:"请输入 ENS 公共解析器地址"`
	EnvLogo                      string `mapstructure:"ENV_LOGO,omitempty" json:"ENV_LOGO" huh_en:"Please input a logo" huh_zh:"请输入 logo"`
}

type ESpaceSettings struct {
	Enabled         bool              `mapstructure:"enabled,omitempty" json:"-" huh_en:"Enable E Space Settings" huh_zh:"是否启用 E Space 设置"`
	EnvApiHost      string            `mapstructure:"ENV_API_HOST,omitempty" json:"ENV_API_HOST" huh_en:"Please input a API host" huh_zh:"请输入 API 地址"`
	EnvCoreApiHost  string            `mapstructure:"ENV_CORE_API_HOST,omitempty" json:"ENV_CORE_API_HOST" huh_en:"Please input a Core API host" huh_zh:"请输入 Core API 地址"`
	EnvCoreScanHost string            `mapstructure:"ENV_CORE_SCAN_HOST,omitempty" json:"ENV_CORE_SCAN_HOST" huh_en:"Please input a Core Scan host" huh_zh:"请输入 Core 扫描地址"`
	EnvRpcServer    string            `mapstructure:"ENV_RPC_SERVER,omitempty" json:"ENV_RPC_SERVER" huh_en:"Please input a RPC server" huh_zh:"请输入 RPC 地址"`
	EnvNetworkID    int               `mapstructure:"ENV_NETWORK_ID,omitempty" json:"ENV_NETWORK_ID" huh_en:"Please input a Network ID" huh_zh:"请输入网络 ID"`
	EnvNetworkType  string            `mapstructure:"ENV_NETWORK_TYPE,omitempty" json:"ENV_NETWORK_TYPE" `
	EnvChainType    string            `mapstructure:"ENV_CHAIN_TYPE,omitempty" json:"ENV_CHAIN_TYPE" huh_en:"Please select a Chain Type" huh_zh:"请选择链类型"`
	EnvWalletConfig WalletConfig      `mapstructure:"ENV_WALLET_CONFIG,omitempty" json:"ENV_WALLET_CONFIG"`
	EnvLogo         string            `mapstructure:"ENV_LOGO,omitempty" json:"ENV_LOGO" huh_en:"Please input a logo" huh_zh:"请输入 logo"`
	EnvTheme        ThemeConfig       `mapstructure:"ENV_THEME,omitempty" json:"ENV_THEME"`
	EnvIcons        IconsConfig       `mapstructure:"ENV_ICONS,omitempty" json:"ENV_ICONS"`
	EnvLocalesEn    map[string]string `mapstructure:"ENV_LOCALES_EN,omitempty" json:"ENV_LOCALES_EN"`
	EnvLocalesCn    map[string]string `mapstructure:"ENV_LOCALES_CN,omitempty" json:"ENV_LOCALES_CN"`
}

type Frontend struct {
	Type string `mapstructure:"type" huh_en:"Please select a frontend type" huh_zh:"请选择前端类型"`
	// SiriusRepo        string            `mapstructure:"sirius_repo,omitempty"`
	// SiriusEthRepo     string            `mapstructure:"sirius_eth_repo,omitempty"`
	PrebuiltRepo      string            `mapstructure:"prebuilt_repo,omitempty" huh_en:"Please input a prebuilt repo" huh_zh:"请输入预构建的仓库"`
	CoreSpaceSettings CoreSpaceSettings `mapstructure:"core_space_settings,omitempty"`
	ESpaceSettings    ESpaceSettings    `mapstructure:"e_space_settings,omitempty"`
}

type Global struct {
	Version int    `mapstructure:"version"`
	Space   string `mapstructure:"space" huh_en:"Please select a Space" huh_zh:"请选择要使用的 Space"`
	Workdir string `mapstructure:"workdir" huh_en:"Please select a workdir" huh_zh:"请选择工作目录"`
}

type ProxySpace struct {
	Port    int    `mapstructure:"port" huh_en:"Please input a port" huh_zh:"请输入端口号"`
	API_URL string `mapstructure:"api_url" huh_en:"Please input a proxy API URL" huh_zh:"请输入代理 API 地址"`
}

type Proxy struct {
	Enabled   bool       `mapstructure:"enabled" huh_en:"Enable Proxy" huh_zh:"是否启用代理"`
	Type      string     `mapstructure:"type"`
	CoreSpace ProxySpace `mapstructure:"core_space" huh_en:"core space" huh_zh:"core space"`
	ESpace    ProxySpace `mapstructure:"e_space" huh_en:"ESpace" huh_zh:"ESpace"`
}

type ContainerSpace struct {
	Port int `mapstructure:"port" huh_en:"Please input a port" huh_zh:"请输入端口号"`
}

type Container struct {
	Enabled bool   `mapstructure:"enabled" huh_en:"Enable Container" huh_zh:"是否启用容器"`
	Type    string `mapstructure:"type"`
	Name    string `mapstructure:"name" huh_en:"Please input a container name" huh_zh:"请输入容器名称"`
	Tag     string `mapstructure:"tag" huh_en:"Please input a container tag" huh_zh:"请输入容器标签"`

	CoreSpace ContainerSpace `mapstructure:"core_space" huh_en:"core space" huh_zh:"core space"`
	ESpace    ContainerSpace `mapstructure:"e_space" huh_en:"ESpace" huh_zh:"ESpace"`
}

type Config struct {
	Global    Global    `mapstructure:"global"`
	Frontend  Frontend  `mapstructure:"frontend"`
	Container Container `mapstructure:"container"`
	Proxy     Proxy     `mapstructure:"proxy"`
}

func init() {
	i18n.ExtractAllTranslations(Config{})
	i18n.DumpTranslations()
}
