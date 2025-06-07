package internal

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func getFileInfo(filePath string) *FileInfo {
	info, err := os.Stat(filePath)
	if err != nil {
		log.Println(err)
		if errors.Is(err, os.ErrPermission) {
			return &FileInfo{}
		}
		if errors.Is(err, os.ErrNotExist) {
			return &FileInfo{}
		}
	}

	mode := info.Mode().String()

	return &FileInfo{
		permissions: mode,
		size:        info.Size(),
		created:     info.ModTime().String(),
		modified:    info.ModTime().String(),
	}
}

func (f FileInfo) Render() string {
	value := ""
	value += fmt.Sprintf(
		titleStyleColor.Render(fitX(" Permissions", f.width))+"\n %s\n",
		f.permissions,
	)
	value += fmt.Sprintf(
		titleStyleColor.Render(" Size")+"\n %s\n",
		getFormattedSize(float64(f.size)),
	)
	value += fmt.Sprintf(titleStyleColor.Render(" Created")+"\n %s\n", f.created)
	value += fmt.Sprintf(titleStyleColor.Render(" Last Modified")+"\n %s\n", f.modified)

	return paneStyleBorder.Render(fitY(value, f.height))
}
