/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved. 
This Project is under BSD-3 License Clause. 
Look at License for more detail. 
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/qascade/yast/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Used to change already set user preferences or reset the user preferences to default",
	Long: `YAST is a TUI utility that will let you stream your favorite movies/tv-series in one command.`,
	RunE: CallUpdateConfig, 
}

func CallSetup(cmd *cobra.Command, args []string) error {
	err := config.SetupYast()
	if err != nil {
		return fmt.Errorf("err %s: could not setup yast", err)
	}
	return nil
}

func CheckIfConfigFlagSet(cmd *cobra.Command, args []string) (bool, bool, bool) {
	if cmd.Flag("reset").Changed {
		return false, false, true
	}
	if cmd.Flag("player").Changed {
		return true, false, false
	}
	if cmd.Flag("target").Changed {
		return false, true, false
	}
	return false, false, false
}

func CallUpdateConfig(cmd *cobra.Command, args []string) error {
	var err error
	playerChangeFlagSet, targetChangeFlagSet, resetFlagSet := CheckIfConfigFlagSet(cmd, args)
	err = config.UpdateConfigJSON(playerChangeFlagSet,targetChangeFlagSet, resetFlagSet)
	if err != nil {
		return fmt.Errorf("err %s: could not update config", err)
	}
	return err
}	

func init() {
	yastCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")
	configCmd.Flags().StringVar(&config.PlayerChoiceFromTui, "player", "", "Player to use for streaming")
	configCmd.Flags().Bool("reset", false, "Reset the user preferences")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
