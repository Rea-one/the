package filetree

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type FileTree struct {
	The      *fyne.Container
	Origin   *Origin
	Activate *widget.Button
}

func NewFileTree() *FileTree {
	the := &FileTree{}
	the.The = container.NewWithoutLayout()
	the.Origin = nil
	the.Activate = widget.NewButton("启动文件树", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				// 获取选中目录路径
				path := uri.Path()
				the.Origin.setTree(path)
				the.The.Remove(the.Activate)
				the.The.Add(the.Origin.The)
				the.The.Refresh()
			}
		}, fyne.CurrentApp().NewWindow("Select Folder"))
	})
	the.The.Add(the.Activate)
	return the
}
