package interactive

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

func runGlobalForm(cfg *config.Global, DataGlobal config.DataGlobalConfig, language string) error {

	cfg.Workdir = "./frontend"

	var globalFields []huh.Field

	cfg.Space = config.ALL_SPACE

	SpaceField, err := CreateHuhField("Space", &DataGlobal.Space.Prompt, &cfg.Space, language)
	if err != nil {
		return fmt.Errorf("failed to create field for Space: %w", err)
	}
	globalFields = append(globalFields, SpaceField)

	cfg.Workdir = "./frontend"
	WorkdirField, err := CreateHuhField("Workdir", &DataGlobal.Workdir.Prompt, &cfg.Workdir, language)
	if err != nil {
		return fmt.Errorf("failed to create field for Workdir: %w", err)
	}
	globalFields = append(globalFields, WorkdirField)

	form := huh.NewForm(
		huh.NewGroup(
			globalFields...,
		),
	)

	return form.Run()

}

func runFrontendForm(cfg *config.Frontend, globalConfig *config.Global, DataFrontend config.DataFrontendConfig, language string) error {

	cfg.Type = "prebuilt"
	cfg.PrebuiltRepo = "https://github.com/Conflux-Chain/sirius-auto-release.git"

	if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.CORE_SPACE {
		cfg.CoreSpaceSettings.EnvNetworkType = "CORE"
		cfg.CoreSpaceSettings.EnvChainType = "DEVNET"

		coreFields := []huh.Field{}

		API_URL_FIELD, err := CreateHuhField("API_URL", &DataFrontend.CoreSpaceSettings.API_URL.Prompt, &cfg.CoreSpaceSettings.API_URL, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace API URL: %w", err)
		}
		coreFields = append(coreFields, API_URL_FIELD)

		ENV_OPEN_API_HOST_FIELD, err := CreateHuhField("ENV_OPEN_API_HOST", &DataFrontend.CoreSpaceSettings.EnvOpenApiHost.Prompt, &cfg.CoreSpaceSettings.EnvOpenApiHost, language)

		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_OPEN_API_HOST: %w", err)
		}
		coreFields = append(coreFields, ENV_OPEN_API_HOST_FIELD)

		EnvRpcServerField, err := CreateHuhField("ENV_RPC_SERVER", &DataFrontend.CoreSpaceSettings.EnvRpcServer.Prompt, &cfg.CoreSpaceSettings.EnvRpcServer, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_RPC_SERVER: %w", err)
		}
		coreFields = append(coreFields, EnvRpcServerField)

		EnvNetworkIDField, err := CreateHuhField("ENV_NETWORK_ID", &DataFrontend.CoreSpaceSettings.EnvNetworkID.Prompt, &cfg.CoreSpaceSettings.EnvNetworkID, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_NETWORK_ID: %w", err)
		}
		coreFields = append(coreFields, EnvNetworkIDField)

		EnvFcAddressField, err := CreateHuhField("ENV_FC_ADDRESS", &DataFrontend.CoreSpaceSettings.EnvFcAddress.Prompt, &cfg.CoreSpaceSettings.EnvFcAddress, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_FC_ADDRESS: %w", err)
		}
		coreFields = append(coreFields, EnvFcAddressField)

		EnvFcExchangeAddressField, err := CreateHuhField("ENV_FC_EXCHANGE_ADDRESS", &DataFrontend.CoreSpaceSettings.EnvFcExchangeAddress.Prompt, &cfg.CoreSpaceSettings.EnvFcExchangeAddress, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_FC_EXCHANGE_ADDRESS: %w", err)
		}
		coreFields = append(coreFields, EnvFcExchangeAddressField)

		EnvFcExchangeInterestAddressField, err := CreateHuhField("ENV_FC_EXCHANGE_INTEREST_ADDRESS", &DataFrontend.CoreSpaceSettings.EnvFcExchangeInterestAddress.Prompt, &cfg.CoreSpaceSettings.EnvFcExchangeInterestAddress, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_FC_EXCHANGE_INTEREST_ADDRESS: %w", err)
		}

		coreFields = append(coreFields, EnvFcExchangeInterestAddressField)

		EnvEnsRegistryAddressField, err := CreateHuhField("ENV_ENS_REGISTRY_ADDRESS", &DataFrontend.CoreSpaceSettings.EnvEnsRegistryAddress.Prompt, &cfg.CoreSpaceSettings.EnvEnsRegistryAddress, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_ENS_REGISTRY_ADDRESS: %w", err)
		}
		coreFields = append(coreFields, EnvEnsRegistryAddressField)

		EnvEnsPublicResolverAddressField, err := CreateHuhField("ENV_ENS_PUBLIC_RESOLVER_ADDRESS", &DataFrontend.CoreSpaceSettings.EnvEnsPublicResolverAddress.Prompt, &cfg.CoreSpaceSettings.EnvEnsPublicResolverAddress, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_ENS_PUBLIC_RESOLVER_ADDRESS: %w", err)
		}
		coreFields = append(coreFields, EnvEnsPublicResolverAddressField)

		EnvLogoField, err := CreateHuhField("ENV_LOGO", &DataFrontend.CoreSpaceSettings.EnvLogo.Prompt, &cfg.CoreSpaceSettings.EnvLogo, language)
		if err != nil {
			return fmt.Errorf("failed to create field for CoreSpace ENV_LOGO: %w", err)
		}
		coreFields = append(coreFields, EnvLogoField)

		if err := huh.NewForm(
			huh.NewGroup(coreFields...),
		).Run(); err != nil {
			return err
		}
	}
	if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.E_SPACE {

		var eSpaceFields []huh.Field
		cfg.ESpaceSettings.EnvNetworkType = "EVM"
		cfg.ESpaceSettings.EnvChainType = "DEVNET"

		API_URL_FIELD, err := CreateHuhField("API_URL", &DataFrontend.ESpaceSettings.API_URL.Prompt, &cfg.ESpaceSettings.API_URL, language)
		if err != nil {
			return fmt.Errorf("failed to create field for ESpace API URL: %w", err)
		}

		eSpaceFields = append(eSpaceFields, API_URL_FIELD)

		EnvOpenApiHostField, err := CreateHuhField("ENV_OPEN_API_HOST", &DataFrontend.ESpaceSettings.EnvOpenApiHost.Prompt, &cfg.ESpaceSettings.EnvOpenApiHost, language)

		if err != nil {
			return fmt.Errorf("failed to create field for ESpace ENV_OPEN_API_HOST: %w", err)
		}
		eSpaceFields = append(eSpaceFields, EnvOpenApiHostField)

		EnvRpcServerField, err := CreateHuhField("ENV_RPC_SERVER", &DataFrontend.ESpaceSettings.EnvRpcServer.Prompt, &cfg.ESpaceSettings.EnvRpcServer, language)
		if err != nil {
			return fmt.Errorf("failed to create field for ESpace ENV_RPC_SERVER: %w", err)
		}
		eSpaceFields = append(eSpaceFields, EnvRpcServerField)

		EnvNetworkIDField, err := CreateHuhField("ENV_NETWORK_ID", &DataFrontend.ESpaceSettings.EnvNetworkID.Prompt, &cfg.ESpaceSettings.EnvNetworkID, language)
		if err != nil {
			return fmt.Errorf("failed to create field for ESpace ENV_NETWORK_ID: %w", err)
		}
		eSpaceFields = append(eSpaceFields, EnvNetworkIDField)

		EnvLogoField, err := CreateHuhField("ENV_LOGO", &DataFrontend.ESpaceSettings.EnvLogo.Prompt, &cfg.ESpaceSettings.EnvLogo, language)
		if err != nil {
			return fmt.Errorf("failed to create field for ESpace ENV_LOGO: %w", err)
		}
		eSpaceFields = append(eSpaceFields, EnvLogoField)

		EnvCoreApiHostField, err := CreateHuhField("ENV_CORE_API_HOST", &DataFrontend.ESpaceSettings.EnvCoreApiHost.Prompt, &cfg.ESpaceSettings.EnvCoreApiHost, language)
		if err != nil {
			return fmt.Errorf("failed to create field for ESpace ENV_CORE_API_HOST: %w", err)
		}

		eSpaceFields = append(eSpaceFields, EnvCoreApiHostField)

		EnvCoreScanHostField, err := CreateHuhField("ENV_CORE_SCAN_HOST", &DataFrontend.ESpaceSettings.EnvCoreScanHost.Prompt, &cfg.ESpaceSettings.EnvCoreScanHost, language)
		if err != nil {
			return fmt.Errorf("failed to create field for ESpace ENV_CORE_SCAN_HOST: %w", err)
		}
		eSpaceFields = append(eSpaceFields, EnvCoreScanHostField)

		formErr := huh.NewForm(
			huh.NewGroup(
				eSpaceFields...,
			),
		).Run()

		if formErr != nil {
			return formErr
		}

		enableWalletConfig := false

		EnvWalletConfigEnabledField, err := CreateHuhField("ENV_WALLET_CONFIG_ENABLE", &DataFrontend.ESpaceSettings.EnvWalletConfig.Enable.Prompt, &enableWalletConfig, language)
		if err != nil {
			return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_ENABLE: %w", err)
		}
		if err := EnvWalletConfigEnabledField.Run(); err != nil {
			return err
		}

		if enableWalletConfig {

			var walletConfigFields []huh.Field

			rpcUrls := ""
			blockExplorerUrls := ""
			chainIDField, err := CreateHuhField("ENV_WALLET_CONFIG_CHAIN_ID", &DataFrontend.ESpaceSettings.EnvWalletConfig.ChainID.Prompt, &cfg.ESpaceSettings.EnvWalletConfig.ChainID, language)
			if err != nil {
				return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_CHAIN_ID: %w", err)
			}

			walletConfigFields = append(walletConfigFields, chainIDField)

			chainNameField, err := CreateHuhField("ENV_WALLET_CONFIG_CHAIN_NAME", &DataFrontend.ESpaceSettings.EnvWalletConfig.ChainName.Prompt, &cfg.ESpaceSettings.EnvWalletConfig.ChainName, language)
			if err != nil {
				return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_CHAIN_NAME: %w", err)
			}
			walletConfigFields = append(walletConfigFields, chainNameField)

			rpcUrlsField, err := CreateHuhField("ENV_WALLET_CONFIG_RPC_URLS", &DataFrontend.ESpaceSettings.EnvWalletConfig.RpcUrls.Prompt, &rpcUrls, language)
			if err != nil {
				return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_RPC_URLS: %w", err)
			}
			walletConfigFields = append(walletConfigFields, rpcUrlsField)

			blockExplorerUrlsField, err := CreateHuhField("ENV_WALLET_CONFIG_BLOCK_EXPLORER_URLS", &DataFrontend.ESpaceSettings.EnvWalletConfig.BlockExplorerUrls.Prompt, &blockExplorerUrls, language)
			if err != nil {
				return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_BLOCK_EXPLORER_URLS: %w", err)
			}
			walletConfigFields = append(walletConfigFields, blockExplorerUrlsField)

			nativeCurrencyNameField, err := CreateHuhField("ENV_WALLET_CONFIG_NATIVE_CURRENCY_NAME", &DataFrontend.ESpaceSettings.EnvWalletConfig.NativeCurrency.Name.Prompt, &cfg.ESpaceSettings.EnvWalletConfig.NativeCurrency.Name, language)
			if err != nil {
				return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_NATIVE_CURRENCY_NAME: %w", err)
			}
			walletConfigFields = append(walletConfigFields, nativeCurrencyNameField)

			symbolField, err := CreateHuhField("ENV_WALLET_CONFIG_NATIVE_CURRENCY_SYMBOL", &DataFrontend.ESpaceSettings.EnvWalletConfig.NativeCurrency.Symbol.Prompt, &cfg.ESpaceSettings.EnvWalletConfig.NativeCurrency.Symbol, language)
			if err != nil {
				return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_NATIVE_CURRENCY_SYMBOL: %w", err)
			}
			walletConfigFields = append(walletConfigFields, symbolField)

			decimalsField, err := CreateHuhField("ENV_WALLET_CONFIG_NATIVE_CURRENCY_DECIMALS", &DataFrontend.ESpaceSettings.EnvWalletConfig.NativeCurrency.Decimals.Prompt, &cfg.ESpaceSettings.EnvWalletConfig.NativeCurrency.Decimals, language)
			if err != nil {
				return fmt.Errorf("failed to create field for ESpace ENV_WALLET_CONFIG_NATIVE_CURRENCY_DECIMALS: %w", err)
			}
			walletConfigFields = append(walletConfigFields, decimalsField)
			if err := huh.NewForm(
				huh.NewGroup(
					walletConfigFields...,
				),
			).Run(); err != nil {
				return err
			}

			cfg.ESpaceSettings.EnvWalletConfig.RpcUrls = strings.Split(rpcUrls, ",")
			cfg.ESpaceSettings.EnvWalletConfig.BlockExplorerUrls = strings.Split(blockExplorerUrls, ",")

		}

	}

	return nil

}

