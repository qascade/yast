/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/qascade/yast/core"
	"github.com/qascade/yast/utils"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func CallSetup(cmd *cobra.Command, args []string) error {
	err := SetupYast()
	if err != nil {
		return fmt.Errorf("err %s: could not setup yast", err)
	}
	return nil
}
func SetupYast() error {
	err := utils.CreateYastWorkDir()
	if err != nil {
		return fmt.Errorf("err %s: could not create default yast work dir %s", err, utils.YastWorkDir)
	}
	var configFile *os.File
	configFile, err = utils.CreateConfigJSON()
	if err != nil {
		return fmt.Errorf("err %s: could not create config.json", err)
	}
	var configBS core.ConfigBuildSpec
	utils.TraceMsg("TODO-Fill Config BS using tui-SetupModel")
	err = core.FillConfigJSON(configFile, &configBS)
	if err != nil {
		return fmt.Errorf("err %s: could not fill config.json", err)
	}
	return nil
}
func init() {
	yastCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
