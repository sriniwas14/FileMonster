package internal

import "os"

func getFiles(path string) []File {
	items, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	files := []File{
		//	{name: "..", itemType: FileDir, mimeType: "folder"},
	}

	for _, item := range items {
		fileType := FileFile
		if item.IsDir() {
			fileType = FileDir
		}

		files = append(files, File{
			itemType: fileType,
			mimeType: item.Type().Perm().String(),
			name:     item.Name(),
		})
	}

	return files
}
