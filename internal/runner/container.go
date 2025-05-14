package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"
	"path/filepath"
)

func RunContainerScript(cfg *config.Container, globalConfig *config.Global) error {
	if cfg.Type == "docker" {
		dockerfile := config.DockerTemplate
		outputPath := filepath.Join(globalConfig.Workdir, "Dockerfile")
		if err := utils.WriteToFile([]byte(dockerfile), outputPath); err != nil {
			return fmt.Errorf("failed to write Dockerfile: %v", err)
		}
	}
	return nil
}
