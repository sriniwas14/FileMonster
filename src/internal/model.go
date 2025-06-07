package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path"
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
		case "h":
			m.showHidden = !m.showHidden
			break
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
		case "o":
			selected := m.list.items[m.list.cursor]
			fp := path.Join(m.path, selected.name)
			cmd := exec.Command(os.Getenv("EDITOR"), fp)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Println("Editor Started!")
			err := cmd.Run()
			fmt.Println("Editor Stopped!")

			if err != nil {
				fmt.Println("Error launching vim:", err)
			}
			break
		case "right":
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

	h -= 1
	w -= 1

	m.width = w
	m.height = h

	// Render Help
	if m.showHelp {
		return m.RenderDialog(70, 50, HELP_CONTENT)
	}

	files := getFiles(m.list.title, m.showHidden)
	list := m.list
	list.items = files
	list.width = (w / 2) - 2
	list.height = h - 2
	m.list.visibleItems = h - 2
	v = list.Render(m.searchText)

	// File Info View
	selected := list.items[list.cursor]
	fileInfo := getFileInfo(filepath.Join(list.title, selected.name))

	fileInfo.width = (w / 2) - 2
	fileInfo.height = (h / 2) - 4
	infoView := fileInfo.Render()

	subPath := filepath.Join(m.list.title, selected.name)

	preview := ""
	if selected.itemType == FileDir {
		files := getFiles(subPath, m.showHidden)
		lr := List{
			title:        subPath,
			width:        list.width,
			height:       list.height,
			items:        files,
			cursor:       0,
			visibleItems: (h / 2),
			showTitle:    false,
		}
		preview = lr.Render(m.searchText)
	} else {
		contents := getFileContents(subPath)
		preview = contents
		preview = fitX(preview, (w/2)-2)
		preview = fitY(preview, h/2)
		preview = paneStyleBorder.Render(preview)
	}

	left := v
	right := lipgloss.JoinVertical(0, preview, infoView)

	mainbody := lipgloss.JoinHorizontal(0, left, right)

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
