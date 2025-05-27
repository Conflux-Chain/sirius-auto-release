package config

import "Conflux-Chain/sirius-auto-release/internal/i18n"

type WalletConfig struct {
	ChainID           int      `toml:"chainId" huh_en:"Please input a chain id" huh_zh:"请输入链 ID"`
	ChainName         string   `toml:"chainName" huh_en:"Please input a chain name" huh_zh:"请输入链名称"`
	RpcUrls           []string `toml:"rpcUrls" huh_en:"Please input a rpc urls" huh_zh:"请输入 rpc 地址"`
	BlockExplorerUrls []string `toml:"blockExplorerUrls" huh_en:"Please input a block explorer urls" huh_zh:"请输入区块浏览器地址"`
	NativeCurrency    struct {
		Name     string `toml:"name" huh_en:"Please input a native currency name" huh_zh:"请输入原生货币名称"`
		Symbol   string `toml:"symbol" huh_en:"Please input a native currency symbol" huh_zh:"请输入本地货币符号"`
		Decimals int    `toml:"decimals" huh_en:"Please input a native currency decimals" huh_zh:"请输入本地货币小数位数"`
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
	API_URL                      string `toml:"api_url" huh_en:"Please input a proxy API URL" huh_zh:"请输入代理 API 地址"`
	Enabled                      bool   `toml:"enabled,omitempty" json:"-" huh_en:"Enable Core Space Settings" huh_zh:"是否启用 Core Space 设置"`
	EnvOpenApiHost               string `toml:"ENV_OPEN_API_HOST,omitempty" toml:"ENV_OPEN_API_HOST" json:"ENV_OPEN_API_HOST" huh_en:"Please input a Open API host" huh_zh:"请输入 Open API 地址"`
	EnvRpcServer                 string `toml:"ENV_RPC_SERVER,omitempty" toml:"ENV_RPC_SERVER" json:"ENV_RPC_SERVER" huh_en:"Please input a RPC server" huh_zh:"请输入 RPC 服务器"`
	EnvNetworkID                 int    `toml:"ENV_NETWORK_ID,omitempty" toml:"ENV_NETWORK_ID" json:"ENV_NETWORK_ID" huh_en:"Please input a Network ID" huh_zh:"请输入网络 ID"`
	EnvNetworkType               string `toml:"ENV_NETWORK_TYPE,omitempty" toml:"ENV_NETWORK_TYPE" json:"ENV_NETWORK_TYPE" huh_en:"Please input a Network Type" huh_zh:"请输入网络类型"`
	EnvChainType                 string `toml:"ENV_CHAIN_TYPE,omitempty" toml:"ENV_CHAIN_TYPE" json:"ENV_CHAIN_TYPE" huh_en:"Please select a Chain Type" huh_zh:"请选择链类型"`
	EnvFcAddress                 string `toml:"ENV_FC_ADDRESS,omitempty" toml:"ENV_FC_ADDRESS" json:"ENV_FC_ADDRESS" huh_en:"Please input a FC address" huh_zh:"请输入 FC 地址"`
	EnvFcExchangeAddress         string `toml:"ENV_FC_EXCHANGE_ADDRESS,omitempty" toml:"ENV_FC_EXCHANGE_ADDRESS" json:"ENV_FC_EXCHANGE_ADDRESS" huh_en:"Please input a FC exchange address" huh_zh:"请输入 FC 交易所地址"`
	EnvFcExchangeInterestAddress string `toml:"ENV_FC_EXCHANGE_INTEREST_ADDRESS,omitempty" toml:"ENV_FC_EXCHANGE_INTEREST_ADDRESS" json:"ENV_FC_EXCHANGE_INTEREST_ADDRESS" huh_en:"Please input a FC exchange interest address" huh_zh:"请输入 FC 交易所利息地址"`
	EnvEnsRegistryAddress        string `toml:"ENV_ENS_REGISTRY_ADDRESS,omitempty" toml:"ENV_ENS_REGISTRY_ADDRESS" json:"ENV_ENS_REGISTRY_ADDRESS" huh_en:"Please input a ENS registry address" huh_zh:"请输入 ENS 注册表地址"`
	EnvEnsPublicResolverAddress  string `toml:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS,omitempty" toml:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS" json:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS" huh_en:"Please input a ENS public resolver address" huh_zh:"请输入 ENS 公共解析器地址"`
	EnvLogo                      string `toml:"ENV_LOGO,omitempty" toml:"ENV_LOGO" json:"ENV_LOGO" huh_en:"Please input a logo" huh_zh:"请输入 logo"`
}

type ESpaceSettings struct {
	API_URL         string            `toml:"api_url" huh_en:"Please input a proxy API URL" huh_zh:"请输入代理 API 地址"`
	Enabled         bool              `toml:"enabled,omitempty" json:"-" huh_en:"Enable E Space Settings" huh_zh:"是否启用 E Space 设置"`
	EnvApiHost      string            `toml:"ENV_API_HOST,omitempty" toml:"ENV_API_HOST" json:"ENV_API_HOST" huh_en:"Please input a API host" huh_zh:"请输入 API 地址"`
	EnvCoreApiHost  string            `toml:"ENV_CORE_API_HOST,omitempty" toml:"ENV_CORE_API_HOST" json:"ENV_CORE_API_HOST" huh_en:"Please input a Core API host" huh_zh:"请输入 Core API 地址"`
	EnvCoreScanHost string            `toml:"ENV_CORE_SCAN_HOST,omitempty" toml:"ENV_CORE_SCAN_HOST" json:"ENV_CORE_SCAN_HOST" huh_en:"Please input a Core Scan host" huh_zh:"请输入 Core 扫描地址"`
	EnvRpcServer    string            `toml:"ENV_RPC_SERVER,omitempty" toml:"ENV_RPC_SERVER" json:"ENV_RPC_SERVER" huh_en:"Please input a RPC server" huh_zh:"请输入 RPC 地址"`
	EnvNetworkID    int               `toml:"ENV_NETWORK_ID,omitempty" toml:"ENV_NETWORK_ID" json:"ENV_NETWORK_ID" huh_en:"Please input a Network ID" huh_zh:"请输入网络 ID"`
	EnvNetworkType  string            `toml:"ENV_NETWORK_TYPE,omitempty" toml:"ENV_NETWORK_TYPE" json:"ENV_NETWORK_TYPE" huh_en:"Please input a Network Type" huh_zh:"请输入网络类型"`
	EnvChainType    string            `toml:"ENV_CHAIN_TYPE,omitempty" toml:"ENV_CHAIN_TYPE" json:"ENV_CHAIN_TYPE" huh_en:"Please select a Chain Type" huh_zh:"请选择链类型"`
	EnvWalletConfig WalletConfig      `toml:"ENV_WALLET_CONFIG,omitempty" json:"ENV_WALLET_CONFIG"`
	EnvLogo         string            `toml:"ENV_LOGO,omitempty" toml:"ENV_LOGO" json:"ENV_LOGO" huh_en:"Please input a logo" huh_zh:"请输入 logo"`
	EnvTheme        ThemeConfig       `toml:"ENV_THEME,omitempty" toml:"ENV_THEME" json:"ENV_THEME"`
	EnvIcons        IconsConfig       `toml:"ENV_ICONS,omitempty" toml:"ENV_ICONS" json:"ENV_ICONS"`
	EnvLocalesEn    map[string]string `toml:"ENV_LOCALES_EN,omitempty" toml:"ENV_LOCALES_EN" json:"ENV_LOCALES_EN"`
	EnvLocalesCn    map[string]string `toml:"ENV_LOCALES_CN,omitempty" toml:"ENV_LOCALES_CN" json:"ENV_LOCALES_CN"`
}

type Frontend struct {
	Type string `toml:"type" huh_en:"Please select a frontend type" huh_zh:"请选择前端类型"`
	// SiriusRepo        string            `toml:"sirius_repo,omitempty"`
	// SiriusEthRepo     string            `toml:"sirius_eth_repo,omitempty"`
	PrebuiltRepo      string            `toml:"prebuilt_repo" huh_en:"Please input a prebuilt repo" huh_zh:"请输入预构建的仓库"`
	CoreSpaceSettings CoreSpaceSettings `toml:"core_space_settings"`
	ESpaceSettings    ESpaceSettings    `toml:"e_space_settings"`
}

type Global struct {
	Version int    `toml:"version"`
	Space   string `toml:"space" huh_en:"Please select a Space" huh_zh:"请选择要使用的 Space"`
	Workdir string ` toml:"workdir" huh_en:"Please select a workdir" huh_zh:"请选择工作目录"`
}

type ProxySpace struct {
	Port int `toml:"port" huh_en:"Please input a port" huh_zh:"请输入端口号"`
}

type Proxy struct {
	Enabled   bool       `toml:"enabled" huh_en:"Enable Proxy" huh_zh:"是否启用代理"`
	Type      string     `toml:"type"`
	CoreSpace ProxySpace `toml:"core_space" huh_en:"core space" huh_zh:"core space"`
	ESpace    ProxySpace `toml:"e_space" huh_en:"ESpace" huh_zh:"ESpace"`
}

type ContainerSpace struct {
	Port int `toml:"port" huh_en:"Please input a port" huh_zh:"请输入端口号"`
}

type Container struct {
	Enabled bool   `toml:"enabled" huh_en:"Enable Container" huh_zh:"是否启用容器"`
	Type    string `toml:"type"`
	Name    string `toml:"name" huh_en:"Please input a container name" huh_zh:"请输入容器名称"`
	Tag     string `toml:"tag" huh_en:"Please input a container tag" huh_zh:"请输入容器标签"`

	CoreSpace ContainerSpace `toml:"core_space" huh_en:"core space" huh_zh:"core space"`
	ESpace    ContainerSpace `toml:"e_space" huh_en:"ESpace" huh_zh:"ESpace"`
}

type Config struct {
	Global    Global    `toml:"global"`
	Frontend  Frontend  `toml:"frontend"`
	Proxy     Proxy     `toml:"proxy"`
	Container Container `toml:"container"`
}

func init() {
	i18n.ExtractAllTranslations(Config{})
}
