/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/qascade/yast/scraper"
)

func SetupYast() error {
	err := CreateYastWorkDir()
	if err != nil {
		return err
	}
	var configFile *os.File
	configFile, err = CreateConfigJSON()
	if err != nil {
		return errors.Errorf("err %w: could not create config.json", err)
	}
	var configBS ConfigBuildSpec
	configBS, err = GetConfigBSFromSetupModel()
	if err != nil {
		return errors.Errorf("err %w: could not get config build spec from setup model", err)
	}
	//Putting default targetPreference as 1337x.to
	configBS.TargetPreference = scraper.TARGET_1337X

	FillConfigJSON(configFile, &configBS)
	return nil

}