func runProxyForm(cfg *config.Proxy, globalConfig *config.Global, dataConfig config.DataProxyConfig, language string) error {

	cfg.Type = "nginx"

	var proxyFields []huh.Field

	if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.CORE_SPACE {

		coreSpacePortField, err := CreateHuhField("CoreSpacePort", &dataConfig.CoreSpace.Port.Prompt, &cfg.CoreSpace.Port, language)
		if err != nil {
			return fmt.Errorf("failed to create field for Proxy CoreSpacePort: %w", err)
		}
		proxyFields = append(proxyFields, coreSpacePortField)

	}

	if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.E_SPACE {
		eSpacePortField, err := CreateHuhField("ESpacePort", &dataConfig.ESpace.Port.Prompt, &cfg.ESpace.Port, language)
		if err != nil {
			return fmt.Errorf("failed to create field for Proxy ESpacePort: %w", err)
		}
		proxyFields = append(proxyFields, eSpacePortField)
	}

	if err := huh.NewForm(
		huh.NewGroup(
			proxyFields...,
		),
	).Run(); err != nil {
		return err
	}

	return nil
}

func runContainerForm(cfg *config.Container, globalConfig *config.Global, dataConfig config.DataContainerConfig, language string) error {

	enableField, err := CreateHuhField("Enabled", &dataConfig.Enabled.Prompt, &cfg.Enabled, language)
	if err != nil {
		return fmt.Errorf("failed to create field for Container Enabled: %w", err)
	}
	if err := enableField.Run(); err != nil {
		return err
	}

	if cfg.Enabled {

		var containerFields []huh.Field

		typeField, err := CreateHuhField("Type", &dataConfig.Type.Prompt, &cfg.Type, language)
		if err != nil {
			return fmt.Errorf("failed to create field for Container Type: %w", err)
		}
		containerFields = append(containerFields, typeField)

		nameField, err := CreateHuhField("Name", &dataConfig.Name.Prompt, &cfg.Name, language)
		if err != nil {
			return fmt.Errorf("failed to create field for Container Name: %w", err)
		}
		containerFields = append(containerFields, nameField)

		tagField, err := CreateHuhField("Tag", &dataConfig.Tag.Prompt, &cfg.Tag, language)
		if err != nil {
			return fmt.Errorf("failed to create field for Container Tag: %w", err)
		}
		containerFields = append(containerFields, tagField)

		if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.CORE_SPACE {

			coreSpacePortField, err := CreateHuhField("CoreSpacePort", &dataConfig.CoreSpace.Port.Prompt, &cfg.CoreSpace.Port, language)
			if err != nil {
				return fmt.Errorf("failed to create field for Container CoreSpacePort: %w", err)
			}
			containerFields = append(containerFields, coreSpacePortField)
		}

		if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.E_SPACE {
			eSpacePortField, err := CreateHuhField("ESpacePort", &dataConfig.ESpace.Port.Prompt, &cfg.ESpace.Port, language)
			if err != nil {
				return fmt.Errorf("failed to create field for Container ESpacePort: %w", err)
			}
			containerFields = append(containerFields, eSpacePortField)
		}

		if err := huh.NewForm(
			huh.NewGroup(
				containerFields...,
			),
		).Run(); err != nil {
			return err
		}
	}
	return nil
}
