package interactive

import (
	"Conflux-Chain/sirius-auto-release/internal/config"

	"github.com/charmbracelet/huh"
)

func CollectConfig(cfg *config.Config) (*config.Config, error) {

	language := "en"

	if err := huh.NewSelect[string]().Title("Select Language(选择语言)").Options(huh.NewOption("English", "en"), huh.NewOption("中文", "zh")).Value(&language).Run(); err != nil {
		return nil, err
	}

	if err := runGlobalForm(&cfg.Global, language); err != nil {
		return nil, err
	}

	if err := runFrontendForm(&cfg.Frontend, &cfg.Global, language); err != nil {
		return nil, err
	}

	if err := runProxyForm(&cfg.Proxy, &cfg.Global, language); err != nil {
		return nil, err
	}

	if err := runContainerForm(&cfg.Container, &cfg.Global, language); err != nil {
		return nil, err
	}

	return cfg, nil
}
