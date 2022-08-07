/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// yastCmd represents the base command when called without any subcommands
// renaming rootCmd to yastCmd for better Context
var yastCmd = &cobra.Command{
	Use:   "yast",
	Short: "Yet Another Streaming Tool",
	Long:  `YAST is a TUI utility that will let you stream your favorite movies/tv-series in one command.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := yastCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yast.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//yastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
