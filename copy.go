package main

import (
	"os"
	"path/filepath"
)

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	// Read source file
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// Write to destination file
	return os.WriteFile(dst, data, 0644)
}

// CopyFilesInDir copies all files from source directory to destination directory
// preserving the directory structure
func CopyFilesInDir(source, destination string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(source, path)
			if err != nil {
				return err
			}

			destPath := filepath.Join(destination, relPath)
			destDir := filepath.Dir(destPath)

			err = os.MkdirAll(destDir, 0755)
			if err != nil {
				return err
			}

			return CopyFile(path, destPath)
		}
		return nil
	})
}
