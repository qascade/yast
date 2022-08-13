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
	"github.com/qascade/yast/query"
	"github.com/qascade/yast/scraper"
	"github.com/qascade/yast/tui"

	"github.com/spf13/cobra"
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
	RunE: Search,
}

//Flags for searchCmd
var (
	MovieName  string
	SeriesName string
	movieSet   bool
	seriesSet  bool
)

// For now we will only search for either movie or series one at a time. If both flags set throw error
var bothSet bool

func CheckIfSearchFlagsSet(cmd *cobra.Command, args []string) (movieSet, seriesSet, bothSet bool, err error) {
	if cmd.Flag("movie").Changed {
		movieSet = true
	}
	if cmd.Flag("series").Changed {
		seriesSet = true
	}
	if cmd.Flag("movie").Changed && cmd.Flag("series").Changed {
		bothSet = true
	}
	if !movieSet && !seriesSet {
		err = fmt.Errorf("you must specify either a movie or a series to search for")
	}
	return
}

func Search(cmd *cobra.Command, args []string) error {
	var err error
	movieSet, seriesSet, bothSet, err = CheckIfSearchFlagsSet(cmd, args)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bothSet {
		err = fmt.Errorf("You can only search for either movie or series at a time")
		return err
	}
	if movieSet {
		MovieName = cmd.Flag("movie").Value.String()
		fmt.Println("Searching for movie: ", MovieName)

		defaultTarget, err := config.GetDefaultTarget()
		if err != nil {
			return err
		}

		context := scraper.NewQueryContext("movie", MovieName, defaultTarget)
		Query := core.NewSearchQuery(context)
		var results []scraper.Result
		if err != nil {
			return err
		}

		results, err = Query.Search()
		if err != nil {
			return err
		}

		err = tui.RenderListModelView("", results)
		if err != nil {
			return err
		}
		fmt.Println(results)
	}
	if seriesSet {
		SeriesName = cmd.Flag("series").Value.String()
		fmt.Println("Searching for Series yet to be implemented...")
	}
	return nil
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
