package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"log/slog"
)

func RunScript(config *config.Config) error {

	slog.Debug("Starting script execution")
	if err := RunFrontendScript(&config.Frontend, &config.Global); err != nil {
		return err
	}

	if config.Proxy.Enabled {
		if err := RunProxyScript(&config.Proxy, &config.Global); err != nil {
			return err
		}
	}

	if config.Container.Enabled {
		if err := RunContainerScript(&config.Container, &config.Global); err != nil {
			return err
		}
	}

	slog.Debug("Script execution completed successfully")
	return nil
}
