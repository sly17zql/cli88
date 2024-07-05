package convert

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckPathIsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		panic(fmt.Sprintf("path does not exist: %v", path))
	} else if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}

	if info.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
}

func ListFiles(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func IsExceptType(fileName string, suffix string) bool {
	parts := strings.Split(fileName, ".")
	lastPart := parts[len(parts)-1]
	if strings.ToLower(lastPart) == suffix {
		return true
	}
	return false
}
