package test

import(
	"os"
	"fmt"
	"testing"
	
	"github.com/qascade/yast/utils"
	"github.com/stretchr/testify/require"
)

func TestUtils(t *testing.T){
	fmt.Println("Executing Utils Test")
	var err error
	
	//Removing YastWorkDir if already exists
	err = os.RemoveAll(utils.YastWorkDir)
	if err != nil {
		t.Errorf("error removing YastWorkDir: %v", err)
	}

	//Testing Constants
	yastWD, err := os.UserHomeDir() 
	yastWD = yastWD + "/.yast/"
	if err != nil {
		t.Errorf("error getting home dir: %v", err)
	}
	require.Equal(t, utils.YastWorkDir, yastWD)
	require.Equal(t, utils.DefaultConfigPath, yastWD + "config.json")

	//Testing Directory/File Creation
	err = utils.CreateYastWorkDir()
	require.NoError(t, err, fmt.Sprintf("error creating YastWorkDir: %s", err))
	require.DirExists(t, utils.YastWorkDir, fmt.Sprintf("YastWorkDir does not exist: %s", utils.YastWorkDir))

	_ , err = utils.CreateConfigJSON()
	require.NoError(t, err, fmt.Sprintf("error creating config file: %s", err))
	require.FileExists(t,utils.DefaultConfigPath, fmt.Sprintf("config.json does not exist: %s", utils.DefaultConfigPath))
	err = utils.RemoveConfigJSON()
	require.NoError(t, err, fmt.Sprintf("error removing config file: %s", err))
	
	//Test Cleanup
	err = os.RemoveAll(utils.YastWorkDir)
	if err != nil {
		t.Errorf("error removing YastWorkDir: %s", err)
	}
}