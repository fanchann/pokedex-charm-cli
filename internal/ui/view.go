package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/fanchann/charm-pokedex/helpers"
)

// renders the current view
func (m Model) View() string {
	if m.Loading {
		return ContainerStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				TitleStyle.Render("Pokedex"),
				"",
				fmt.Sprintf("%s Loading Pokemon data...", m.Spinner.View()),
			),
		)
	}

	switch m.State {
	case ListView:
		return m.renderListView()
	case DetailView:
		return m.renderDetailView()
	case SearchView:
		return m.renderSearchView()
	default:
		return "Unknown state"
	}
}

func (m Model) renderListView() string {
	help := HelpStyle.Render("- Enter: View details - /: Search - r: Refresh - q: Quit")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		TitleStyle.Render("Pokedex - Pokemon List"),
		fmt.Sprintf("Total All Pokemon: %d", len(m.List.Items())),
		"",
		m.List.View(),
		help,
	)

	return ContainerStyle.Render(content)
}

func (m Model) renderDetailView() string {
	if m.Err != nil {
		errorContent := lipgloss.JoinVertical(
			lipgloss.Center,
			TitleStyle.Render("Pokedex"),
			"",
			ErrorStyle.Render(fmt.Sprintf("Error: %v", m.Err)),
			"",
			HelpStyle.Render("- ESC: Back to list - /: Search - q: Quit"),
		)
		return ContainerStyle.Render(errorContent)
	}

	pokemon := m.SelectedPokemon

	// basic info
	header := lipgloss.JoinHorizontal(
		lipgloss.Left,
		PokemonNameStyle.Render(fmt.Sprintf("#%d %s", pokemon.ID, helpers.FormatPokemonName(pokemon.Name))),
	)

	// types
	var types []string
	for _, t := range pokemon.Types {
		types = append(types, TypeStyle.Render(strings.ToUpper(t.Type.Name)))
	}
	typesRow := lipgloss.JoinHorizontal(lipgloss.Left, types...)

	// basic stats
	basicInfo := fmt.Sprintf(
		"Height: %.1fm  |  Weight: %.1fkg  |  Base Experience: %d",
		helpers.FormatHeight(pokemon.Height),
		helpers.FormatWeight(pokemon.Weight),
		pokemon.BaseExp,
	)

	// stats
	var statsRows []string
	for _, stat := range pokemon.Stats {
		statName := helpers.FormatStatName(stat.Stat.Name)
		statBar := helpers.RenderStatBar(stat.BaseStat)
		row := lipgloss.JoinHorizontal(
			lipgloss.Left,
			StatNameStyle.Render(fmt.Sprintf("%-15s", statName)),
			StatValueStyle.Render(fmt.Sprintf("%3d", stat.BaseStat)),
			" ",
			statBar,
		)
		statsRows = append(statsRows, row)
	}

	// abilities
	var abilities []string
	for _, ability := range pokemon.Abilities {
		abilityName := helpers.FormatAbilityName(ability.Ability.Name, ability.IsHidden)
		abilities = append(abilities, abilityName)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		TitleStyle.Render("Pokedex - Pokemon Details"),
		"",
		header,
		"",
		HeaderStyle.Render("Types:"),
		typesRow,
		"",
		HeaderStyle.Render("Basic Info:"),
		basicInfo,
		"",
		HeaderStyle.Render("Base Stats:"),
		strings.Join(statsRows, "\n"),
		"",
		HeaderStyle.Render("Abilities:"),
		strings.Join(abilities, ", "),
		"",
		HelpStyle.Render("- ESC/b: Back to list - /: Search - q: Quit"),
	)

	return ContainerStyle.Render(content)
}

func (m Model) renderSearchView() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		TitleStyle.Render("Pokedex - Search Pokemon"),
		"",
		HeaderStyle.Render("Enter Pokemon name or ID:"),
		m.SearchInput.View(),
		"",
		HelpStyle.Render("- Enter: Search - ESC: Cancel"),
	)

	return ContainerStyle.Render(content)
}
