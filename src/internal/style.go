package internal

import "github.com/charmbracelet/lipgloss"

var (
	listStyleItem         = lipgloss.NewStyle().PaddingLeft(1)
	listStyleSelectedItem = lipgloss.NewStyle().Foreground(lipgloss.Color("150")).PaddingLeft(1)
	paneStyleBorder       = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("0"))
	paneStyleBottomBorder = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, true, false).
				BorderForeground(lipgloss.Color("0"))

	titleStyleColor = lipgloss.NewStyle().Foreground(lipgloss.Color("140"))
)
