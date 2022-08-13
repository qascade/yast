package tui

import (
	"fmt"
	"os/exec"
)

//Need to implement getMagnet from Function Model.
func GenerateCommand(magnet, defaultPlayer string) string {
	var cmdStr string
	if defaultPlayer == "vlc" {
		cmdStr = fmt.Sprintf("webtorrent %s --%s", magnet, defaultPlayer)
	}
	if defaultPlayer == "web-torrent" {
		cmdStr = fmt.Sprintf("webtorrent %s", magnet)
	}
	return cmdStr
}

func StartStream() error {
	chosenMagnet := GetMagnetFromListModel()
	cmdStr := GenerateCommand(chosenMagnet, GetPlayerChoice())
	cmd := exec.Command(cmdStr)
	fmt.Printf("Stream is Starting!!\n")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error starting stream: %v", err)
	}
	return nil
}
