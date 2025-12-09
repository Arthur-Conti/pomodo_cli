package styles

import "github.com/charmbracelet/lipgloss"

type Style struct {
	Help func(strs ...string) string
	Start lipgloss.Color
	End lipgloss.Color
}

func NewStyle() *Style {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	startColor := lipgloss.Color("#04B575")
	endColor := lipgloss.Color("#FF4136")
	return &Style{
		Help: helpStyle,
		Start: startColor,
		End: endColor,
	}
}