package webtorrent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/qascade/yast/config"
	"github.com/qascade/yast/tui"
)

func getPlayerChoiceFromConfig() (string, error) {
	var configBS config.ConfigBuildSpec
	configJsonFile, err := os.Open(config.DefaultConfigPath)
	if err != nil {
		err = fmt.Errorf("error opening config file: %v", err)
		return "", err
	}

	configJson, err := ioutil.ReadAll(configJsonFile)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(configJson, &configBS)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling config file: %v", err)
	}
	return configBS.Player, nil
}

//Need to implement getMagnet from Function Model.
func generateCommand(magnet, defaultPlayer string) string {
	var cmdStr string
	if defaultPlayer == "vlc" {
		cmdStr = fmt.Sprintf("webtorrent %s --%s", magnet, defaultPlayer)
	}
	if defaultPlayer == "web-torrent" {
		cmdStr = fmt.Sprintf("webtorrent %s", magnet)
	}
	return cmdStr
}

func startStream() error {
	defaultPlayerChoice, err := getPlayerChoiceFromConfig()
	chosenMagnet := tui.GetMagnetFromListModel()
	cmdStr := generateCommand(chosenMagnet, defaultPlayerChoice)
	cmd := exec.Command(cmdStr)
	fmt.Printf("Stream is Starting!!\n")
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("error starting stream: %v", err)
	}
	return nil
}
