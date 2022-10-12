/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package cmd

import (
	"github.com/pkg/errors"
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
		return errors.Errorf("err %w: could not setup yast", err)
	}
	return nil
}

func init() {
	yastCmd.AddCommand(setupCmd)
}
