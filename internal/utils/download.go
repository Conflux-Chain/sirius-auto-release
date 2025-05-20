package utils

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func DownloadFile(url string, filePath string) error {

	client := http.Client{
		Timeout: 10 * time.Minute,
	}

	slog.Debug("Downloading file", "url", url, "path", filePath)
	resp, err := client.Get(url)

	if err != nil {
		slog.Error("Failed to download file", "error", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to download file", "status", resp.Status)
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	out, err := os.Create(filePath)

	if err != nil {
		slog.Error("Failed to create file", "error", err)
		return fmt.Errorf("failed to create file: %v", err)
	}

	defer out.Close()

	slog.Debug("Writing to file", "path", filePath)
	slog.Debug("Response size", "size", resp.ContentLength)
	slog.Debug("Response status", "status", resp.Status)
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		slog.Error("Failed to download file", "error", err)
		return fmt.Errorf("failed to download file: %v", err)
	}
	slog.Debug("File downloaded successfully", "path", filePath)

	return nil
}
