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

	stuff, button := AppVersion1.makeStuff()

	myWindow.SetContent(container.NewBorder(stuff, button, nil, nil))

	screenSize := GetScreenSize()
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}

func (v AppVersion) makeStuff() (stuff *fyne.Container, button *widget.Button) {
	switch v {
	case AppVersion1:
		return App1()
	default:
		panic("unsupported version")
	}
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
