package internal

import (
	"fmt"
	"strings"
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

func (l *List) Render(search string) string {
	visibleItems := l.visibleItems
	start := 0
	end := visibleItems

	if l.cursor+1 > visibleItems {
		// start = l.cursor - visibleItems + 1
		// end = start + visibleItems
		start = l.cursor - visibleItems + 1
		end = l.cursor + 1
	}

	if end > len(l.items) {
		end = len(l.items)
	}

	list := ""
	if l.showTitle {
		list = paneStyleBottomBorder.Render(fitX(l.title, l.width)) + "\n"
	}
	for i, f := range l.items[start:end] {
		// One is for files the other is for folders
		temp := f.name

		// Highlight searched item
		if len(search) > 0 && strings.Contains(temp, search) {
			temp = fmt.Sprintf(listStyleSearchMatch.Render("%s"), temp)
		}

		if i == l.cursor-start {
			temp = fmt.Sprintf(
				listStyleSelectedItem.Render("%s %s")+"\n",
				getIcon(f.name, f.itemType),
				temp,
			)
		} else {
			temp = fmt.Sprintf(listStyleItem.Render("%s %s")+"\n", getIcon(f.name, f.itemType), temp)
		}

		list += temp
	}

	list += padX("", l.width) + "\n"

	return paneStyleBorder.Render(padY(list, l.visibleItems))
}
