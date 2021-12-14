package util

import (
	"archive/zip"
	"io"
	"os"
)

// zipファイルを生成する処理
func GenerateZipFile(filePath string, filePaths []string) error {
	archive, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	for _, filePath := range filePaths {
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		w, err := zipWriter.Create(filePath)
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, f); err != nil {
			return err
		}
	}

	zipWriter.Close()

	return nil
}
