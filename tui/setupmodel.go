/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/termenv"
)

var (
	term   = termenv.EnvColorProfile()
	subtle = makeFgStyle("241")
	dot    = colorFg(" • ", "236")
)

type SetupModel struct {
	PlayerChoice string
	Choice       int
	Chosen       bool
	Quitting     bool
}

func (m SetupModel) Init() tea.Cmd {
	return nil
}

// The main view, which just calls the appropriate sub-view
func (m SetupModel) View() string {
	var s string
	if m.Quitting {
		tea.Quit()
	}
	if !m.Chosen {
		s = choicesView(m)
	} else {
		s = chosenView(m)
	}
	return indent.String("\n"+s+"\n\n", 2)
}

func choicesView(m SetupModel) string {
	c := m.Choice

	tpl := "Welcome to YAST!! 🥳\n\n"
	tpl += "%s\n\n"
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"Please Choose your preferred default player to stream\n \n%s\n%s",
		checkbox("web-torrent default player", c == 0),
		checkbox("vlc", c == 1),
	)
	return fmt.Sprintf(tpl, choices)
}
func chosenView(m SetupModel) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = "You Chose WebTorrent Player. Happy Streaming 🙂"
		PlayerChoice = "web-torrent"
	case 1:
		msg = "You Chose VLC. Happy Streaming 🙂"
		PlayerChoice = "vlc"
	}
	m.Quitting = true
	return msg + "\n\nPress q to Quit."
}
func checkbox(label string, checked bool) string {
	if checked {
		return colorFg("[x] "+label, "212")
	}
	return fmt.Sprintf("[ ] %s", label)
}

//update Function for updating chosen item movement based on triggers (Arrow Keys)
func updateChoices(msg tea.Msg, m SetupModel) tea.Model {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice += 1
			if m.Choice > 1 {
				m.Choice = 1
			}
		case "k", "up":
			m.Choice -= 1
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
			return m
		}
	}
	return m
}

//Main Update Function for SetupModel
func (m SetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if !m.Chosen {
		return updateChoices(msg, m), nil
	}
	m.Quitting = true
	return updateChosen(m), nil
}

//Update Function after the user has chosen the default player to stream
func updateChosen(m SetupModel) tea.Model {
	//Quit After Player has made his choice
	m.Quitting = true
	return m
}

//Function Call to Render SetupModel
func RenderSetupModelView() error {
	setupModel := SetupModel{}
	if err := tea.NewProgram(setupModel).Start(); err != nil {
		err = fmt.Errorf("error: not able to render setup model")
		return err
	}
	return nil
}

//Coloring Function Definitions

// Color a string's foreground with the given value.
func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}
