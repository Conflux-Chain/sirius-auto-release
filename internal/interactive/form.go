package interactive

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/i18n"

	"github.com/charmbracelet/huh"
)

func runGlobalForm(cfg *config.Global, language string) error {

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

func runFrontendForm(cfg *config.Frontend, language string) error {

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

	if cfg.CoreSpaceSettings.Enabled {
		// TODO: add core space settings
	}

	if err := huh.NewConfirm().Title(i18n.T(language, "Config.Frontend.ESpaceSettings.Enabled")).Value(&cfg.ESpaceSettings.Enabled).Run(); err != nil {
		return err
	}

	if cfg.ESpaceSettings.Enabled {
		// TODO: add ESpace settings
	}

	return nil
}
