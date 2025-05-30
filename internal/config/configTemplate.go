package config

type WalletConfig struct {
	ChainID           int      `toml:"chainId,omitempty" json:"chainId,omitempty"`
	ChainName         string   `toml:"chainName,omitempty" json:"chainName,omitempty"`
	RpcUrls           []string `toml:"rpcUrls,omitempty" json:"rpcUrls,omitempty"`
	BlockExplorerUrls []string `toml:"blockExplorerUrls,omitempty" json:"blockExplorerUrls,omitempty"`
	NativeCurrency    struct {
		Name     string `toml:"name,omitempty" json:"name,omitempty"`
		Symbol   string `toml:"symbol,omitempty" json:"symbol,omitempty"`
		Decimals int    `toml:"decimals,omitempty" json:"decimals,omitempty"`
	} `toml:"nativeCurrency,omitempty" json:"nativeCurrency"`
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
	API_URL         string       `toml:"api_url,omitempty" json:"-"`
	EnvOpenApiHost  string       `toml:"ENV_OPEN_API_HOST,omitempty" json:"ENV_OPEN_API_HOST,omitempty"`
	EnvCoreApiHost  string       `toml:"ENV_CORE_API_HOST,omitempty" json:"ENV_CORE_API_HOST,omitempty"`
	EnvCoreScanHost string       `toml:"ENV_CORE_SCAN_HOST,omitempty" json:"ENV_CORE_SCAN_HOST,omitempty"`
	EnvRpcServer    string       `toml:"ENV_RPC_SERVER,omitempty" json:"ENV_RPC_SERVER,omitempty"`
	EnvNetworkID    int          `toml:"ENV_NETWORK_ID,omitempty" json:"ENV_NETWORK_ID,omitempty"`
	EnvNetworkType  string       `toml:"ENV_NETWORK_TYPE,omitempty" json:"ENV_NETWORK_TYPE,omitempty"`
	EnvChainType    string       `toml:"ENV_CHAIN_TYPE,omitempty" json:"ENV_CHAIN_TYPE,omitempty"`
	EnvWalletConfig WalletConfig `toml:"ENV_WALLET_CONFIG,omitempty" json:"ENV_WALLET_CONFIG,omitempty"`
	EnvLogo         string       `toml:"ENV_LOGO,omitempty" json:"ENV_LOGO,omitempty"`
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
