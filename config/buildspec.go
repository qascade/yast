/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package config

var PlayerChoiceFromTui string
var ConfigJsonExists bool

// This function will call interactive tui for taking input of user preferences.
type ConfigBuildSpec struct {
	Player           string `json:"player"`
	TargetPreference string `json:"target-preference"`
	QueryHistory     bool   `json:"query-history"`
}

func NewConfigBuildSpec() *ConfigBuildSpec {
	return &ConfigBuildSpec{}
}
