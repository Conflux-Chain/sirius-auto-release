package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
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
		if err := RunContainerScript(&cfg.Container, &cfg.Global); err != nil {
			return err
		}
	}

	// Print help message
	if cfg.Container.Enabled {
		if cfg.Container.Type == config.CONTAINER_TYPE_DOCKER {
			fmt.Println("Dockerfile generated successfully.")
			fmt.Println("To build the Docker image, run:")
			fmt.Print("\n\n")
			fmt.Printf("cd %s &&", cfg.Global.Workdir)
			fmt.Printf(" docker build -t %s:%s . &&", cfg.Container.Name, cfg.Container.Tag)
			fmt.Printf(" docker run ")
			if cfg.Global.Space == config.ALL_SPACE {
				fmt.Printf(" -p %d:%d -p %d:%d", cfg.Container.CoreSpace.Port, cfg.Proxy.CoreSpace.Port, cfg.Container.ESpace.Port, cfg.Proxy.ESpace.Port)
			} else if cfg.Global.Space == config.CORE_SPACE {
				fmt.Printf(" -p %d:%d", cfg.Container.CoreSpace.Port, cfg.Proxy.CoreSpace.Port)
			} else if cfg.Global.Space == config.E_SPACE {
				fmt.Printf(" -p %d:%d", cfg.Container.ESpace.Port, cfg.Proxy.ESpace.Port)
			}
			fmt.Printf(" %s:%s \n", cfg.Container.Name, cfg.Container.Tag)

			fmt.Print("\n\n")
		}

	}

	return nil
}
