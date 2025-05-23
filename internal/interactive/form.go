package interactive

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/i18n"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

func runGlobalForm(cfg *config.Global, language string) error {

	cfg.Workdir = "./frontend"
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(i18n.T(language, "config.global.space")).Options(
				huh.NewOption(config.CORE_SPACE, config.CORE_SPACE),
				huh.NewOption(config.E_SPACE, config.E_SPACE),
				huh.NewOption(config.ALL_SPACE, config.ALL_SPACE),
			).Value(&cfg.Space),

			huh.NewInput().Title(i18n.T(language, "config.global.workdir")).Value(&cfg.Workdir).Placeholder("./frontend"),
		),
	)

	return form.Run()

}

func runFrontendForm(cfg *config.Frontend, globalConfig *config.Global, language string) error {

	cfg.PrebuiltRepo = "https://github.com/Conflux-Chain/sirius-auto-release.git"
	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(i18n.T(language, "Config.Frontend.Type")).Options(
				huh.NewOption("prebuilt", "prebuilt"),
			).Value(&cfg.Type),
			huh.NewInput().Title(i18n.T(language, "Config.Frontend.PrebuiltRepo")).Value(&cfg.PrebuiltRepo).Placeholder("https://github.com/Conflux-Chain/sirius-auto-release.git"),
		),
	).Run(); err != nil {
		return err
	}

	if err := huh.NewConfirm().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.Enabled")).Value(&cfg.CoreSpaceSettings.Enabled).Run(); err != nil {
		return err
	}

	if (globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.CORE_SPACE) && cfg.CoreSpaceSettings.Enabled {

		networkId := ""
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvOpenApiHost")).Value(&cfg.CoreSpaceSettings.EnvOpenApiHost),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvRpcServer")).Value(&cfg.CoreSpaceSettings.EnvRpcServer),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvNetworkID")).Value(&networkId).Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("networkId cannot be empty")
					}
					if _, err := strconv.Atoi(s); err != nil {
						return fmt.Errorf("networkId must be a number")
					}

					return nil
				}),

				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvNetworkType")).Value(&cfg.CoreSpaceSettings.EnvNetworkType),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvChainType")).Value(&cfg.CoreSpaceSettings.EnvChainType),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvFcAddress")).Value(&cfg.CoreSpaceSettings.EnvFcAddress),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvFcExchangeAddress")).Value(&cfg.CoreSpaceSettings.EnvFcExchangeAddress),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvFcExchangeInterestAddress")).Value(&cfg.CoreSpaceSettings.EnvFcExchangeInterestAddress),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvEnsRegistryAddress")).Value(&cfg.CoreSpaceSettings.EnvEnsRegistryAddress),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvEnsPublicResolverAddress")).Value(&cfg.CoreSpaceSettings.EnvEnsPublicResolverAddress),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.CoreSpaceSettings.EnvLogo")).Value(&cfg.CoreSpaceSettings.EnvLogo),
			),
		).Run(); err != nil {
			return err
		}

		// network id is already validated
		netId, _ := strconv.Atoi(networkId)
		cfg.CoreSpaceSettings.EnvNetworkID = netId

	}

	if err := huh.NewConfirm().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.Enabled")).Value(&cfg.ESpaceSettings.Enabled).Run(); err != nil {
		return err
	}

	if (globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.E_SPACE) && cfg.ESpaceSettings.Enabled {

		networkId := ""
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvApiHost")).Value(&cfg.ESpaceSettings.EnvApiHost),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvCoreApiHost")).Value(&cfg.ESpaceSettings.EnvCoreApiHost),

				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvCoreScanHost")).Value(&cfg.ESpaceSettings.EnvCoreScanHost),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvRpcServer")).Value(&cfg.ESpaceSettings.EnvRpcServer),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvNetworkID")).Value(&networkId).Validate(utils.ValidateNumber("networkId")),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvChainType")).Value(&cfg.ESpaceSettings.EnvChainType),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvLogo")).Value(&cfg.ESpaceSettings.EnvLogo),
			),
		).Run(); err != nil {
			return err
		}

		netId, _ := strconv.Atoi(networkId)
		cfg.ESpaceSettings.EnvNetworkID = netId

		chainId := ""

		rpcUrl := ""
		blockExplorerUrl := ""

		decimals := "18"
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvWalletConfig.ChainID")).Value(&chainId).Validate(utils.ValidateNumber("chainId")),

				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvWalletConfig.ChainName")).Value(&cfg.ESpaceSettings.EnvWalletConfig.ChainName),

				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvWalletConfig.RpcUrls")).Value(&rpcUrl),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvWalletConfig.BlockExplorerUrls")).Value(&blockExplorerUrl),

				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvWalletConfig.NativeCurrency.Name")).Value(&cfg.ESpaceSettings.EnvWalletConfig.NativeCurrency.Name),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvWalletConfig.NativeCurrency.Symbol")).Value(&cfg.ESpaceSettings.EnvWalletConfig.NativeCurrency.Symbol),
				huh.NewInput().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.EnvWalletConfig.NativeCurrency.Decimals")).Value(&decimals).Validate(utils.ValidateNumber("decimals")).Placeholder("18"),
			),
		).Run(); err != nil {
			return err
		}

		chainIdInt, _ := strconv.Atoi(chainId)
		cfg.ESpaceSettings.EnvWalletConfig.ChainID = chainIdInt
		cfg.ESpaceSettings.EnvWalletConfig.RpcUrls = []string{rpcUrl}
		cfg.ESpaceSettings.EnvWalletConfig.BlockExplorerUrls = []string{blockExplorerUrl}
		decimalsInt, _ := strconv.Atoi(decimals)
		cfg.ESpaceSettings.EnvWalletConfig.NativeCurrency.Decimals = decimalsInt

	}

	return nil
}

