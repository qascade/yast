/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/qascade/yast/config"
	"github.com/stretchr/testify/require"
)

func TestConfigJson(t *testing.T) {
	fmt.Println("Executing Config JSON Test")

	testConfigFile, err := setupYastEnv()
	if err != nil {
		t.Error(err)
	}

	//Stub for ConfigBuildSpec
	var testConfigBS = config.ConfigBuildSpec{
		Player:           "vlc",
		TargetPreference: "pirate-bay",
		QueryHistory:     true,
	}

	err = config.FillConfigJSON(testConfigFile, &testConfigBS)
	require.NoError(t, err, fmt.Sprintf("error filling config file: %s", err))
	assertFilesMatch(t, config.DefaultConfigPath, TestConfigJsonPath)

	existingPlayerChoice, err := config.GetPlayerFromExistingConfig()
	require.NoError(t, err, fmt.Sprintf("error getting player from existing config: %s", err))
	require.NotEmpty(t, existingPlayerChoice)
	require.Equal(t, existingPlayerChoice, "vlc")

	existingTarget, err := config.GetExistingTargetFromConfig()
	require.NoError(t, err, fmt.Sprintf("error getting target from existing config: %s", err))
	require.NotEmpty(t, existingTarget)
	require.Equal(t, existingTarget, "pirate-bay")

	err = os.RemoveAll(config.YastWorkDir)
	if err != nil {
		panic(err)
	}

}
