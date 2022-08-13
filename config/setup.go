/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package config

import (
	"fmt"
	"os"

	"github.com/qascade/yast/scraper"
)

func SetupYast() error {
	err := CreateYastWorkDir()
	if err != nil {
		return fmt.Errorf("err %s: could not create default yast work dir %s", err, YastWorkDir)
	}
	var configFile *os.File
	configFile, err = CreateConfigJSON()
	if err != nil {
		return fmt.Errorf("err %s: could not create config.json", err)
	}
	var configBS ConfigBuildSpec
	configBS, err = GetConfigBSFromSetupModel()
	if err != nil {
		return fmt.Errorf("err %s: could not get config build spec from setup model", err)
	}
	//Putting default targetPreference as piratebay
	configBS.TargetPreference = scraper.TARGET_PIRATEBAY

	FillConfigJSON(configFile, &configBS)
	return nil

}
