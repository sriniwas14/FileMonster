package internal

type ActionType string

const (
	ActionNew    ActionType = "new"
	ActionRename            = "rename"
	ActionDelete            = "delete"
)

type Action struct {
	action  ActionType
	visible bool
}

type Model struct {
	path         string
	list         *List
	width        int
	height       int
	showHelp     bool
	showSearch   bool
	showHidden   bool
	searchText   string
	actionDialog *Action
}

type List struct {
	visibleItems int
	items        []File
	cursor       int
	title        string
	width        int
	height       int
	showTitle    bool
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
