package core

import (
	"encoding/json"
	"os"
	//"github.com/qascade/yast/tui"
	//"github.com/qascade/yast/utils"
	//"github.com/tidwall/sjson"
)

//string where choice from SetupModel will be stored
var PlayerChoice string

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
