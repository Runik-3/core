package core

import (
	_ "embed"
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

//go:embed defaultConfig.json
var defaultConfig []byte

type Config struct {
	KindlegenPath string `json:"kindlegenPath"`
	Theme         string `json:"theme"`
}

// Checks for and returns the kindlegen path if it exists. This function checks
// the default install location for Kindle Previewer on windows and macos -- 
// kindle previewer does not have a linux verion.
func getKindlegenBinaryPath() string {
	sysOs := runtime.GOOS

	// C:\Users\[your_username]\AppData\Local\Amazon\Kindle Previewer 3\lib\fc\bin\kindlegen
	if sysOs == "windows" {
		appdata, err := os.UserConfigDir()
		if err != nil {
			return ""
		}

		maybeKindlegenPath:= filepath.Join(appdata, "Local", "Amazon", "Kindle Previewer 3", "lib", "fc", "bin", "kindlegen")
		if FileExists(maybeKindlegenPath) {
			return maybeKindlegenPath
		}
	}

	// /Applications/Kindle Previewer 3.app/Contents/lib/fc/bin/kindlegen
	if sysOs == "darwin" {
		maybeKindlegenPath:= filepath.Join("/", "Applications", "Kindle Previewer 3.app", "Contents", "lib", "fc", "bin", "kindlegen")
		if FileExists(maybeKindlegenPath) {
			return maybeKindlegenPath
		}
	}
	return ""
}

func GetOrCreateConfig(appDir string) (Config, error) {
	configPath := path.Join(appDir, "config.json")

	if !FileExists(configPath) {
		err := os.WriteFile(configPath, defaultConfig, os.ModePerm)
		if err != nil {
			return Config{}, errors.New("There was an issue creating a new configuration file.")
		}
	}

	configData, err := os.ReadFile(configPath)

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return Config{}, err
	}

	if config.KindlegenPath == "" {
		config.KindlegenPath = getKindlegenBinaryPath()
	}

	return config, nil
}

func UpdateConfig(appDir string, c Config) error {
	configPath := path.Join(appDir, "config.json")

	config, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, config, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
