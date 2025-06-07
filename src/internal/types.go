package internal

type Model struct {
	path       string
	list       *List
	width      int
	height     int
	showHelp   bool
	showSearch bool
	showHidden bool
	searchText string
}

type List struct {
	items     []File
	cursor    int
	title     string
	width     int
	height    int
	showTitle bool
}

type FileType int

const (
	FileDir FileType = iota
	FileFile
)

type FileInfo struct {
	filePath    string
	size        int64
	permissions string
	created     string
	modified    string
	width       int
	height      int
}

type File struct {
	name     string
	itemType FileType
	mimeType string
	size     int
}
