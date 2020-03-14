package app

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// GetFiles – list files in the folder
func GetFiles(basePath string) ([]string, error) {
	var files []string
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if path != basePath {
			files = append(files, strings.TrimPrefix(path, basePath))
		}
		return nil
	})
	return files, err
}

// ReadFile – safe read of file in a folder
func ReadFile(basePath, filePath string) ([]byte, error) {
	filePath = strings.ReplaceAll(filePath, "/", "")
	return ioutil.ReadFile(path.Join(basePath, filePath))
}
