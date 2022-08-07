package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/qascade/yast/config"
	"github.com/stretchr/testify/require"
)

func TestConfigUtils(t *testing.T) {
	fmt.Println("Executing Config Test")
	var err error

	//Removing YastWorkDir if already exists
	err = os.RemoveAll(config.YastWorkDir)
	if err != nil {
		t.Errorf("error removing YastWorkDir: %v", err)
	}

	//Testing Constants
	yastWD, err := os.UserHomeDir()
	yastWD = yastWD + "/.yast/"
	if err != nil {
		t.Errorf("error getting home dir: %v", err)
	}
	require.Equal(t, config.YastWorkDir, yastWD)
	require.Equal(t, config.DefaultConfigPath, yastWD+"config.json")

	//Testing Directory/File Creation
	err = config.CreateYastWorkDir()
	require.NoError(t, err, fmt.Sprintf("error creating YastWorkDir: %s", err))
	require.DirExists(t, config.YastWorkDir, fmt.Sprintf("YastWorkDir does not exist: %s", config.YastWorkDir))

	_, err = config.CreateConfigJSON()
	require.NoError(t, err, fmt.Sprintf("error creating config file: %s", err))
	require.FileExists(t, config.DefaultConfigPath, fmt.Sprintf("config.json does not exist: %s", config.DefaultConfigPath))
	err = config.RemoveConfigJSON()
	require.NoError(t, err, fmt.Sprintf("error removing config file: %s", err))

	//Test Cleanup
	err = os.RemoveAll(config.YastWorkDir)
	if err != nil {
		t.Errorf("error removing YastWorkDir: %s", err)
	}
}
