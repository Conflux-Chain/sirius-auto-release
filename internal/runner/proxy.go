package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"bytes"
	"fmt"
	"html/template"
	"path"
)

type NginxTemplateData struct {
	APIURL string
	Port   int
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
	data := NginxTemplateData{
		APIURL: proxyConfig.API_URL,
		Port:   proxyConfig.Port,
	}
	output := path.Join(globalConfig.Workdir, "nginx.conf")
	err := data.writeToFile(output)
	return err
}
