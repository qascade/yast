package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	
	"github.com/qascade/yast/scraper"
	"github.com/qascade/yast/movie"
)

const (
	MOVIE_RESULTS_TITLE = "Movie Results"
	//SERIES_RESULTS_TITLE = "Series Results"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FFFDF5")).
					Background(lipgloss.Color("#25A065")).
					Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
							Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
							Render
)

type listKeyMap struct {
	//Do we need a reload/refresh toggle? 

	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
			toggleTitleBar: key.NewBinding(
					key.WithKeys("T"),
					key.WithHelp("T", "toggle title"),
			),
			toggleStatusBar: key.NewBinding(
					key.WithKeys("S"),
					key.WithHelp("S", "toggle status"),
			),
			togglePagination: key.NewBinding(
					key.WithKeys("P"),
					key.WithHelp("P", "toggle pagination"),
			),
			toggleHelpMenu: key.NewBinding(
					key.WithKeys("H"),
					key.WithHelp("H", "toggle help"),
			),
	}
}

type ListModel struct {
	list          list.Model
	items 	   	  []*list.Item
	keys          *listKeyMap
	delegateKeys  *delegateKeyMap
}

func NewListModel(title string, results scraper.Results) ListModel {
	var (
		//itemGenerator randomItemGenerator
		delegateKeys  = newDelegateKeyMap()
		listKeys      = newListKeyMap()
	)
	// Need for model item list here. 
	
	var items []*list.Item	
	for _, result := range results {
		if queryItem, ok := result.(movie.Movie); ok {
			items = append(items, queryItem)
		}
		if result.(movie.Movie)
	}
	delegate := newItemDelegate(delegateKeys)
	queryItemList := list.New(items, delegate, 0, 0)
	queryItemList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
			//listKeys.insertItem,
		}
	}

	return ListModel{
		list:          queryItemList,
		keys:          listKeys,
		//items stuff was here as well 
		items:         items,
		delegateKeys:  delegateKeys,
	}

}

func (m ListModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}
func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd

		case key.Matches(msg, m.keys.toggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil

		case key.Matches(msg, m.keys.toggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil

		case key.Matches(msg, m.keys.togglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil

		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

		// case key.Matches(msg, m.keys.insertItem):
		// 	m.delegateKeys.remove.SetEnabled(true)
		// 	newItem := m.itemGenerator.next()
		// 	insCmd := m.list.InsertItem(0, newItem)
		// 	statusCmd := m.list.NewStatusMessage(statusMessageStyle("Added " + newItem.Title()))
		// 	return m, tea.Batch(insCmd, statusCmd)
		// }
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m ListModel) View() string{
	return appStyle.Render(m.list.View())
}

//Also need to pass Results struct here for rendering.
func RenderListModelView(title string, results scraper.Results) string {
	if err := tea.NewProgram(NewListModel(title, results)).Start(); err != nil {
			fmt.Errorf("error: not able to render list model")
			os.Exit(1)
	}
}



//Delegate Key Map Functions ------------------//

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			//This will call the View of ResultModel for the chosen queryItem
			case key.Matches(msg, keys.choose):
				return m.NewStatusMessage(statusMessageStyle("You chose " + title))

			}
		}

		return nil
	}

	help := []key.Binding{keys.choose,}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
	}
}