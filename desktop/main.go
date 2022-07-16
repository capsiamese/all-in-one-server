package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

var list = NewLinkedList[int]()

func init() {
	for i := 0; i < 10; i++ {
		var t = i
		list.PushBack(&t)
	}
}

func main() {
	cnt := 0
	a := app.New()
	w := a.NewWindow("AIO")
	lst := widget.NewList(
		func() int {
			return cnt
		},
		func() fyne.CanvasObject {
			log.Println("create")
			return widget.NewLabel("item")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			log.Println("update", id)
			if l, ok := object.(*widget.Label); ok {
				l.Text = fmt.Sprintf("Item:%d", id)
			}
			object.Refresh()
		},
	)
	btn := widget.NewButton("Click", func() {
		cnt++
	})
	sc := container.NewVScroll(
		container.New(
			layout.NewVBoxLayout(),
			widget.NewButton("1", func() {
			}),
			widget.NewButton("2", func() {
			}),
		))
	sc.OnScrolled = func(position fyne.Position) {
		log.Println(position)
	}
	w.SetContent(container.New(layout.NewVBoxLayout(), lst, btn, sc))
	w.ShowAndRun()
}
