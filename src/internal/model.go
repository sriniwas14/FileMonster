package internal

import (
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) handleTextInput(key string) {
	switch key {
	case "backspace":
		if len(m.searchText) > 0 {
			m.searchText = m.searchText[:len(m.searchText)-1]
		}
		break
	default:
		m.searchText += key
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		default:
			if m.showSearch {
				m.handleTextInput(msg.String())
				break
			}
		case "?":
			m.showHelp = true
			return m, nil
		case "/":
			m.showSearch = true
			return m, nil
		case "esc":
			m.showHelp = false
			m.showSearch = false
			m.searchText = ""
			return m, nil
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
				m.list.title = filepath.Join(m.list.title, selected.name)
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

	// Render Help
	if m.showHelp {
		return m.RenderDialog(70, 50, HELP_CONTENT)
	}

	files := getFiles(m.list.title)
	list := m.list
	list.items = files
	list.width = (w / 2) - 2
	list.height = ((h / 3) * 2) - 2
	v = list.Render(m.searchText)

	// File Info View
	selected := list.items[list.cursor]
	fileInfo := getFileInfo(filepath.Join(list.title, selected.name))

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
		r = lr.Render(m.searchText)
	} else {
		contents := getFileContents(subPath)
		r = contents
		r += padX("", (w/2)-2)
		r = fitY(r, ((h/3)*2)-2)
		r = paneStyleBorder.Render(r)
	}

	top := lipgloss.JoinHorizontal(0, v, r)
	bottom := infoView

	mainbody := lipgloss.JoinVertical(0, top, bottom)

	if !m.showSearch {
		return mainbody
	}

	commandPane := "Search: " + m.searchText

	return lipgloss.JoinVertical(0, mainbody, commandPane)
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
		showHelp: false,
	}
}
