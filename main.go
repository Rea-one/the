package main

import (
	"the/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	theWork := view.NewWork()
	theApp := app.New()
	theWin := theApp.NewWindow("Canvas")
	theCav := theWin.Canvas()
	theCav.SetContent(theWork.The)

	theWin.Resize(fyne.NewSize(800, 500))
	theWin.ShowAndRun()
}
