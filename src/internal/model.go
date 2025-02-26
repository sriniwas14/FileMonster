package internal

import (
	"log"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
			return m, tea.Quit
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
			m.list.title = filepath.Dir(m.list.title)
			m.list.cursor = 0
			break
		case "right":
			break
		case "enter":
			selected := m.list.items[m.list.cursor]

			if selected.itemType == FileDir {
				m.list.cursor = 0
				m.list.title += "/" + selected.name
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

	files := getFiles(m.list.title)
	list := m.list
	list.items = files
	list.width = (w / 2) - 2
	list.height = ((h / 3) * 2) - 2
	v = list.Render()

	selected := list.items[list.cursor]
	fileInfo := getFileInfo(filepath.Join(list.title, selected.name))
	if fileInfo == nil {
		log.Println("error reading file information")
	}

	fileInfo.width = (w / 3) - 2
	fileInfo.height = (h / 3) - 1
	infoView := fileInfo.Render()

	subPath := filepath.Join(m.list.title, selected.name)

	r := ""
	if selected.itemType == FileDir {
		files := getFiles(subPath)
		lr := List{
			title:     subPath,
			width:     list.width,
			height:    list.height,
			items:     files,
			cursor:    0,
			showTitle: false,
		}
		r = lr.Render()
	} else {
		contents := getFileContents(subPath)
		r = contents
		r += padX("", (w/2)-2)
		r = fitY(r, ((h/3)*2)-2)
		r = paneStyleBorder.Render(r)
	}

	top := lipgloss.JoinHorizontal(0, v, r)
	bottom := infoView

	return lipgloss.JoinVertical(0, top, bottom)
}

func Start(path string) Model {
	return Model{
		path: path,
		list: &List{
			items:     []File{},
			cursor:    0,
			title:     path,
			showTitle: true,
		},
	}
}
