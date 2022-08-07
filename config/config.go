/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package config

import (
	"encoding/json"
	"fmt"
	"github.com/qascade/yast/tui"
	"github.com/qascade/yast/utils"
	"os"
	//"github.com/tidwall/sjson"
)

var PlayerChoiceFromTui string
var ConfigJsonExists bool

//This function will call interactive tui for taking input of user preferences.
type ConfigBuildSpec struct {
	Player           string `json:"player"`
	TargetPreference string `json:"target-preference"`
	QueryHistory     bool   `json:"query-history"`
}

func NewConfigBuildSpec() *ConfigBuildSpec {
	return &ConfigBuildSpec{}
}

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
			return fmt.Errorf("err %s: could not remove config.json", err)
		}
		ConfigJsonExists = false
		var configBS ConfigBuildSpec
		var configFile *os.File
		configFile, err = CreateConfigJSON()
		if err != nil {
			return fmt.Errorf("err %s: could not create config.json", err)
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
		err := RemoveConfigJSON()
		if err != nil {
			return fmt.Errorf("err %s: could not remove config.json", err)
		}
		ConfigJsonExists = false
		var configBS ConfigBuildSpec
		var configFile *os.File
		configFile, err = CreateConfigJSON()
		if err != nil {
			return fmt.Errorf("err %s: could not create config.json", err)
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
	if targetChange {
		utils.TraceMsg("TODO-Target Preference yet to be added in SetupModel. Defaulting to piratebay for now.")
		return nil
	}
	return nil
}

func GetConfigBSFromSetupModel() (ConfigBuildSpec, error) {
	var configBS ConfigBuildSpec
	utils.TraceMsg("TODO-Fill Config BS using tui-SetupModel")
	err := tui.RenderSetupModelView()
	if err != nil {
		return configBS, fmt.Errorf("err %s: could not render setup model view", err)
	}
	PlayerChoiceFromTui = tui.GetPlayerChoice()
	configBS.Player = PlayerChoiceFromTui
	utils.TraceMsg("TODO-Target Preference yet to be added in SetupModel. Defaulting to piratebay for now.")
	return configBS, nil
}
