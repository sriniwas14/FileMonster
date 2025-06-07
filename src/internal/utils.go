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
	if spacesReq < 0 {
		return ""
	}
	return str + strings.Repeat(" ", spacesReq)
}

func padY(str string, height int) string {
	spacesReq := height - len(strings.Split(str, "\n"))
	if spacesReq < 0 {
		return str
	}
	return str + strings.Repeat("\n", spacesReq)
}

func fitY(str string, height int) string {
	lines := strings.Split(str, "\n")

	if len(lines) > height {
		return strings.Join(lines[0:height], "\n")
	}
	return padY(str, height)
}

func getFormattedSize(size float64) string {
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

func genSpaces(count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(" ", count)
}

func applyWidth(content string, width int) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		spacesToAdd := width - len(line)
		lines[i] = fmt.Sprintf("%s%s", line, genSpaces(spacesToAdd))
	}

	return strings.Join(lines, "\n")
}

func centerX(content string, screenWidth int) string {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		sidePad := (screenWidth - len(line)) / 2
		lines[i] = fmt.Sprintf("%s%s%s", genSpaces(sidePad), line, genSpaces(sidePad))
	}

	return strings.Join(lines, "\n")
}

func centerY(content string, screenHeight int) string {
	lineCount := len(strings.Split(content, "\n"))
	linesToAdd := (screenHeight - lineCount) / 2

	v := strings.Repeat("\n", linesToAdd) + content + strings.Repeat("\n", linesToAdd)
	return v
}
