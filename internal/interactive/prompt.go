package interactive

import (
	"Conflux-Chain/sirius-auto-release/internal/config"

	"github.com/charmbracelet/huh"
	"github.com/pelletier/go-toml/v2"
)

func CollectConfig(cfg *config.Config) (*config.Config, error) {

	language := "en"

	if err := huh.NewSelect[string]().Title("Select Language(选择语言)").Options(huh.NewOption("English", "en"), huh.NewOption("中文", "zh")).Value(&language).Run(); err != nil {
		return nil, err
	}

	var dataConfig config.DataConfig
	if err := toml.Unmarshal([]byte(config.ConfigDataEmbed), &dataConfig); err != nil {
		return nil, err
	}
	if err := runGlobalForm(&cfg.Global, dataConfig.Global, language); err != nil {
		return nil, err
	}

	if err := runFrontendForm(&cfg.Frontend, &cfg.Global, dataConfig.Frontend, language); err != nil {
		return nil, err
	}

	if err := runContainerForm(&cfg.Container, &cfg.Global, dataConfig.Container, language); err != nil {
		return nil, err
	}

	// only run proxy form if container is enabled

	if cfg.Container.Enabled {
		if err := runProxyForm(&cfg.Proxy, &cfg.Global, dataConfig.Proxy, language); err != nil {
			return nil, err
		}

	}

	return cfg, nil
}
