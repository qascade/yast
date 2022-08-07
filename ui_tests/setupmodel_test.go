package ui_tests

import (
	"fmt"
	//"sync"

	"github.com/qascade/yast/tui"
)

func UITestSetupModel() error {
	// Used a wait group to force the test to stop at the rendering go routine to see how the view is rendered.
	// WaitGroup is not to be pushed on github as will it will fail the test CI Workflow.
	// Uncomment the waitgroup code if you want to see the view rendered.
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// errc := make(chan error)
	// go func() {
	err := tui.RenderSetupModelView()
	playerChoice := tui.PlayerChoice
	// 	errc <- err
	// 	wg.Done()
	// }()
	// err := <-errc
	if err != nil {
		return fmt.Errorf("error rendering setup model view: %s", err)
	}
	if playerChoice == "" {
		return fmt.Errorf("player choice is empty")
	}
	if playerChoice != "vlc" && playerChoice != "web-torrent" {
		return fmt.Errorf("player choice string:%s not same as the choices available", playerChoice)
	}
	return nil

}