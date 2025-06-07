package internal

import (
	"errors"
	"log"
	"os"
)

func getFiles(path string, showHidden bool) []File {
	items, err := os.ReadDir(path)

	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			log.Println(err)
			return []File{}
		}

	}

	files := []File{}

	for _, item := range items {
		fileType := FileFile
		if item.IsDir() {
			fileType = FileDir
		}
		fName := item.Name()

		if fName[0] == '.' && !showHidden {
			continue
		}

		files = append(files, File{
			itemType: fileType,
			mimeType: item.Type().Perm().String(),
			name:     item.Name(),
		})
	}

	return files
}

func getFileContents(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer file.Close()
	n := 1536
	buffer := make([]byte, n)

	bytesRead, err := file.Read(buffer)

	return string(buffer[:bytesRead])
}
