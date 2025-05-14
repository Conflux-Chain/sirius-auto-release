package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"log/slog"
)

func RunDockerScript() error {
	// TODO
	return nil
}

func RunScript(config *config.Config) error {

	slog.Debug("Starting script execution")
	// frontend
	if err := RunFrontendScript(&config.Frontend, &config.Global); err != nil {
		return err
	}

	if config.Proxy.Enabled {
		if err := RunProxyScript(&config.Proxy, &config.Global); err != nil {
			return err
		}
	}

	if config.Docker.Enabled {
		if err := RunDockerScript(); err != nil {
			return err
		}
	}

	slog.Debug("Script execution completed successfully")
	return nil
}
