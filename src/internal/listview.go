package internal

import (
	"fmt"
	"strings"
)

const (
	visibleItems = 15
)

func getIcon(filename string, t FileType) string {
	icon := ""
	if t == FileDir {
		return "\ueaf7"
	} else {
		// Get extension
		nameParts := strings.Split(filename, ".")

		if len(nameParts) > 1 {
			icon = FileIcons[nameParts[1]]
		} else {
			icon = FileIcons[strings.ToLower(nameParts[0])]
		}

		if icon == "" {
			icon = "\uea7b"
		}
	}
	return icon
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
				getIcon(f.name, f.itemType),
				f.name,
			)
		} else {
			list += fmt.Sprintf(listStyleItem.Render("%s %s")+"\n", getIcon(f.name, f.itemType), f.name)
		}
	}

	list += padX("", l.width) + "\n"

	return paneStyleBorder.Render(padY(list, l.height))
}
