package infrastructure

import (
	"os"
)

type FileChecker struct {
}

func (f FileChecker) Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func (f FileChecker) IsDirectory(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	mode := fi.Mode()
	return mode.IsDir()
}

func NewFileChecker() *FileChecker {
	return &FileChecker{}
}
