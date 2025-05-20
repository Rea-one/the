package editor

import (
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
	return ed
}
