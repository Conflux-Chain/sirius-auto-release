package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

func Unzip(src string, out string) error {

	r, err := zip.OpenReader(src)
	log.Debug("Opening zip file", "path", src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}

	defer r.Close()

	for _, f := range r.File {

		fpath := filepath.Join(out, f.Name)

		cleanedFpath := filepath.Clean(fpath)
		cleanedDest := filepath.Clean(out)

		relPath, err := filepath.Rel(cleanedDest, cleanedFpath)

		if err != nil {
			return fmt.Errorf("failed to get relative path: %v", err)
		}

		if strings.HasPrefix(relPath, "..") || relPath == ".." {
			return fmt.Errorf("zip file contains illegal path: %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}

			continue
		}
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}

		rc, err := f.Open()

		if err != nil {
			return fmt.Errorf("failed to open file in zip: %v", err)
		}

		outfile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {

			return fmt.Errorf("failed to create file: %v", err)
		}

		_, err = io.Copy(outfile, rc)

		outfile.Close()
		rc.Close()
		if err != nil {
			return fmt.Errorf("failed to copy file: %v", err)
		}

	}

	return nil
}
