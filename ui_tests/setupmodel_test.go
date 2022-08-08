/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/

//Uncomment the code to test UI locally.

package ui_tests

// import (
// 	"fmt"
//"sync"

// 	"github.com/qascade/yast/tui"
// )

// func TestSetupModel(t *testing.T) error {
// Used a wait group to force the test to stop at the rendering go routine to see how the view is rendered.
// WaitGroup is not to be pushed on github as will it will fail the test CI Workflow.
// Uncomment the waitgroup code if you want to see the view rendered.
// wg := sync.WaitGroup{}
// wg.Add(1)
// errc := make(chan error)
// go func() {
// err := tui.RenderSetupModelView()
// playerChoice := tui.PlayerChoice
// 	errc <- err
// 	wg.Done()
// }()
// err := <-errc
// if err != nil {
// 	t.Errorf("error rendering setup model view: %s", err)
// }
// if playerChoice == "" {
// 	t.Errorf("player choice is empty")
// }
// if playerChoice != "vlc" && playerChoice != "web-torrent" {
// 	t.Errorf("player choice string:%s not same as the choices available", playerChoice)
// }
// return nil
//}
