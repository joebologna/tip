package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Tip")

	strings := make([]binding.String, 0)

	size := fyne.NewSize(80, 30)
	entries := make([]fyne.CanvasObject, 0)
	rows := 5
	cols := 4
	for i := 0; i < cols*rows; i++ {
		entryString := binding.NewString()
		if (i % cols) == 0 {
			entryString.Set(fmt.Sprintf("Entry %d", i/cols+1))
		}
		strings = append(strings, entryString)
		entries = append(entries, MakeEntry(&entryString, size))
	}

	button := widget.NewButton("update rows", TestUpdate(strings, cols, rows))
	button.Alignment = widget.ButtonAlignLeading
	button.Importance = widget.HighImportance
	button.Resize(size)

	grid := container.NewAdaptiveGrid(4, entries...)
	myWindow.SetContent(container.NewBorder(grid, button, nil, nil))

	screenSize := GetScreenSize()
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
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

func MakeROTile(text string) []fyne.CanvasObject {
	size := fyne.NewSize(100, 30)
	rect := canvas.NewRectangle(color.RGBA{0, 0, 0, 32})
	rect.StrokeWidth = 1
	rect.StrokeColor = color.Black
	rect.Resize(size)
	label := widget.NewLabel(text)
	return []fyne.CanvasObject{rect, label}
}

func GetDeviceType() (is_mobile, is_browser, is_desktop bool) {
	is_mobile, is_browser = fyne.CurrentDevice().IsMobile(), fyne.CurrentDevice().IsBrowser()
	is_desktop = !(is_mobile || is_browser)
	return
}

func GetScreenSize() fyne.Size {
	is_mobile, is_browser, _ := GetDeviceType()
	o := fyne.CurrentDevice().Orientation()
	if is_mobile || is_browser {
		if o == fyne.OrientationVertical || o == fyne.OrientationVerticalUpsideDown {
			return fyne.NewSize(768, 1024)
		}
		return fyne.NewSize(1024, 768)
	}
	if o == fyne.OrientationVertical || o == fyne.OrientationVerticalUpsideDown {
		return fyne.NewSize(768, 1024)
	}
	return fyne.NewSize(1024, 768)
}

type O fyne.DeviceOrientation

func (o O) String() string {
	if o == O(fyne.OrientationHorizontalLeft) {
		return "OrientationHorizontalLeft"
	}
	if o == O(fyne.OrientationHorizontalRight) {
		return "OrientationHorizontalRight"
	}
	if o == O(fyne.OrientationVertical) {
		return "OrientationVertical"
	}
	if o == O(fyne.OrientationVerticalUpsideDown) {
		return "OrientationVerticalUpsideDown"
	}
	return ""
}
