package internal

import (
	"fmt"
	"strings"
)

var stepSizeMap = map[int]string{
	0: "bytes",
	1: "kb",
	2: "mb",
	3: "gb",
	4: "tb",
}

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

func fitY(str string, height int) string {
	lines := strings.Split(str, "\n")

	if len(lines) > height {
		return strings.Join(lines[0:height], "\n")
	}
	return padY(str, height)
}

func getPrettySize(size float64) string {
	value := size
	step := 0

	for range 10 {
		if value < 1000 {
			break
		}
		value = size / 1000
		step++
	}

	return fmt.Sprintf("%.2f %s", value, stepSizeMap[step])
}
