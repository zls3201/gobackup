package helper

import (
	"os"
	"path"
	fp "path/filepath"
)

// IsExistsPath check path exist
func IsExistsPath(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// MkdirP like mkdir -p
func MkdirP(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0777)
	}
}

// ExplandHome ~/foo -> /home/jason/foo
func ExplandHome(filePath string) string {
	if len(filePath) < 2 {
		return filePath
	}

	if filePath[:2] != "~/" {
		return filePath
	}

	return path.Join(os.Getenv("HOME"), filePath[2:])
}

// ExplandHome ~/foo -> /home/jason/foo
// ExplandHome /a/b/../c -> /a/c
func GetAbsPath(filePath string) (string, error) {
	return fp.Abs(fp.Clean(filePath))
}
