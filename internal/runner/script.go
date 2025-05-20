package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"
	"log/slog"
)

func RunScript(cfg *config.Config) error {

	slog.Debug("Starting script execution")
	if err := RunFrontendScript(&cfg.Frontend, &cfg.Global); err != nil {
		return err
	}
	utils.ShowSeparator()

	if cfg.Proxy.Enabled {
		if err := RunProxyScript(&cfg.Proxy, &cfg.Global); err != nil {
			return err
		}
		utils.ShowSeparator()
	}

	if cfg.Container.Enabled {
		if err := RunContainerScript(cfg); err != nil {
			return err
		}
		utils.ShowSeparator()
	}

	// Print help message
	if cfg.Container.Enabled {

		if cfg.Container.Type == config.CONTAINER_TYPE_DOCKER {

			utils.ShowCommandSuggestion("To build and run the Docker container, use the following commands:")

			message := fmt.Sprint("To build the Docker image, run:")

			message += "\n\n"
			message += fmt.Sprintf("cd %s &&", cfg.Global.Workdir)

			message += fmt.Sprintf(" docker build -t %s:%s . &&", cfg.Container.Name, cfg.Container.Tag)

			message += fmt.Sprintf(" docker run ")
			if ports := utils.GetPortBindingForConfig(&cfg.Global, &cfg.Proxy, &cfg.Container); len(ports) > 0 {
				for _, port := range ports {
					message += fmt.Sprintf(" -p %s", port)
				}
			}
			message += fmt.Sprintf(" %s:%s \n", cfg.Container.Name, cfg.Container.Tag)
			utils.ShowCommand(message)

		} else if cfg.Container.Type == config.CONTAINER_TYPE_DOCKER_COMPOSE {

			utils.ShowCommandSuggestion("To start the Docker container, run:")
			utils.ShowCommand(fmt.Sprintf("cd %s &&  docker compose up -d --build ", cfg.Global.Workdir))
		}

	}

	return nil
}
