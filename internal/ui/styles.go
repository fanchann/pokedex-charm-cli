package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#1078eeff")).
			Padding(0, 1).
			MarginBottom(1)

	HeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000ff")).
			Background(lipgloss.Color("#ffffffff")).
			Padding(0, 1).
			Bold(true).
			MarginBottom(1)

	PokemonNameStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFE66D")).
				Bold(true)

	TypeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#00b7ffff")).
			Padding(0, 1).
			MarginRight(1)

	StatNameStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#0bb677ff")).
			Width(15)

	StatValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD93D")).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Bold(true)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)

	ContainerStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#afa40dff"))
)
