/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package cmd

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/qascade/yast/config"
	"github.com/qascade/yast/query"
	"github.com/qascade/yast/scraper"
	"github.com/qascade/yast/tui"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a movie or tv-series that you want to watch",
	Long:  `YAST is a TUI utility that will let you stream your favorite movies/tv-series in one command.`,
	RunE:  Search,
}

// Flags for searchCmd
var (
	MovieName  string
	SeriesName string
	movieSet   bool
	seriesSet  bool

	ErrNoSelectionWhatToWatch = errors.New("you must specify either a movie or a series to search for")
	ErrBothSelectedToWatch    = errors.New("you can only search for either movie or series at a time")
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
		err = ErrNoSelectionWhatToWatch
	}
	return
}

func Search(cmd *cobra.Command, args []string) error {
	var err error
	movieSet, seriesSet, bothSet, err = CheckIfSearchFlagsSet(cmd, args)
	if err != nil {
		return err
	}
	if bothSet {
		return ErrBothSelectedToWatch
	}
	if movieSet {
		MovieName = cmd.Flag("movie").Value.String()
		fmt.Println("Searching for movie: ", MovieName)

		defaultTarget, err := config.GetExistingTargetFromConfig()
		if err != nil {
			return err
		}

		context := scraper.NewQueryContext("movie", MovieName, defaultTarget)
		query := core.NewSearchQuery(context)
		if err != nil {
			return err
		}
		// Adding wait group to wait for search results to come before rendering the list model.
		var wg sync.WaitGroup

		wg.Add(1)
		errc := make(chan error, 1)
		resultc := make(chan scraper.Result, 1)
		resultLen := make(chan int, 1)
		go func() {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			var err error
			var results []scraper.Result
			results, err = query.Search()
			if err != nil {
				errc <- err
			}
			resultLen <- len(results)
			for _, result := range results {
				resultc <- result
			}
			wg.Wait()
		}()
		//Stream will be called by tui
		if err != nil {
			return err
		}
		results := make([]scraper.Result, <-resultLen)
		for i := 0; i < len(results); i++ {
			results[i] = <-resultc
		}
		err = tui.RenderListModelView("", results)
		if err != nil {
			return err
		}
		tui.StartStream()
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
	yastCmd.AddCommand(SearchCmd)
	SearchCmd.Flags().StringVarP(&MovieName, "movie", "m", "", "name of the movie to be searched")
	SearchCmd.Flags().StringVar(&SeriesName, "series", "", "name of the series to be searched")
}
