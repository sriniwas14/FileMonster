package cmd

import (
	"filemonster/src/internal"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	p := tea.NewProgram(internal.Start(path), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Println(err)
	}
}
