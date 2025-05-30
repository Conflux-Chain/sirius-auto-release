package config

type Label struct {
	EN string `toml:"en"`
	ZH string `toml:"zh"`
}

type PromptOption struct {
	Value string `toml:"value"`
	Label Label  `toml:"label"`
}
type Title struct {
	EN string `toml:"en"`
	ZH string `toml:"zh"`
}

type Description struct {
	EN string `toml:"en"`
	ZH string `toml:"zh"`
}

type PromptConfig struct {
	Title       Title          `toml:"title"`
	Description Description    `toml:"description"`
	Type        string         `toml:"type"`
	Options     []PromptOption `toml:"options"`
}

type StringConfigItem struct {
	Value  string       `toml:"value"`
	Prompt PromptConfig `toml:"_prompt"`
}

type IntConfigItem struct {
	Value  int          `toml:"value"`
	Prompt PromptConfig `toml:"_prompt"`
}

type BoolConfigItem struct {
	Value  bool         `toml:"value"`
	Prompt PromptConfig `toml:"_prompt"`
}

type StringSliceConfigItem struct {
	Value  []string     `toml:"value"`
	Prompt PromptConfig `toml:"_prompt"`
}

type DataGlobalConfig struct {
	Version int              `toml:"version"`
	Space   StringConfigItem `toml:"space"`
	Workdir StringConfigItem `toml:"workdir"`
}

type DataCoreSpaceSettingsConfig struct {
	API_URL                      StringConfigItem `toml:"api_url"`
	EnvOpenApiHost               StringConfigItem `toml:"ENV_OPEN_API_HOST"`
	EnvRpcServer                 StringConfigItem `toml:"ENV_RPC_SERVER"`
	EnvNetworkID                 IntConfigItem    `toml:"ENV_NETWORK_ID"`
	EnvChainType                 StringConfigItem `toml:"ENV_CHAIN_TYPE"`
	EnvFcAddress                 StringConfigItem `toml:"ENV_FC_ADDRESS"`
	EnvFcExchangeAddress         StringConfigItem `toml:"ENV_FC_EXCHANGE_ADDRESS"`
	EnvFcExchangeInterestAddress StringConfigItem `toml:"ENV_FC_EXCHANGE_INTEREST_ADDRESS"`
	EnvEnsRegistryAddress        StringConfigItem `toml:"ENV_ENS_REGISTRY_ADDRESS"`
	EnvEnsPublicResolverAddress  StringConfigItem `toml:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS"`
	EnvLogo                      StringConfigItem `toml:"ENV_LOGO"`
}
type DataNativeCurrencyConfig struct {
	Name     StringConfigItem `toml:"name"`
	Symbol   StringConfigItem `toml:"symbol"`
	Decimals IntConfigItem    `toml:"decimals"`
}

type DataWalletConfig struct {
	Enable            BoolConfigItem           `toml:"enable"`
	ChainID           IntConfigItem            `toml:"chainId"`
	ChainName         StringConfigItem         `toml:"chainName"`
	RpcUrls           StringSliceConfigItem    `toml:"rpcUrls"`
	BlockExplorerUrls StringSliceConfigItem    `toml:"blockExplorerUrls"`
	NativeCurrency    DataNativeCurrencyConfig `toml:"nativeCurrency"`
}
type DataESpaceSettingsConfig struct {
	API_URL         StringConfigItem `toml:"api_url"`
	EnvOpenApiHost  StringConfigItem `toml:"ENV_OPEN_API_HOST"`
	EnvCoreApiHost  StringConfigItem `toml:"ENV_CORE_API_HOST"`
	EnvCoreScanHost StringConfigItem `toml:"ENV_CORE_SCAN_HOST"`
	EnvRpcServer    StringConfigItem `toml:"ENV_RPC_SERVER"`
	EnvNetworkID    IntConfigItem    `toml:"ENV_NETWORK_ID"`
	EnvChainType    StringConfigItem `toml:"ENV_CHAIN_TYPE"`
	EnvWalletConfig DataWalletConfig `toml:"ENV_WALLET_CONFIG"`
	EnvLogo         StringConfigItem `toml:"ENV_LOGO"`
}
type DataFrontendConfig struct {
	Type         StringConfigItem `toml:"type"`
	PrebuiltRepo StringConfigItem `toml:"prebuilt_repo"`
	// SiriusRepo        StringConfigItem      `toml:"sirius_repo"` // Uncomment if needed
	// SiriusEthRepo     StringConfigItem      `toml:"sirius_eth_repo"` // Uncomment if needed
	CoreSpaceSettings DataCoreSpaceSettingsConfig `toml:"core_space_settings"`
	ESpaceSettings    DataESpaceSettingsConfig    `toml:"e_space_settings"`
}

type ContainerSpacePortConfig struct {
	Port IntConfigItem `toml:"port"`
}

type DataContainerConfig struct {
	Enabled   BoolConfigItem           `toml:"enabled"`
	Type      StringConfigItem         `toml:"type"`
	Name      StringConfigItem         `toml:"name"`
	Tag       StringConfigItem         `toml:"tag"`
	CoreSpace ContainerSpacePortConfig `toml:"core_space"`
	ESpace    ContainerSpacePortConfig `toml:"e_space"`
}

type ProxySpaceConfig struct {
	Port IntConfigItem `toml:"port"`
}

type DataProxyConfig struct {
	CoreSpace ProxySpaceConfig `toml:"core_space"`
	ESpace    ProxySpaceConfig `toml:"e_space"`
}

type DataConfig struct {
	Global    DataGlobalConfig    `toml:"global"`
	Frontend  DataFrontendConfig  `toml:"frontend"`
	Container DataContainerConfig `toml:"container"`
	Proxy     DataProxyConfig     `toml:"proxy"`
}
