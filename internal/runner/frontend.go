package runner

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type GitHubReleaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func (g *GitHubReleaseAsset) WriteSettingsToFile(cfg *config.Frontend, globalConfig *config.Global) error {
	dirName := strings.TrimSuffix(g.Name, ".zip")

	// Determine if we need to inject settings and what data to inject
	var injectSettings bool
	var injectData any

	switch {
	case globalConfig.Space == config.ALL_SPACE:
		if dirName == config.CORE_SPACE_DIR {
			injectSettings = true
			injectData = cfg.CoreSpaceSettings
		} else if dirName == config.E_SPACE_DIR {
			injectSettings = true
			injectData = cfg.ESpaceSettings
		}
	case globalConfig.Space == config.CORE_SPACE && dirName == config.CORE_SPACE_DIR:
		injectSettings = true
		injectData = cfg.CoreSpaceSettings
	case globalConfig.Space == config.E_SPACE && dirName == config.E_SPACE_DIR:
		injectSettings = true
		injectData = cfg.ESpaceSettings
	}

	if !injectSettings {
		return nil
	}

	// Marshal settings to JSON
	jsonData, err := json.Marshal(injectData)
	if err != nil {
		return fmt.Errorf("failed to marshal inject data: %v", err)
	}

	// Create script tag with settings
	scriptTag := fmt.Sprintf("<script type=\"text/javascript\">\n  window.customConfig = %s;\n</script>", string(jsonData))
	log.Debug("Injecting settings into HTML", "scriptTag", scriptTag)

	// Read HTML file
	htmlFilePath := filepath.Join(globalConfig.Workdir, dirName, "build", "index.html")
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %v", err)
	}

	// Inject script tag before closing body tag
	newHTMLContent := strings.Replace(string(htmlContent), "<body>", scriptTag+"<body>", 1)

	// Write modified HTML back to file
	if err := os.WriteFile(htmlFilePath, []byte(newHTMLContent), 0644); err != nil {
		return fmt.Errorf("failed to write HTML file: %v", err)
	}

	log.Debug("Injected settings into HTML", "filePath", htmlFilePath)
	fmt.Println("Injected settings into HTML file:", htmlFilePath)

	return nil
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

	log.Debug("Parsed URL", "url", parseURL)

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

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", repoOwner, repoName)

	log.Debug("Fetching release from", "url", apiURL)

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

	var releaseInfo []GitHubRelease

	if err := json.NewDecoder(resp.Body).Decode(&releaseInfo); err != nil {
		return GitHubRelease{}, fmt.Errorf("failed to decode response: %v", err)
	}

	index := slices.IndexFunc(releaseInfo, func(release GitHubRelease) bool {

		for _, asset := range release.Assets {
			if strings.HasPrefix(asset.Name, "scan") {
				return true
			}
		}
		return false
	})

	if index == -1 {
		return GitHubRelease{}, fmt.Errorf("no matching release found")
	}

	return releaseInfo[index], nil

}

func prebuilt(frontendConfig *config.Frontend, globalConfig *config.Global) error {

	github := GitHub{
		RepoURL: frontendConfig.PrebuiltRepo,
	}

	release, err := github.GetLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to get latest release: %v", err)
	}

	utils.ShowStep(fmt.Sprintf("Latest prebuilt release: %s\n", release.TagName))

	tempDir, err := os.MkdirTemp("", "sirius")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}

	defer func() {
		log.Debug("Removing temporary directory", "path", tempDir)
		if err := os.RemoveAll(tempDir); err != nil {
			log.Error("Failed to remove temporary directory", "error", err)
		}
	}()

	assetsToDownload := selectAssetsToDownload(globalConfig.Space, release.Assets)
	if len(assetsToDownload) == 0 {
		return fmt.Errorf("no matching assets found for space: %s", globalConfig.Space)
	}

	for _, asset := range assetsToDownload {

		utils.ShowStep(fmt.Sprintf("Downloading release %s... (this may take a while)\n", asset.Name))

		downloadPath := filepath.Join(tempDir, asset.Name)
		if err := utils.DownloadFile(asset.BrowserDownloadURL, downloadPath); err != nil {
			return fmt.Errorf("failed to download file %s: %v", asset.Name, err)
		}

		extractName := strings.TrimSuffix(asset.Name, ".zip")
		extractPath := filepath.Join(globalConfig.Workdir, extractName)
		if err := utils.Unzip(downloadPath, extractPath); err != nil {
			return fmt.Errorf("failed to extract %s: %v", asset.Name, err)
		}

		utils.ShowStep(fmt.Sprintf("Frontend prepared in %s\n", extractPath))

		if err := asset.WriteSettingsToFile(frontendConfig, globalConfig); err != nil {
			return fmt.Errorf("failed to write settings to file: %v", err)
		}

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
		if strings.Contains(asset.Name, assetName) {
			return []GitHubReleaseAsset{asset}
		}
	}

	return nil
}

func RunFrontendScript(frontendConfig *config.Frontend, globalConfig *config.Global) error {

	utils.ShowTitle("Frontend")
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
