package internal

import (
	"log"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return nil, tea.Quit
		case "up":
			if m.list.cursor > 0 {
				m.list.cursor -= 1
			}
			break
		case "down":
			if m.list.cursor < len(m.list.items)-1 {
				m.list.cursor += 1
			}
			break
		case "left":
			m.path = filepath.Dir(m.path)
			m.list.cursor = 0
			break
		case "right":
			break
		case "enter":
			selected := m.list.items[m.list.cursor]

			if selected.itemType == FileDir {
				m.list.cursor = 0
				m.path += "/" + selected.name
			}
			break
		}
	}

	return m, nil
}

func (m Model) View() string {
	v := ""

	w, h, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	m.width = w
	m.height = h

	log.Println(w, h)

	files := getFiles(m.path)
	m.list.items = files
	v += m.listRender()

	return v
}

func Start(path string) Model {
	return Model{
		path: path,
		list: &List{
			items:  []File{},
			cursor: 0,
		},
	}
}
