package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
)

type NginxTemplateData struct {
	APIURL string
	Port   int
}

func (d *NginxTemplateData) writeToFile(outPath string) error {

	fmt.Println("Generating Nginx configuration file...")

	temp, err := template.New("nginx").Parse(config.NginxTemplate)

	if err != nil {
		return fmt.Errorf("failed to parse nginx template: %v", err)
	}

	var buf bytes.Buffer

	if err := temp.Execute(&buf, d); err != nil {
		return fmt.Errorf("failed to execute nginx template: %v", err)
	}

	dir := filepath.Dir(outPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
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
