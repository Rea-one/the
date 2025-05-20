package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type EditorConfig struct {
	Theme string
}
type EditorCodeField struct {
	width  int
	height int
}

type Editor struct {
	The *widget.TextGrid
}

func NewEditor(config EditorConfig) *Editor {
	ed := &Editor{}
	ed.The = widget.NewTextGrid()
	ed.The.ShowLineNumbers = true
	return ed
}

func (tar *Editor) MinSize() fyne.Size {
	return fyne.NewSize(200, 100)
}

func (tar *Editor) Move(pos fyne.Position) {
	tar.The.Move(pos)
}

func (tar *Editor) Position() fyne.Position {
	return tar.The.Position()
}

func (tar *Editor) Resize(size fyne.Size) {
	tar.The.Resize(size)
}

func (tar *Editor) Size() fyne.Size {
	return tar.The.Size()
}

func (tar *Editor) Hide() {
}

func (tar *Editor) Visible() bool {
	return tar.The.Visible()
}

func (tar *Editor) Show() {
}
