package cmd

import (
	"filemonster/src/internal"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(internal.Start(path), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
