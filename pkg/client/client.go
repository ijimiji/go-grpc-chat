package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Hello")

	txts := []fyne.CanvasObject{
		widget.NewLabel("txt1"),
		widget.NewLabel("txt2"),
		widget.NewLabel("txt3"),
		widget.NewLabel("txt4"),
	}

	w.SetContent(container.NewVBox(txts...))
	w.ShowAndRun()
}
