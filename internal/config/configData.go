package config

type PromptOption struct {
	Value   string `toml:"value"`
	LabelEN string `toml:"label.en"`
	LabelZH string `toml:"label.zh"`
}

type PromptConfig struct {
	TitleEN       string         `toml:"title.en"`
	TitleZH       string         `toml:"title.zh"`
	DescriptionEN string         `toml:"description.en,omitempty"`
	DescriptionZH string         `toml:"description.zh,omitempty"`
	Type          string         `toml:"type"`
	Options       []PromptOption `toml:"options,omitempty"`
}

type StringConfigItem struct {
	Value  string       `toml:"value"`
	Prompt PromptConfig `toml:"_prompt,omitempty"`
}

type IntConfigItem struct {
	Value  int          `toml:"value"`
	Prompt PromptConfig `toml:"_prompt,omitempty"`
}

type BoolConfigItem struct {
	Value  bool         `toml:"value"`
	Prompt PromptConfig `toml:"_prompt,omitempty"`
}

type StringSliceConfigItem struct {
	Value  []string     `toml:"value"`
	Prompt PromptConfig `toml:"_prompt,omitempty"`
}

type DataGlobalConfig struct {
	Version int              `toml:"version"` // Direct value, no _prompt
	Space   StringConfigItem `toml:"space"`
	Workdir StringConfigItem `toml:"workdir"`
}

type CoreSpaceSettingsConfig struct {
	ApiURL                           StringConfigItem `toml:"api_url"`
	Enable                           BoolConfigItem   `toml:"enable"`
	ENV_OPEN_API_HOST                StringConfigItem `toml:"ENV_OPEN_API_HOST"`
	ENV_RPC_SERVER                   StringConfigItem `toml:"ENV_RPC_SERVER"`
	ENV_NETWORK_ID                   IntConfigItem    `toml:"ENV_NETWORK_ID"`
	ENV_CHAIN_TYPE                   StringConfigItem `toml:"ENV_CHAIN_TYPE"`
	ENV_FC_ADDRESS                   StringConfigItem `toml:"ENV_FC_ADDRESS"`
	ENV_FC_EXCHANGE_ADDRESS          StringConfigItem `toml:"ENV_FC_EXCHANGE_ADDRESS"`
	ENV_FC_EXCHANGE_INTEREST_ADDRESS StringConfigItem `toml:"ENV_FC_EXCHANGE_INTEREST_ADDRESS"`
	ENV_ENS_REGISTRY_ADDRESS         StringConfigItem `toml:"ENV_ENS_REGISTRY_ADDRESS"`
	ENV_ENS_PUBLIC_RESOLVER_ADDRESS  StringConfigItem `toml:"ENV_ENS_PUBLIC_RESOLVER_ADDRESS"`
	ENV_LOGO                         StringConfigItem `toml:"ENV_LOGO"`
}
type NativeCurrencyConfig struct {
	Name     StringConfigItem `toml:"name"`
	Symbol   StringConfigItem `toml:"symbol"`
	Decimals IntConfigItem    `toml:"decimals"`
}

type ESpaceWalletConfig struct {
	Enable            BoolConfigItem        `toml:"enable"`
	ChainId           IntConfigItem         `toml:"chainId"`
	ChainName         StringConfigItem      `toml:"chainName"`
	RpcUrls           StringSliceConfigItem `toml:"rpcUrls"`
	BlockExplorerUrls StringSliceConfigItem `toml:"blockExplorerUrls"`
	NativeCurrency    NativeCurrencyConfig  `toml:"nativeCurrency"`
}
type ESpaceSettingsConfig struct {
	ApiURL             StringConfigItem   `toml:"api_url"`
	Enable             BoolConfigItem     `toml:"enable"`
	ENV_API_HOST       StringConfigItem   `toml:"ENV_API_HOST"`
	ENV_CORE_API_HOST  StringConfigItem   `toml:"ENV_CORE_API_HOST"`
	ENV_CORE_SCAN_HOST StringConfigItem   `toml:"ENV_CORE_SCAN_HOST"`
	ENV_RPC_SERVER     StringConfigItem   `toml:"ENV_RPC_SERVER"`
	ENV_NETWORK_ID     IntConfigItem      `toml:"ENV_NETWORK_ID"`
	ENV_CHAIN_TYPE     StringConfigItem   `toml:"ENV_CHAIN_TYPE"`
	ENV_LOGO           StringConfigItem   `toml:"ENV_LOGO"`
	ENV_WALLET_CONFIG  ESpaceWalletConfig `toml:"ENV_WALLET_CONFIG"`
}
type DataFrontendConfig struct {
	Type         StringConfigItem `toml:"type"`
	PrebuiltRepo StringConfigItem `toml:"prebuilt_repo"`
	// SiriusRepo        StringConfigItem      `toml:"sirius_repo,omitempty"` // Uncomment if needed
	// SiriusEthRepo     StringConfigItem      `toml:"sirius_eth_repo,omitempty"` // Uncomment if needed
	CoreSpaceSettings CoreSpaceSettingsConfig `toml:"core_space_settings"`
	ESpaceSettings    ESpaceSettingsConfig    `toml:"e_space_settings"`
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
