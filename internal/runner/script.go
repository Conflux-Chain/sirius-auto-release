package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"fmt"
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

	// Print help message
	if config.Container.Enabled {
		if config.Container.Type == "docker" {
			fmt.Println("Dockerfile generated successfully.")
			fmt.Println("To build the Docker image, run:")
			fmt.Print("\n\n")
			fmt.Printf("cd %s && docker build -t %s:%s . && docker run -p %d:3000  %s:%s \n", config.Global.Workdir, config.Container.Name, config.Container.Tag, config.Container.Port, config.Container.Name, config.Container.Tag)
			fmt.Print("\n\n")
		}

	}

	return nil
}
