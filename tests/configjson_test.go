package test

import (
	"fmt"
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
		TargetPreference: "piratebay",
		QueryHistory:     true,
	}

	err = config.FillConfigJSON(testConfigFile, &testConfigBS)
	require.NoError(t, err, fmt.Sprintf("error filling config file: %s", err))
	assertFilesMatch(t, config.DefaultConfigPath, TestConfigJsonPath)
}