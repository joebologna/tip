package exp

import (
	"fmt"
	"tip/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Generates a grid with entries and a button which updates them.
//
// Issues: The AdaptiveGrid does not behave as expected when rotating the phone. It does change the layout of the grid,
// but the device orientation is always vertical on mobile and horizontal left on desktop (regardless of screenSize)
func App1() (*fyne.Container, *widget.Button) {
	strings := make([]binding.String, 0)

	entrySize := fyne.NewSize(80, 30)
	entries := make([]fyne.CanvasObject, 0)
	rows := 5
	cols := 4
	for i := 0; i < cols*rows; i++ {
		entryString := binding.NewString()
		if (i % cols) == 0 {
			entryString.Set(fmt.Sprintf("Entry %d", i/cols+1))
		}
		strings = append(strings, entryString)
		entries = append(entries, MakeEntry(&entryString, entrySize))
	}

	button := widget.NewButton("update rows", TestUpdate(strings, cols, rows))
	button.Alignment = widget.ButtonAlignLeading
	button.Importance = widget.HighImportance
	button.Resize(entrySize)

	stuff := container.NewVBox()
	grid := container.NewAdaptiveGrid(cols, entries...)
	stuff.Add(grid)

	stuff.Add(widget.NewLabel(utils.O(fyne.CurrentDevice().Orientation()).String()))

	return stuff, button
}

func MakeEntry(text *binding.String, size fyne.Size) fyne.CanvasObject {
	entry := widget.NewEntryWithData(*text)
	entry.Validator = nil
	entry.Resize(size)
	entry.OnChanged = func(text string) { fmt.Println(text) }
	return fyne.CanvasObject(entry)
}

func TestUpdate(strings []binding.String, cols, rows int) func() {
	return func() {
		fmt.Println("update rows")
		for i, v := range strings {
			if (i % cols) == 0 {
				v.Set("Col 1")
			} else {
				v.Set(".")
			}
		}
	}
}
