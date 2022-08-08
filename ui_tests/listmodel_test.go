/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/

//Uncomment the test to check UI Rendering Locally.
package ui_tests

// import (
// 	"fmt"
// 	"testing"
//"sync"
//"time"

// "github.com/qascade/yast/movie"
// "github.com/qascade/yast/scraper"
// "github.com/qascade/yast/tui"
//)

// func TestListModel(t *testing.T) error {
// 	 fmt.Println("Running TestListModel")
// var resultsStub scraper.Results
// queryItem1 := movie.Movie{
// 	Name:     "Test Movie 1",
// 	Uploaded: time.Now(),
// 	Magnet:   "random Magnet",
// 	Size:     "random Size",
// 	Seeds:    1,
// 	Uploader: "random Uploader",
// }
// for i := 0; i < 40; i++ {
// 	resultsStub = append(resultsStub, queryItem1)
// }
//Used a wait group to force the test to stop at the rendering go routine to see how the view is rendered.
//WaitGroup is not to be pushed on github as will it will fail the test CI Workflow.
//Uncomment the waitgroup code if you want to see the view rendered.
//wg := sync.WaitGroup{}
// wg.Add(1)
// errc := make(chan error)
// go func() {
//err := tui.RenderListModelView("movie", resultsStub)
//errc <- err
//wg.Done()
//}()
//err := <-errc
//if err != nil {
//t.Error(err)
//}
//wg.Wait()
//}
