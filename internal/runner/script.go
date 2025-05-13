package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
)

func RunFrontendScript(frontendConfig *config.Frontend, globalConfig *config.Global) error {
	switch frontendConfig.Type {
	case "prebuilt":
		{
			githubRelease := GitHub{
				RepoURL: frontendConfig.PrebuiltRepo,
			}

			// Get the latest release
			release, err := githubRelease.GetLatestRelease()
			if err != nil {
				return fmt.Errorf("failed to get latest release: %v", err)
			}
			fmt.Printf("Latest prebuilt release: %s\n", release.TagName)

			tempDir, err := os.MkdirTemp("", "sirius")
			if err != nil {
				return fmt.Errorf("failed to create temp dir: %v", err)
			}

			for _, asset := range release.Assets {
				if strings.Contains(asset.Name, "scan") {
					fmt.Printf("Downloading release %s... (this may take a while)\n", asset.Name)

					outPath := path.Join(tempDir, asset.Name)

					if err := utils.DownloadFile(asset.BrowserDownloadURL, outPath); err != nil {
						return fmt.Errorf("failed to download file: %v", err)
					}
				}
			}

			defer func() {
				slog.Debug("Removing temp dir", "path", tempDir)
				if err := os.RemoveAll(tempDir); err != nil {
					slog.Error("Failed to remove temp dir", "error", err)
				}
			}()

			// TODO unzip the downloaded file to the workdir
			fmt.Printf("Preparing frontend in %s\n", globalConfig.Workdir)
		}
	default:
		{
			return fmt.Errorf("unknown frontend type: %s", frontendConfig.Type)
		}
	}

	return nil
}

func RunDockerScript() error {
	// TODO
	return nil
}

func RunProxyScript() error {
	// TODO
	return nil
}

func RunScript(config *config.Config) error {

	slog.Debug("Starting script execution")
	// frontend
	if err := RunFrontendScript(&config.Frontend, &config.Global); err != nil {
		return err
	}

	if config.Docker.Enabled {
		if err := RunDockerScript(); err != nil {
			return err
		}
	}

	if config.Proxy.Enabled {
		if err := RunProxyScript(); err != nil {
			return err
		}
	}
	slog.Debug("Script execution completed successfully")
	return nil
}
