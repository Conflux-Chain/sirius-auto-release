package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
)

type DockerfileTemplateData struct {
	CopyPaths string
}

func getCopyPathsForSpace(space string) (string, error) {
	switch space {
	case "all":
		return "COPY ./scan /app/frontend/scan\nCOPY ./scan-eth /app/frontend/scan-eth\n", nil
	case "core":
		return "COPY ./scan /app/frontend/scan\n", nil
	case "eSpace":
		return "COPY ./scan-eth /app/frontend/scan-eth\n", nil
	default:
		return "", fmt.Errorf("unknown space: %s", space)
	}
}

func docker(cfg *config.Container, globalConfig *config.Global) error {

	dockerfile := config.DockerTemplate

	tmpl, err := template.New("dockerfile").Parse(dockerfile)
	if err != nil {
		return fmt.Errorf("failed to parse Dockerfile template: %v", err)
	}

	var buf bytes.Buffer

	copyPaths, err := getCopyPathsForSpace(globalConfig.Space)

	if err != nil {
		return err
	}
	data := DockerfileTemplateData{
		CopyPaths: copyPaths,
	}

	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute Dockerfile template: %v", err)
	}

	outputPath := filepath.Join(globalConfig.Workdir, "Dockerfile")
	if err := utils.WriteToFile(buf.Bytes(), outputPath); err != nil {
		return fmt.Errorf("failed to write Dockerfile: %v", err)
	}

	return nil
}

type DockerComposeTemplateData struct {
	ServiceName string
	Ports       []string
}

func dockerCompose(cfg *config.Config) error {

	dockerCompose := config.DockerComposeTemplate

	tmpl, err := template.New("docker-compose").Parse(dockerCompose)
	if err != nil {
		return fmt.Errorf("failed to parse docker-compose template: %v", err)
	}

	var buf bytes.Buffer
	ports := utils.GetPortBindingForConfig(&cfg.Global, &cfg.Proxy, &cfg.Container)

	data := DockerComposeTemplateData{
		ServiceName: cfg.Container.Name,
		Ports:       ports,
	}

	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute docker-compose template: %v", err)
	}
	outputPath := filepath.Join(cfg.Global.Workdir, "docker-compose.yml")
	if err := utils.WriteToFile(buf.Bytes(), outputPath); err != nil {
		return fmt.Errorf("failed to write docker-compose.yml: %v", err)
	}
	return nil

}

func RunContainerScript(cfg *config.Config) error {

	switch cfg.Container.Type {
	case config.CONTAINER_TYPE_DOCKER:
		{

			return docker(&cfg.Container, &cfg.Global)
		}
	case config.CONTAINER_TYPE_DOCKER_COMPOSE:
		{
			if err := docker(&cfg.Container, &cfg.Global); err != nil {
				return err
			}
			return dockerCompose(cfg)
		}
	default:
		return fmt.Errorf("unknown container type: %s", cfg.Container.Type)

	}

}
