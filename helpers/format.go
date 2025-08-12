package helpers

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func FormatPokemonName(name string) string {
	return strings.Title(name)
}

func FormatHeight(height int) float64 {
	return float64(height) / 10
}

func FormatWeight(weight int) float64 {
	return float64(weight) / 10
}

func FormatStatName(statName string) string {
	return strings.Title(strings.ReplaceAll(statName, "-", " "))
}

func FormatAbilityName(abilityName string, isHidden bool) string {
	formatted := strings.Title(strings.ReplaceAll(abilityName, "-", " "))
	if isHidden {
		formatted += " (Hidden)"
	}
	return formatted
}

func RenderStatBar(value int) string {
	maxWidth := 20
	maxStat := 100

	filledWidth := int(float64(value) / float64(maxStat) * float64(maxWidth))
	if filledWidth > maxWidth {
		filledWidth = maxWidth
	}

	filled := strings.Repeat("+", filledWidth)
	empty := strings.Repeat("-", maxWidth-filledWidth)

	barStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#4ECDC4"))
	emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#333333"))

	return barStyle.Render(filled) + emptyStyle.Render(empty)
}
