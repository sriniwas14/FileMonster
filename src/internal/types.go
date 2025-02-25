package internal

type Model struct {
	path   string
	list   *List
	width  int
	height int
}

type List struct {
	items  []File
	cursor int
}

type FileType int

const (
	FileDir FileType = iota
	FileFile
)

type File struct {
	name     string
	itemType FileType
	mimeType string
	size     int
}
