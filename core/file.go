package core

import (
	"log"
	"os"
)

// common filesystem methods

func GetUserConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return configDir
}

func MkdirIfNotExists(path string) {
	appDirErr := os.MkdirAll(path, os.ModePerm)
	if appDirErr != nil {
		log.Fatal(appDirErr)
	}
}
