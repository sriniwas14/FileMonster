package internal

import "github.com/charmbracelet/lipgloss"

var (
	listStyleItem         = lipgloss.NewStyle().PaddingLeft(1)
	listStyleSelectedItem = lipgloss.NewStyle().Foreground(lipgloss.Color("150")).PaddingLeft(1)
	paneStyleBorder       = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	paneStyleBottomBorder = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, true, false)
)