func runProxyForm(cfg *config.Proxy, globalConfig *config.Global, language string) error {

	if err := huh.NewConfirm().Title(i18n.T(language, "Config.Proxy.Enabled")).Value(&cfg.Enabled).Run(); err != nil {
		return err
	}
	if cfg.Enabled {
		cfg.Type = "nginx"

		if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.CORE_SPACE {

			port := ""
			if err := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title(i18n.T(language, "Config.Proxy.CoreSpace.Port")).Value(&port).Validate(utils.ValidateNumber("port")),
					huh.NewInput().Title(i18n.T(language, "Config.Proxy.CoreSpace.API_URL")).Value(&cfg.CoreSpace.API_URL),
				),
			).Run(); err != nil {
				return err
			}
			portInt, _ := strconv.Atoi(port)
			cfg.CoreSpace.Port = portInt
		}

		if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.E_SPACE {
			port := ""
			if err := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title(i18n.T(language, "Config.Proxy.ESpace.Port")).Value(&port).Validate(utils.ValidateNumber("port")),
					huh.NewInput().Title(i18n.T(language, "Config.Proxy.ESpace.API_URL")).Value(&cfg.ESpace.API_URL),
				),
			).Run(); err != nil {
				return err
			}

			portInt, _ := strconv.Atoi(port)
			cfg.ESpace.Port = portInt
		}
	}
	return nil
}

func runContainerForm(cfg *config.Container, globalConfig *config.Global, language string) error {

	if err := huh.NewConfirm().Title(i18n.T(language, "Config.Container.Enabled")).Value(&cfg.Enabled).Run(); err != nil {
		return err
	}

	if cfg.Enabled {
		cfg.Type = "docker"

		if err := huh.NewInput().Title(i18n.T(language, "Config.Container.Name")).Value(&cfg.Name).Run(); err != nil {
			return err
		}

		if err := huh.NewInput().Title(i18n.T(language, "Config.Container.Tag")).Value(&cfg.Tag).Run(); err != nil {
			return err
		}

		if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.CORE_SPACE {

			port := ""
			if err := huh.NewInput().Title(fmt.Sprintf("%s %s", i18n.T(language, "Config.Container.CoreSpace"), i18n.T(language, "Config.Container.CoreSpace.Port"))).Value(&port).Validate(utils.ValidateNumber("port")).Run(); err != nil {
				return err
			}

			portInt, _ := strconv.Atoi(port)
			cfg.CoreSpace.Port = portInt
		}

		if globalConfig.Space == config.ALL_SPACE || globalConfig.Space == config.E_SPACE {

			port := ""
			if err := huh.NewInput().Title(fmt.Sprintf("%s %s", i18n.T(language, "Config.Container.ESpace"), i18n.T(language, "Config.Container.ESpace.Port"))).Value(&port).Validate(utils.ValidateNumber("port")).Run(); err != nil {
				return err
			}

			portInt, _ := strconv.Atoi(port)
			cfg.ESpace.Port = portInt

		}
	}

	return nil
}
