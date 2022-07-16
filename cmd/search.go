/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"fmt"
	"github.com/spf13/cobra"
	//"github.com/qascade/yast/core"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Search, 
}
//Flags for searchCmd
var MovieName string
var SeriesName string
var MovieSet bool
var SeriesSet bool
// For now we will only search for either movie or series one at a time. If both flags set throw error
var BothSet bool

func CheckIfSet(cmd *cobra.Command, args []string) (movieSet, seriesSet, bothSet bool, err error){
	if cmd.Flag("movie").Changed {
		movieSet = true
	}
	if cmd.Flag("series").Changed {
		seriesSet = true
	}
	if cmd.Flag("movie").Changed  && cmd.Flag("series").Changed {
		bothSet = true
	}
	if !movieSet && !seriesSet {
		err = fmt.Errorf("you must specify either a movie or a series to search for")
	}
	return
}

func Search(cmd *cobra.Command, args []string){
	var err error
	MovieSet, SeriesSet, BothSet, err = CheckIfSet(cmd, args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if BothSet {
		fmt.Println("You can only search for either movie or series at a time")
		return
	}
	if MovieSet {
		MovieName = cmd.Flag("movie").Value.String()
		fmt.Println("Searching for movie: ", MovieName)
	}
	if SeriesSet {
		SeriesName = cmd.Flag("series").Value.String()
		fmt.Println("Searching for series: ", SeriesName)
	}
}
	

// func SearchMovie(movieName string, targetSite url) (error){
	
// }

func init() {
	yastCmd.AddCommand(searchCmd)
	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")
	searchCmd.Flags().StringVarP(&MovieName, "movie", "m", "", "name of the movie to be searched")
	searchCmd.Flags().StringVar(&SeriesName, "series", "", "name of the series to be searched")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
