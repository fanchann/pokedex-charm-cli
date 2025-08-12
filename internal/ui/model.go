package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/fanchann/charm-pokedex/internal/models"
	"github.com/fanchann/charm-pokedex/internal/service"
)

type ViewState int

const (
	ListView   ViewState = iota // 0
	DetailView                  // 1
	SearchView                  // 2
)

type Model struct {
	State           ViewState
	List            list.Model
	SearchInput     textinput.Model
	Spinner         spinner.Model
	SelectedPokemon models.Pokemon
	Loading         bool
	Err             error
	Width           int
	Height          int
}

func InitialModel() Model {
	// Initialize list
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Pokedex - Pokemon List"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	// Initialize search input
	ti := textinput.New()
	ti.Placeholder = "Search Pokemon by name or ID..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 50

	s := spinner.New()
	s.Spinner = spinner.Hamburger
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return Model{
		State:       ListView,
		List:        l,
		SearchInput: ti,
		Spinner:     s,
		Loading:     true,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		service.FetchPokemonList(0, 150),
		m.Spinner.Tick,
	)
}
