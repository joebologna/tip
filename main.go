package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Tip")

	stuff, button := makeStuff()

	myWindow.SetContent(container.NewBorder(stuff, button, nil, nil))

	screenSize := GetScreenSize()
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}

func makeStuff() (stuff *fyne.Container, button *widget.Button) {
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

	button = widget.NewButton("update rows", TestUpdate(strings, cols, rows))
	button.Alignment = widget.ButtonAlignLeading
	button.Importance = widget.HighImportance
	button.Resize(entrySize)

	stuff = container.NewVBox()
	grid := container.NewAdaptiveGrid(cols, entries...)
	stuff.Add(grid)

	stuff.Add(widget.NewLabel(O(fyne.CurrentDevice().Orientation()).String()))

	return stuff, button
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

func MakeEntry(text *binding.String, size fyne.Size) fyne.CanvasObject {
	entry := widget.NewEntryWithData(*text)
	entry.Validator = nil
	entry.Resize(size)
	entry.OnChanged = func(text string) { fmt.Println(text) }
	return fyne.CanvasObject(entry)
}
