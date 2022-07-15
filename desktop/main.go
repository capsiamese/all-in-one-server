package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("AIO")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
