/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package cmd

import (
	"fmt"
	"github.com/qascade/yast/config"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup yast for first-time users",
	RunE:  CallSetup,
}

func CallSetup(cmd *cobra.Command, args []string) error {
	err := config.SetupYast()
	if err != nil {
		return fmt.Errorf("err %s: could not setup yast", err)
	}
	return nil
}

func init() {
	yastCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
