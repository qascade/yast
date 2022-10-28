/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package config

import (
	"fmt"
	"github.com/qascade/yast/scraper"
	"os"
)

func SetupYast() error {
	err := CreateYastWorkDir()
	if err != nil {
		return fmt.Errorf("err %s: could not create default yast work dir %s", err, YastWorkDir)
	}
	var configBS ConfigBuildSpec
	configBS, err = GetConfigBSFromSetupModel()
	if err != nil {
		return fmt.Errorf("err %s: could not get config build spec from setup model", err)
	}
	if configBS.Player == "" {
		existingPlayer, err := GetPlayerFromExistingConfig()
		if err != nil || existingPlayer == "" {
			os.Remove(DefaultConfigPath)
			return nil
		}
		configBS.Player = existingPlayer
	}
	//Putting default targetPreference as 1337x.to
	configBS.TargetPreference = scraper.TARGET_1337X
	configFile, err := CreateConfigJSON()
	if err != nil {
		return fmt.Errorf("err %s: could not create config.json", err)
	}

	FillConfigJSON(configFile, &configBS)
	return nil
}
