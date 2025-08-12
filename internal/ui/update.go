package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/fanchann/charm-pokedex/internal/globals"
	"github.com/fanchann/charm-pokedex/internal/models"
	"github.com/fanchann/charm-pokedex/internal/service"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.List.SetWidth(msg.Width - 5)
		m.List.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		switch m.State {
		case ListView:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				if selectedItem := m.List.SelectedItem(); selectedItem != nil {
					if pokemon, ok := selectedItem.(models.PokemonListItem); ok {
						m.Loading = true
						m.State = DetailView
						return m, tea.Batch(
							service.FetchPokemonDetail(pokemon.Name),
							m.Spinner.Tick,
						)
					}
				}
			case "/":
				m.State = SearchView
				m.SearchInput.Focus()
				return m, textinput.Blink
			case "r":
				m.Loading = true
				return m, tea.Batch(
					service.FetchPokemonList(0, 150),
					m.Spinner.Tick,
				)
			}

		case DetailView:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "esc", "b":
				m.State = ListView
				return m, nil
			case "/":
				m.State = SearchView
				m.SearchInput.Focus()
				return m, textinput.Blink
			}

		case SearchView:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.State = ListView
				m.SearchInput.Blur()
				return m, nil
			case "enter":
				query := m.SearchInput.Value()
				if query != "" {
					m.Loading = true
					m.State = DetailView
					m.SearchInput.Blur()
					return m, tea.Batch(
						service.SearchPokemon(query),
						m.Spinner.Tick,
					)
				}
				m.State = ListView
				m.SearchInput.Blur()
				return m, nil
			}
		}

	case globals.PokemonListMsg:
		m.Loading = false
		items := make([]list.Item, len(msg.Pokemons))
		for i, pokemon := range msg.Pokemons {
			items[i] = pokemon
		}
		m.List.SetItems(items)
		return m, nil

	case globals.PokemonDetailMsg:
		m.Loading = false
		m.SelectedPokemon = msg.Pokemon
		m.Err = nil
		return m, nil

	case globals.ErrorMsg:
		m.Loading = false
		m.Err = msg.Err
		return m, nil

	case spinner.TickMsg:
		if m.Loading {
			m.Spinner, cmd = m.Spinner.Update(msg)
			return m, cmd
		}
	}

	// update component based state
	switch m.State {
	case ListView:
		m.List, cmd = m.List.Update(msg)
	case SearchView:
		m.SearchInput, cmd = m.SearchInput.Update(msg)
	}

	return m, cmd
}
