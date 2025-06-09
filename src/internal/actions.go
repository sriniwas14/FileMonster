package internal

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (a *Action) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			a.visible = false
			a.action = ""
			break
		}
	}

	return a, nil
}

func (a *Action) View() string {
	switch a.action {
	case ActionDelete:
		return dialogStyle.Render("Are you sure?\n[y]es / [n]o")
	}

	return ""
}

func (a Action) Init() tea.Cmd {
	return nil
}
