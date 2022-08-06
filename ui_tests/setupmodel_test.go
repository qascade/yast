package ui_tests

import (
	"fmt"
	"testing"
	//"sync"

	"github.com/qascade/yast/core"
	"github.com/qascade/yast/tui"
	"github.com/stretchr/testify/require"
)

func TestSetupModel(t *testing.T) {
	// Used a wait group to force the test to stop at the rendering go routine to see how the view is rendered.
	// WaitGroup is not to be pushed on github as will it will fail the test CI Workflow.
	// Uncomment the waitgroup code if you want to see the view rendered.
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// errc := make(chan error)
	// go func() {
	err := tui.RenderSetupModelView()
	playerChoice := core.PlayerChoice
	// 	errc <- err
	// 	wg.Done()
	// }()
	// err := <-errc
	require.NoError(t, err, fmt.Sprintf("error rendering setup model view: %s", err))
	require.NotEmpty(t, playerChoice, "player choice is empty")
	if playerChoice != "vlc" && playerChoice != "web-torrent" {
		t.Errorf("player choice string:%s not same as the choices available", playerChoice)
	}
	if err != nil {
		t.Error(err)
	}
	// wg.Wait()
}
