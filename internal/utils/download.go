package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

func DownloadFile(url string, filePath string) error {

	client := http.Client{
		Timeout: 10 * time.Minute,
	}

	log.Debug("Downloading file", "url", url, "path", filePath)
	resp, err := client.Get(url)

	if err != nil {
		log.Error("Failed to download file", "error", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Failed to download file", "status", resp.Status)
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	out, err := os.Create(filePath)

	if err != nil {
		log.Error("Failed to create file", "error", err)
		return fmt.Errorf("failed to create file: %v", err)
	}

	defer out.Close()

	log.Debug("Writing to file", "path", filePath)
	log.Debug("Response size", "size", resp.ContentLength)
	log.Debug("Response status", "status", resp.Status)
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Error("Failed to download file", "error", err)
		return fmt.Errorf("failed to download file: %v", err)
	}
	log.Debug("File downloaded successfully", "path", filePath)

	return nil
}
