package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Wait struct {
	The    *fyne.Container
	starts []*fyne.CanvasObject
}

func NewWait() *Wait {
	con := &Wait{}
	welcome := canvas.NewText("欢迎来到the", color.Black)
	welcome.TextSize = 20
	welcome.TextStyle = fyne.TextStyle{Bold: true}

	con.The = container.New(layout.NewHBoxLayout(), layout.NewSpacer(), layout.NewSpacer())

	return con
}

func (tar *Wait) MinSize() fyne.Size {
	return fyne.NewSize(200, 100)
}

func (tar *Wait) Move(pos fyne.Position) {
	tar.The.Move(pos)
}

func (tar *Wait) Position() fyne.Position {
	return tar.The.Position()
}

func (tar *Wait) Resize(size fyne.Size) {
	tar.The.Resize(size)
}

func (tar *Wait) Size() fyne.Size {
	return tar.The.Size()
}

func (tar *Wait) Hide() {
}

func (tar *Wait) Visible() bool {
	return tar.The.Visible()
}

func (tar *Wait) Show() {
}
