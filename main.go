package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fanchann/charm-pokedex/internal/ui"
)

func main() {
	p := tea.NewProgram(
		ui.InitialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
