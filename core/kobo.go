package core

import (
	"errors"
	"os"
	"path/filepath"
)

func FindKoboDictDir(path string) (string, error) {
	if path == "" {
		return "", errors.New("No device selected.")
	}
	pathCustomDict := filepath.Join(path, ".kobo", "custom-dict")
	pathDict := filepath.Join(path, ".kobo", "dict")
	_, err := os.Stat(pathCustomDict)
	if err == nil {
		return pathCustomDict, nil
	}
	_, err = os.Stat(pathDict)
	if err == nil {
		return pathDict, nil
	}

	return "", errors.New("No dictionary location detected on device.")
}
