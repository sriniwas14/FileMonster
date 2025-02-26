package internal

import (
	"fmt"
)

const (
	visibleItems = 15
)

func getIcon(t FileType) string {
	if t == FileDir {
		return "\ueaf7"
	} else {
		return "\uea7b"
	}
}

func (l *List) Render() string {
	start := 0
	end := visibleItems

	if l.cursor > visibleItems {
		start = l.cursor - visibleItems + 1
		end = start + visibleItems
	}

	if end > len(l.items) {
		end = len(l.items)
	}

	list := ""
	if l.showTitle {
		list = paneStyleBottomBorder.Render(padX(l.title, l.width)) + "\n"
	}
	for i, f := range l.items[start:end] {
		if i == l.cursor-start {
			list += fmt.Sprintf(
				listStyleSelectedItem.Render("%s %s")+"\n",
				getIcon(f.itemType),
				f.name,
			)
		} else {
			list += fmt.Sprintf(listStyleItem.Render("%s %s")+"\n", getIcon(f.itemType), f.name)
		}
	}

	list += padX("", l.width) + "\n"

	return paneStyleBorder.Render(padY(list, l.height))
}
