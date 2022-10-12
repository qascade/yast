/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"

	"github.com/qascade/yast/scraper"
	"github.com/qascade/yast/tui"
	"github.com/qascade/yast/utils"
)

var (
	CouldNotRemoveConfig     = "err %w: could not remove config.json"
	CouldNotCreateConfig     = "err %w: could not create config.json"
	CouldNotOpenConfig       = "err %w: could not open config.json"
	CouldNotReadConfig       = "err %w: could not read config.json"
	CouldNotUnmarshallConfig = "err %w: could not unmarshall config.json"
	CouldNotRenderSetup      = "err %w: could not render setup model view"
	ProblemWithPlayer        = "err %w: could not get player from existing config"
)

func FillConfigJSON(configFile *os.File, configBS *ConfigBuildSpec) error {
	ConfigJsonExists = true
	encoder := json.NewEncoder(configFile)
	encoder.Encode(configBS)
	return nil
}

func UpdateConfigJSON(playerChange bool, targetChange bool, reset bool) error {
	if reset {
		err := RemoveConfigJSON()
		if err != nil {
			return errors.Errorf(CouldNotRemoveConfig, err)
		}
		var configBS ConfigBuildSpec
		var configFile *os.File
		configFile, err = CreateConfigJSON()
		if err != nil {
			return errors.Errorf(CouldNotCreateConfig, err)
		}
		configBS, err = GetConfigBSFromSetupModel()
		if err != nil {
			return err
		}
		err = FillConfigJSON(configFile, &configBS)
		if err != nil {
			return err
		}
		return nil
	}
	if playerChange {
		//Doing Code duplication here, Should Keep it same as it won't remain same once target-preference is added.
		var configBS ConfigBuildSpec
		var configFile *os.File

		configBS, err := GetConfigBSFromSetupModel()
		if err != nil {
			return err
		}
		//Player quits without choosing a player.
		if configBS.Player == "" {
			configBS.Player, err = GetPlayerFromExistingConfig()
			if err != nil {
				err = errors.Errorf(ProblemWithPlayer, err)
				return err
			}
		} else {
			err := RemoveConfigJSON()
			ConfigJsonExists = false
			if err != nil {
				return errors.Errorf(CouldNotRemoveConfig, err)
			}

			configFile, err = CreateConfigJSON()
			if err != nil {
				return errors.Errorf(CouldNotCreateConfig, err)
			}
		}
		err = FillConfigJSON(configFile, &configBS)
		if err != nil {
			return err
		}
		return nil
	}
	if targetChange {
		utils.TraceMsg("TODO-Target Preference yet to be added in SetupModel. Defaulting to 1337x.to for now.")
		return nil
	}
	return nil
}

func GetConfigBSFromSetupModel() (ConfigBuildSpec, error) {
	var configBS ConfigBuildSpec
	utils.TraceMsg("TODO-Fill Config BS using tui-SetupModel")
	err := tui.RenderSetupModelView()
	if err != nil {
		return configBS, errors.Errorf(CouldNotRenderSetup, err)
	}
	PlayerChoiceFromTui = tui.GetPlayerChoice()
	configBS.Player = PlayerChoiceFromTui
	utils.TraceMsg("TODO-Target Preference yet to be added in SetupModel. Defaulting to 1337x.to for now.")
	configBS.TargetPreference = scraper.TARGET_1337X
	return configBS, nil
}

func GetExistingTargetFromConfig() (string, error) {
	configFile, err := os.Open(DefaultConfigPath)
	if err != nil {
		err = errors.Errorf(CouldNotOpenConfig, err)
		return "", err
	}
	var configBSJson []byte
	configBSJson, err = ioutil.ReadAll(configFile)
	if err != nil {
		err = errors.Errorf(CouldNotReadConfig, err)
		return "", err
	}
	configBS := ConfigBuildSpec{}
	err = json.Unmarshal(configBSJson, &configBS)
	if err != nil {
		err = errors.Errorf(CouldNotUnmarshallConfig, err)
		return "", err
	}
	return configBS.TargetPreference, nil
}

func GetPlayerFromExistingConfig() (string, error) {
	configFile, err := os.Open(DefaultConfigPath)
	if err != nil {
		err = errors.Errorf(CouldNotOpenConfig, err)
		return "", err
	}
	var configBSJson []byte
	configBSJson, err = ioutil.ReadAll(configFile)
	if err != nil {
		err = errors.Errorf(CouldNotReadConfig, err)
		return "", err
	}
	configBS := ConfigBuildSpec{}
	err = json.Unmarshal(configBSJson, &configBS)
	if err != nil {
		err = errors.Errorf(CouldNotUnmarshallConfig, err)
		return "", err
	}
	return configBS.Player, nil
}
