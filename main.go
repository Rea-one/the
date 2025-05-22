package main

import (
	"the/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	the := view.NewWork()
	theApp := app.New()
	theWin := theApp.NewWindow("Canvas")
	theCav := theWin.Canvas()
	theCav.SetContent(the.The)

	theWin.Resize(fyne.NewSize(800, 500))
	theWin.ShowAndRun()
}
