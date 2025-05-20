package main

import (
	"the/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	theWait := view.NewWait()
	theApp := app.New()
	theWin := theApp.NewWindow("Canvas")
	theCav := theWin.Canvas()
	theCav.SetContent(theWait.The)

	theWin.Resize(fyne.NewSize(800, 500))
	theWin.ShowAndRun()
}
