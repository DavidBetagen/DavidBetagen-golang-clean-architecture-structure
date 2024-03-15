package helper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Construct the relative path within the zip file
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		// Create a zip file entry header
		header := &zip.FileHeader{
			Name:   relPath,
			Method: zip.Deflate,
		}

		// If it's a directory, don't compress, just store
		if info.IsDir() {
			header.Method = zip.Store
		}

		// Create the zip file entry
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If it's not a directory, write the file content to the zip file
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
