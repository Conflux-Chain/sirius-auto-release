package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteToFile(content []byte, outPath string) error {
	dir := filepath.Dir(outPath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	if err := os.WriteFile(outPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
