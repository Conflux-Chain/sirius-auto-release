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

func prebuilt(frontendConfig *config.Frontend, globalConfig *config.Global) error {

	github := GitHub{
		RepoURL: frontendConfig.PrebuiltRepo,
	}

	release, err := github.GetLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to get latest release: %v", err)
	}
	fmt.Printf("Latest prebuilt release: %s\n", release.TagName)

	tempDir, err := os.MkdirTemp("", "sirius")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}

	defer func() {
		slog.Debug("Removing temporary directory", "path", tempDir)
		if err := os.RemoveAll(tempDir); err != nil {
			slog.Error("Failed to remove temporary directory", "error", err)
		}
	}()

	assetsToDownload := selectAssetsToDownload(globalConfig.Space, release.Assets)
	if len(assetsToDownload) == 0 {
		return fmt.Errorf("no matching assets found for space: %s", globalConfig.Space)
	}

	for _, asset := range assetsToDownload {
		fmt.Printf("Downloading release %s... (this may take a while)\n", asset.Name)

		downloadPath := filepath.Join(tempDir, asset.Name)
		if err := utils.DownloadFile(asset.BrowserDownloadURL, downloadPath); err != nil {
			return fmt.Errorf("failed to download file %s: %v", asset.Name, err)
		}

		extractName := strings.TrimSuffix(asset.Name, ".zip")
		extractPath := filepath.Join(globalConfig.Workdir, extractName)
		if err := utils.Unzip(downloadPath, extractPath); err != nil {
			return fmt.Errorf("failed to extract %s: %v", asset.Name, err)
		}

		fmt.Printf("Frontend prepared in %s\n", extractPath)
	}

	return nil
}

// Helper function to select which assets to download based on space configuration
func selectAssetsToDownload(space string, allAssets []GitHubReleaseAsset) []GitHubReleaseAsset {
	if space == config.ALL_SPACE {
		return allAssets
	}

	var assetName string
	switch space {
	case config.E_SPACE:
		assetName = "scan-eth.zip"
	case config.CORE_SPACE:
		assetName = "scan.zip"
	default:
		return nil
	}

	for _, asset := range allAssets {
		if asset.Name == assetName {
			return []GitHubReleaseAsset{asset}
		}
	}

	return nil
}

func RunFrontendScript(frontendConfig *config.Frontend, globalConfig *config.Global) error {
	switch frontendConfig.Type {
	case "prebuilt":
		{
			return prebuilt(frontendConfig, globalConfig)
		}
	default:
		{
			return fmt.Errorf("unknown frontend type: %s", frontendConfig.Type)
		}
	}
}
