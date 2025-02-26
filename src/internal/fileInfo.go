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
		if errors.Is(err, os.ErrPermission) {
			log.Println(err)
			return nil
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
		titleStyleColor.Render(padX(" Permissions", f.width))+"\n %s\n",
		f.permissions,
	)
	value += fmt.Sprintf(titleStyleColor.Render(" Size")+"\n %s\n", getPrettySize(float64(f.size)))
	value += fmt.Sprintf(titleStyleColor.Render(" Created")+"\n %s\n", f.created)
	value += fmt.Sprintf(titleStyleColor.Render(" Last Modified")+"\n %s\n", f.modified)

	return paneStyleBorder.Render(padY(value, f.height))
}
