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

	if cfg.Proxy.Enabled {
		if err := RunProxyScript(&cfg.Proxy, &cfg.Global); err != nil {
			return err
		}
	}

	if cfg.Container.Enabled {
		if err := RunContainerScript(cfg); err != nil {
			return err
		}
	}

	// Print help message
	if cfg.Container.Enabled {
		fmt.Println("Dockerfile generated successfully.")
		if cfg.Container.Type == config.CONTAINER_TYPE_DOCKER {
			fmt.Println("To build the Docker image, run:")
			fmt.Print("\n\n")
			fmt.Printf("cd %s &&", cfg.Global.Workdir)
			fmt.Printf(" docker build -t %s:%s . &&", cfg.Container.Name, cfg.Container.Tag)
			fmt.Printf(" docker run ")

			if ports := utils.GetPortBindingForConfig(&cfg.Global, &cfg.Proxy, &cfg.Container); len(ports) > 0 {
				for _, port := range ports {
					fmt.Printf(" -p %s", port)
				}
			}

			fmt.Printf(" %s:%s \n", cfg.Container.Name, cfg.Container.Tag)

			fmt.Print("\n\n")
		} else if cfg.Container.Type == config.CONTAINER_TYPE_DOCKER_COMPOSE {
			fmt.Println("To start the Docker container, run:")
			fmt.Print("\n\n")
			fmt.Printf("cd %s &&", cfg.Global.Workdir)
			fmt.Printf(" docker compose up -d --build  \n")
			fmt.Print("\n\n")
		}

	}

	return nil
}
