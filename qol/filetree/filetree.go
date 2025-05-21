package filetree

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type FileTree struct {
	The      *fyne.Container
	Origin   *Origin
	Activate *widget.Button
}

func NewFileTree() *FileTree {
	the := &FileTree{}
	the.Origin = nil
	the.Activate = widget.NewButton("启动文件树", func() {

	})
}
