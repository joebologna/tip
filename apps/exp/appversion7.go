package exp

import (
	"image/color"
	"tip/apps/keypad"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// App7 is a stub
func App7() (*fyne.Container, *widget.Button) {
	// b := widget.NewButton("Bye", func() { os.Exit(0) })

	t := widget.NewLabel("V7")
	t.Alignment = fyne.TextAlignCenter

	stuff, b := keypad.App6()
	r1 := canvas.NewRectangle(color.RGBA{204, 128, 0, 255})
	r1.SetMinSize(stuff.MinSize().Add(fyne.NewSquareSize(10)))

	k := container.NewCenter(r1, stuff)

	return container.NewVBox(k, layout.NewSpacer(), b), b
}
