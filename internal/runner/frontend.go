package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

type GitHubReleaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type GitHubRelease struct {
	TagName string               `json:"tag_name"`
	Assets  []GitHubReleaseAsset `json:"assets"`
}

type GitHub struct {
	RepoURL string
}

func (g *GitHub) GetLatestRelease() (GitHubRelease, error) {

	parseURL, err := url.Parse(g.RepoURL)

	slog.Debug("Parsed URL", "url", parseURL)

	if err != nil {
		return GitHubRelease{}, fmt.Errorf("failed to parse URL: %v", err)
	}

	if parseURL.Host != "github.com" {
		return GitHubRelease{}, fmt.Errorf("invalid GitHub URL: %s", g.RepoURL)
	}

	pathParts := strings.Split(strings.Trim(parseURL.Path, "/"), "/")

	if len(pathParts) != 2 {
		return GitHubRelease{}, fmt.Errorf("invalid GitHub URL: %s", g.RepoURL)
	}

	repoOwner := pathParts[0]
	repoName := strings.TrimSuffix(pathParts[1], ".git")

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)

	slog.Debug("Fetching latest release from", "url", apiURL)

	client := http.Client{
		Timeout: 10 * time.Minute,
	}

	resp, err := client.Get(apiURL)

	if err != nil {
		return GitHubRelease{}, fmt.Errorf("failed to get latest release: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GitHubRelease{}, fmt.Errorf("failed to get latest release: %s", resp.Status)
	}

	var releaseInfo GitHubRelease

	if err := json.NewDecoder(resp.Body).Decode(&releaseInfo); err != nil {
		return GitHubRelease{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return releaseInfo, nil
}

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

			var asset GitHubReleaseAsset

			var assetName string
			if frontendConfig.Space == "eSpace" {
				assetName = "scan-eth.zip"
			} else if frontendConfig.Space == "coreSpace" {
				assetName = "scan.zip"
			} else {
				return fmt.Errorf("unknown space: %s", frontendConfig.Space)
			}

			idx := slices.IndexFunc(release.Assets, func(asset GitHubReleaseAsset) bool {
				return asset.Name == assetName
			})
			if idx == -1 {
				return fmt.Errorf("failed to find asset: %s", assetName)
			}
			asset = release.Assets[idx]

			fmt.Printf("Downloading release %s... (this may take a while)\n", asset.Name)

			outPath := filepath.Join(tempDir, asset.Name)

			if err := utils.DownloadFile(asset.BrowserDownloadURL, outPath); err != nil {
				return fmt.Errorf("failed to download file: %v", err)
			}
			name := strings.TrimSuffix(asset.Name, ".zip")
			utils.Unzip(outPath, filepath.Join(globalConfig.Workdir, name))
			fmt.Printf("Preparing frontend in %s\n", filepath.Join(globalConfig.Workdir, name))

			// move frontend to workdir
			if _, err := os.Stat(filepath.Join(globalConfig.Workdir, "build")); err == nil {
				os.RemoveAll(filepath.Join(globalConfig.Workdir, "build"))

			}

			// move build to workdir
			if err := os.Rename(filepath.Join(globalConfig.Workdir, name, "build"), filepath.Join(globalConfig.Workdir, "build")); err != nil {
				return fmt.Errorf("failed to move build: %v", err)
			}

			// remove empty dir
			if err := os.RemoveAll(filepath.Join(globalConfig.Workdir, name)); err != nil {
				return fmt.Errorf("failed to remove dir: %v", err)
			}

			// cleanup
			defer func() {
				slog.Debug("Removing temp dir", "path", tempDir)
				if err := os.RemoveAll(tempDir); err != nil {
					slog.Error("Failed to remove temp dir", "error", err)
				}
			}()

		}
	default:
		{
			return fmt.Errorf("unknown frontend type: %s", frontendConfig.Type)
		}
	}

	return nil
}
