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
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Used to change already set user preferences or reset the user preferences to default",
	Long:  `YAST is a TUI utility that will let you stream your favorite movies/tv-series in one command.`,
	RunE:  CallUpdateConfig,
}

func CheckIfConfigFlagSet(cmd *cobra.Command) (bool, bool, bool) {
	if cmd.Flag("reset").Changed {
		return false, false, true
	}
	if cmd.Flag("player").Changed {
		return true, false, false
	}
	if cmd.Flag("target").Changed {
		return false, true, false
	}
	cmd.Help()
	return false, false, false
}

func CallUpdateConfig(cmd *cobra.Command, args []string) error {
	var err error
	playerChangeFlagSet, targetChangeFlagSet, resetFlagSet := CheckIfConfigFlagSet(cmd)
	err = config.UpdateConfigJSON(playerChangeFlagSet, targetChangeFlagSet, resetFlagSet)
	if err != nil {
		return fmt.Errorf("err %s: could not update config", err)
	}
	return err
}

func init() {
	yastCmd.AddCommand(ConfigCmd)

	ConfigCmd.Flags().Bool("player", true, "Change Default Player to use for streaming")
	ConfigCmd.Flags().Bool("reset", true, "Reset the user preferences")
	ConfigCmd.Flags().Bool("target", true, "Change Default Target for searching movies/tv-series")
}
