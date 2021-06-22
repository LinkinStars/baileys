// Copyright 2019 LinkinStar
// license that can be found in the LICENSE file.

package util

import (
	"os"
	"path/filepath"
	"strings"
)

// CreateDirIfNotExist create dir recursion
func CreateDirIfNotExist(dir string) error {
	if CheckPathIfNotExist(dir) {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}

// CheckPathIfNotExist return true if file exist
func CheckPathIfNotExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetOnlyFilename return a filename without extension
func GetOnlyFilename(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(path)
	return strings.Split(filename, ext)[0]
}
