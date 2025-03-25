package core

import (
	_ "embed"
	"encoding/json"
	"errors"
	"os"
	"path"
)

//go:embed defaultConfig.json
var defaultConfig []byte

type Config struct {
	KindleGenPath string `json:"kindlegenPath"`
}

func GetOrCreateConfig(appDir string) (Config, error) {
	configPath := path.Join(appDir, "config.json")

	_, err := os.Stat(configPath)
	// Config file does not exist
	if err != nil {
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

	return config, nil
}
