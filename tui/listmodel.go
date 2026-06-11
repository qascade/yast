/*
Copyright © 2022 Shubh Karman Singh <sksingh2211@gmail.com>
All rights reserved.
This Project is under BSD-3 License Clause.
Look at License for more detail.
*/
package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/qascade/yast/movie"
	"github.com/qascade/yast/scraper"
)

const (
	MOVIE_RESULTS_TITLE = "Movie Results"
	//SERIES_RESULTS_TITLE = "Series Results"
)

var qualityLabels = map[string]string{
	"":        "All",
	"4K":      "4K",
	"1080p":   "1080p",
	"720p":    "720p",
	"480p":    "480p",
	"360p":    "360p",
	"Unknown": "Unknown",
}

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
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	filter1080p      key.Binding
	filter720p       key.Binding
	filter4K         key.Binding
	filterAll        key.Binding
	clearFilter      key.Binding
	sortBySeeds      key.Binding
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
		filter1080p: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", "filter 1080p"),
		),
		filter720p: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", "filter 720p"),
		),
		filter4K: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("3", "filter 4K"),
		),
		filterAll: key.NewBinding(
			key.WithKeys("0"),
			key.WithHelp("0", "show all"),
		),
		clearFilter: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "clear filters"),
		),
		sortBySeeds: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "sort by seeds"),
		),
	}
}

type ListModel struct {
	list          list.Model
	items         []list.Item
	keys          *listKeyMap
	delegateKeys  *delegateKeyMap
	qualityFilter string
	sortBySeeds   bool
}

//title argument will be used later once tv-series is implemented.
func NewListModel(title string, results scraper.Results) ListModel {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)
	// Need for model item list here.

	var items []list.Item
	for _, result := range results {
		if queryItem, ok := result.(movie.Movie); ok {
			items = append(items, queryItem)
		}
	}
	delegate := newItemDelegate(delegateKeys)
	queryItemList := list.New(items, delegate, 0, 0)
	//This will have to be handled differently once series type is added to yast
	queryItemList.Title = MOVIE_RESULTS_TITLE
	queryItemList.Styles.Title = titleStyle
	queryItemList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
			listKeys.filter1080p,
			listKeys.filter720p,
			listKeys.filter4K,
			listKeys.filterAll,
			listKeys.sortBySeeds,
			listKeys.clearFilter,
		}
	}

	return ListModel{
		list:          queryItemList,
		keys:          listKeys,
		items:         items,
		delegateKeys:  delegateKeys,
		qualityFilter: "",
		sortBySeeds:   false,
	}

}

func (m ListModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m ListModel) View() string {
	return appStyle.Render(m.list.View())
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
		case tea.Quit():
			return nil, tea.Quit
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

		case key.Matches(msg, m.keys.filter1080p):
			m.qualityFilter = "1080p"
			m.applyFilter()
			m.list.SetShowStatusBar(true)
			return m, m.list.NewStatusMessage(statusMessageStyle(m.filterStatus()))

		case key.Matches(msg, m.keys.filter720p):
			m.qualityFilter = "720p"
			m.applyFilter()
			m.list.SetShowStatusBar(true)
			return m, m.list.NewStatusMessage(statusMessageStyle(m.filterStatus()))

		case key.Matches(msg, m.keys.filter4K):
			m.qualityFilter = "4K"
			m.applyFilter()
			m.list.SetShowStatusBar(true)
			return m, m.list.NewStatusMessage(statusMessageStyle(m.filterStatus()))

		case key.Matches(msg, m.keys.filterAll):
			m.qualityFilter = ""
			m.applyFilter()
			m.list.SetShowStatusBar(true)
			return m, m.list.NewStatusMessage(statusMessageStyle(m.filterStatus()))

		case key.Matches(msg, m.keys.sortBySeeds):
			m.sortBySeeds = !m.sortBySeeds
			m.applyFilter()
			m.list.SetShowStatusBar(true)
			statusMsg := "Sort: default"
			if m.sortBySeeds {
				statusMsg = "Sort: Seeds ↓"
			}
			return m, m.list.NewStatusMessage(statusMessageStyle(statusMsg))

		case key.Matches(msg, m.keys.clearFilter):
			if m.list.FilterState() == list.FilterApplied || m.qualityFilter != "" || m.sortBySeeds {
				m.list.ResetFilter()
				m.qualityFilter = ""
				m.sortBySeeds = false
				m.list.SetItems(m.items)
				return m, m.list.NewStatusMessage(statusMessageStyle("Filter cleared"))
			}
		}
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

//This method is used to render ListModel View once query results are aquired by the scrper.
//Also need to pass Results struct here for rendering.
func RenderListModelView(title string, results scraper.Results) (err error) {
	if err := tea.NewProgram(NewListModel(title, results)).Start(); err != nil {
		err = fmt.Errorf("error: not able to render list model")
		return err
	}
	return nil
}

// applyFilter recomputes the displayed item list based on the current quality filter
// and sort settings, then updates the underlying list.Model.
func (m *ListModel) applyFilter() {
	filtered := make([]list.Item, 0, len(m.items))
	for _, item := range m.items {
		mv, ok := item.(movie.Movie)
		if !ok {
			filtered = append(filtered, item)
			continue
		}
		if m.qualityFilter == "" || mv.Quality == m.qualityFilter {
			filtered = append(filtered, item)
		}
	}
	if m.sortBySeeds {
		m.sortItemsBySeeds(filtered)
	}
	m.list.SetItems(filtered)
}

// sortItemsBySeeds sorts items in-place by seed count (descending).
func (m *ListModel) sortItemsBySeeds(items []list.Item) {
	for i := 0; i < len(items); i++ {
		for j := i + 1; j < len(items); j++ {
			si, oki := items[i].(movie.Movie)
			sj, okj := items[j].(movie.Movie)
			if !oki || !okj {
				continue
			}
			seedI := parseSeedCount(si.Seeds)
			seedJ := parseSeedCount(sj.Seeds)
			if seedJ > seedI {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
}

// parseSeedCount converts a seed string (e.g. "4521") to an int.
func parseSeedCount(s string) int {
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		} else {
			break
		}
	}
	return n
}

// filterStatus returns a human-readable status line showing active filters.
func (m *ListModel) filterStatus() string {
	label := qualityLabels[m.qualityFilter]
	if label == "" {
		label = "All"
	}
	status := "Filter: " + label
	if m.sortBySeeds {
		status += " | Sort: Seeds ↓"
	}
	return status
}

//--------------Deletegate Key Map-------------------//

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		//var title string
		var magnet string
		if i, ok := m.SelectedItem().(movie.Movie); ok {
			//title = i.Name
			magnet = i.Magnet
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			//This case will be used to call ResultModel View for the selected result
			//For now just printing the Title
			case key.Matches(msg, keys.choose):
				//webtorrent api call
				SetMagnetChoice(magnet)
				// var wg sync.WaitGroup
				// wg.Add(1)
				// go func() {
				// 	defer wg.Done()
				// 	err := StartStream()
				// 	if err != nil {
				// 		panic(err)
				// 	}
				// }()
				// wg.Wait()
				return tea.Quit
				//return m.NewStatusMessage(statusMessageStyle("You chose " + title))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose}

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
