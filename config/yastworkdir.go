package config

import (
	"os"
)

var (
	YastWorkDir       string
	DefaultConfigPath string
)

func CreateYastWorkDir() error {
	err := os.MkdirAll(YastWorkDir, 0755)
	if err != nil {
		return err
	}
	return nil
}

func CreateConfigJSON() (*os.File, error) {
	configFile, err := os.Create(DefaultConfigPath)
	if err != nil {
		return nil, err
	}
	return configFile, nil
}

func RemoveConfigJSON() error {
	err := os.Remove(DefaultConfigPath)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	var userHomeDir string
	var err error
	userHomeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	YastWorkDir = userHomeDir + "/.yast/"
	DefaultConfigPath = YastWorkDir + "config.json"
}
