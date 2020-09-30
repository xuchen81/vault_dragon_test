package config

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/viper"
)

var Vi *viper.Viper

func GetConfigFilePath() (string, error) {
	configPathPriority := []string{"settings", "/usr/local/etc/vaultdragon"}
	for _, p := range configPathPriority {
		configFile := path.Join(p, "local.yml")
		if os.Getenv("VD_ENV") != "local" {
			configFile = path.Join(p, "config.yml")
		}
		if _, err := os.Stat(configFile); !os.IsNotExist(err) {
			return configFile, nil
		}
	}
	return "", errors.New("Missing config file")
}
