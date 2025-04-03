package main

import (
	"os"
	"tip/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Generates a grid with entries and a button which updates them.
//
// Issues: The AdaptiveGrid does not behave as expected when rotating the phone. It does change the layout of the grid,
// but the device orientation is always vertical on mobile and horizontal left on desktop (regardless of screenSize)
func App6() (*fyne.Container, *widget.Button) {
	entryString := utils.NewBS()
	entry := widget.NewEntryWithData(entryString)
	// entry.PlaceHolder = ""
	entry.Validator = nil
	stuff := container.NewVBox(entry)
	keys := make([]fyne.CanvasObject, 0)
	for _, key := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", ",", "0", ".", "AC", " ", "DEL"} {
		b := widget.NewButton(" "+key+" ", func() {
			if key == "DEL" {
				s := entryString.GetS()
				n := len(s)
				if n > 0 {
					s = s[:n-1]
					entryString.Set(s)
				}
			} else if key == "AC" {
				entryString.Set("")
			} else {
				s := entryString.GetS()
				s += key
				entryString.Set(s)
			}
		})
		keys = append(keys, b)
	}
	stuff.Add(container.NewGridWithColumns(3, keys...))
	return stuff, widget.NewButton("Bye", func() { os.Exit(0) })
}
