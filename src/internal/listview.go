package internal

import (
	"fmt"
)

const (
	visibleItems = 20
)

func (m Model) listRender() string {
	start := 0
	end := visibleItems

	if m.list.cursor > visibleItems {
		start = m.list.cursor - visibleItems + 1
		end = start + visibleItems
	}

	if end > len(m.list.items) {
		end = len(m.list.items)
	}

	list := paneStyleBottomBorder.Render(m.path) + "\n"
	for i, f := range m.list.items[start:end] {
		if i == m.list.cursor-start {
			list += fmt.Sprintf(listStyleSelectedItem.Render("\ueaf7 %s")+"\n", f.name)
		} else {
			list += fmt.Sprintf(listStyleItem.Render("\ueaf7 %s")+"\n", f.name)
		}
	}

	return paneStyleBorder.Render(list)
}
