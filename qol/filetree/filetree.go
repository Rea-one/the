package filetree

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type FileTree struct {
	The           *widget.Tree
	BaseDirectory string
}

func NewFileTree() *FileTree {
	ft := &FileTree{}
	ft.The = widget.NewTree(
		fileDirSearch,
		fileCheck,
		file2node,
		nodeUpdate,
	)
	return ft
}

func fileDirSearch(path string) []string {
	res, err := getSubDirectories(path)
	if err != nil {
		return nil
	}
	return res
}

func getSubDirectories(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs, nil
}

func fileCheck(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func file2node(branch bool) fyne.CanvasObject {
	return widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
}

func nodeUpdate(path string, branch bool, obj fyne.CanvasObject) {
	label := obj.(*widget.Label)
	if branch {
		label.SetText("📁 " + path)
	} else {
		label.SetText("📄 " + path)
	}
}
