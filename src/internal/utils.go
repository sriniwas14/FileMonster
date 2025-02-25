package internal

import (
	"strings"
)

func padX(str string, width int) string {
	spacesReq := width - len(str)
	spaces := ""

	for range spacesReq {
		spaces += " "
	}

	return str + spaces
}

func padY(str string, height int) string {
	spacesReq := height - len(strings.Split(str, "\n"))
	spaces := ""

	for range spacesReq {
		spaces += "\n"
	}

	return str + spaces
}
