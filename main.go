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
	for col := float32(0); col < 4; col++ {
		for row := float32(0); row < 5; row++ {
			entryString := binding.NewString()
			entryString.Set(fmt.Sprintf("Entry %d,%d", int(col+1), int(row+1)))
			strings = append(strings, entryString)
			entries = append(entries, MakeEntry(&entryString, size, fyne.NewPos((size.Width+2)*col, (size.Height+2)*row)))
		}
	}
	// pos := fyne.NewPos(0, (size.Height+2)*5)
	button := widget.NewButton("update rows", func() {
		fmt.Println("update rows")
		for _, v := range strings {
			v.Set("hi")
		}
	})
	button.Alignment = widget.ButtonAlignLeading
	button.Importance = widget.HighImportance
	button.Resize(size)

	grid := container.NewAdaptiveGrid(4, entries...)
	myWindow.SetContent(container.NewBorder(grid, button, nil, nil))

	screenSize := GetScreenSize()
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}

func MakeEntry(text *binding.String, size fyne.Size, loc fyne.Position) fyne.CanvasObject {
	entry := widget.NewEntryWithData(*text)
	entry.Validator = nil
	entry.Resize(size)
	entry.Move(loc)
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
