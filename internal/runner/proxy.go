package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"path/filepath"
)

type NginxTemplateData struct {
	APIURL   string
	API_HOST string
}

func (d *NginxTemplateData) generateNginxConfig() ([]byte, error) {
	fmt.Println("Generating Nginx configuration file...")

	temp, err := template.New("nginx").Parse(config.NginxTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nginx template: %v", err)
	}

	var buf bytes.Buffer
	if err := temp.Execute(&buf, d); err != nil {
		return nil, fmt.Errorf("failed to execute nginx template: %v", err)
	}

	return buf.Bytes(), nil
}

func (d *NginxTemplateData) writeToFile(outPath string) error {
	content, err := d.generateNginxConfig()
	if err != nil {
		return err
	}

	if err := utils.WriteToFile(content, outPath); err != nil {
		return err
	}
	fmt.Println("Nginx configuration file generated at:", outPath)
	return nil
}
func RunProxyScript(proxyConfig *config.Proxy, globalConfig *config.Global) error {
	parseURL, err := url.Parse(proxyConfig.API_URL)

	if err != nil {
		return fmt.Errorf("failed to parse URL: %v", err)
	}

	data := NginxTemplateData{
		APIURL:   proxyConfig.API_URL,
		API_HOST: parseURL.Hostname(),
	}
	output := filepath.Join(globalConfig.Workdir, "nginx.conf")

	return data.writeToFile(output)
}
