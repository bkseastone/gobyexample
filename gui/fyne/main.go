package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func toggle(label *widget.Label) func() {
	return func() {
		if label.Text == "Hi!" {
			label.SetText("Welcome :)")
		} else {
			label.SetText("Hi!")
		}
	}
}
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", toggle(hello)),
	))
	w.Resize(fyne.NewSize(200, 100))

	w.ShowAndRun()
}
