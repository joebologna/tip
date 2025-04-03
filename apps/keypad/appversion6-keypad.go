package keypad

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Generates a grid with entries and a button which updates them.
//
// Issues: The AdaptiveGrid does not behave as expected when rotating the phone. It does change the layout of the grid,
// but the device orientation is always vertical on mobile and horizontal left on desktop (regardless of screenSize)
func App6() (*fyne.Container, *widget.Button) {
	stuff := container.NewVBox(widget.NewLabel("keypad goes here"))

	return stuff, widget.NewButton("Bye", func() { os.Exit(0) })
}
