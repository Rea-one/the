package editor

import (
	"github.com/twgh/xcgui/widget"
)

type EditorConfig struct {
	Theme string
}
type EditorCodeField struct {
	width  int
	height int
}

type Editor struct {
	The *widget.Editor
}

func NewEditor(handle int) *Editor {
	var editor = &Editor{}
	editor.The = widget.NewEditor(0, 0, 0, 0, handle)
	return editor
}
