package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"bytes"
	"fmt"
	"net"
	"net/url"
	"path/filepath"
	"text/template"
)

type NginxServerTemplate struct {
	APIUrl     string
	APIHost    string
	ListenPort int
	AssetDir   string
}

func (t *NginxServerTemplate) generateNginxConfig() ([]byte, error) {

	parseURL, err := url.Parse(t.APIUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	hostname := parseURL.Hostname()

	if net.ParseIP(hostname) == nil {
		t.APIHost = hostname
	} else {
		t.APIHost = ""
	}

	return generateFromTemplate(config.NginxServerTemplate, t)
}

type NginxBaseTemplateData struct {
	Servers string
}

func (d *NginxBaseTemplateData) generateNginxConfig() ([]byte, error) {
	return generateFromTemplate(config.NginxBaseTemplate, d)
}

func generateFromTemplate(templateText string, data interface{}) ([]byte, error) {
	tmpl, err := template.New("nginx").Parse(templateText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nginx template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute nginx template: %w", err)
	}

	return buf.Bytes(), nil
}

func writeToFile(content []byte, outPath string) error {
	if err := utils.WriteToFile(content, outPath); err != nil {
		return fmt.Errorf("failed to write nginx config: %w", err)
	}
	fmt.Println("Nginx configuration file generated at:", outPath)
	return nil
}

func getServerConfigs(proxyConfig *config.Proxy, globalConfig *config.Global) ([]NginxServerTemplate, error) {
	var configs []NginxServerTemplate

	switch globalConfig.Space {
	case "all":
		configs = append(configs,
			NginxServerTemplate{
				APIUrl:     proxyConfig.CoreSpace.API_URL,
				AssetDir:   "scan/build",
				ListenPort: proxyConfig.CoreSpace.Port,
			},
			NginxServerTemplate{
				APIUrl:     proxyConfig.ESpace.API_URL,
				AssetDir:   "scan-eth/build",
				ListenPort: proxyConfig.ESpace.Port,
			},
		)
	case "core":
		configs = append(configs, NginxServerTemplate{
			APIUrl:     proxyConfig.CoreSpace.API_URL,
			AssetDir:   "scan/build",
			ListenPort: proxyConfig.CoreSpace.Port,
		})
	case "eSpace":
		configs = append(configs, NginxServerTemplate{
			APIUrl:     proxyConfig.ESpace.API_URL,
			AssetDir:   "scan-eth/build",
			ListenPort: proxyConfig.ESpace.Port,
		})
	default:
		return nil, fmt.Errorf("no matching assets found for space: %s", globalConfig.Space)
	}

	return configs, nil
}

func RunProxyScript(proxyConfig *config.Proxy, globalConfig *config.Global) error {

	utils.ShowTitle("Proxy")

	configs, err := getServerConfigs(proxyConfig, globalConfig)

	utils.ShowStep(fmt.Sprintf("Generating Nginx configuration for %s space...\n", globalConfig.Space))

	if err != nil {
		return err
	}

	// nginx server config
	var content bytes.Buffer
	for _, config := range configs {
		buf, err := config.generateNginxConfig()
		if err != nil {
			return fmt.Errorf("failed to generate server config: %w", err)
		}
		content.Write(buf)
		content.WriteString("\n")
	}

	// nginx base config
	data := NginxBaseTemplateData{
		Servers: string(content.Bytes()),
	}

	buf, err := data.generateNginxConfig()
	if err != nil {
		return fmt.Errorf("failed to generate base config: %w", err)
	}

	output := filepath.Join(globalConfig.Workdir, "nginx.conf")
	utils.ShowStep(fmt.Sprintf("Writing Nginx configuration to %s...\n", output))
	return writeToFile(buf, output)
}
