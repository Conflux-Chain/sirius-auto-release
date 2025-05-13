package runner

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
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

	if resp.StatusCode != http.StatusOK {
		return GitHubRelease{}, fmt.Errorf("failed to get latest release: %s", resp.Status)
	}

	var releaseInfo GitHubRelease

	if err := json.NewDecoder(resp.Body).Decode(&releaseInfo); err != nil {
		return GitHubRelease{}, fmt.Errorf("failed to decode response: %v", err)
	}

	defer resp.Body.Close()

	return releaseInfo, nil
}
