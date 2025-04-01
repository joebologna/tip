package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func App2() (*fyne.Container, *widget.Button) {
	stuff := container.NewVBox(widget.NewLabel("This is a stub for App2"))
	button := widget.NewButton("stub button", func() {})
	return stuff, button
}
