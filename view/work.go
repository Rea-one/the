package view

import (
	"the/qol/editor"
	"the/qol/filetree"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Work struct {
	The *fyne.Container
}

func NewWork() *Work {
	return &Work{The: container.NewBorder(
		nil,
		nil,
		filetree.NewFileTree().The,
		nil,
		editor.NewEditor(editor.EditorConfig{Theme: "dark"}).The,
	)}

}
