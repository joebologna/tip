package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Tip")
	tileSize := float32(92)

	size := fyne.NewSize(80, 30)
	entries := make([]fyne.CanvasObject, 0)
	for col := float32(0); col < 4; col++ {
		for row := float32(0); row < 5; row++ {
			entries = append(entries, MakeEntry(fmt.Sprintf("Entry %d,%d", int(col+1), int(row+1)), size, fyne.NewPos((size.Width+2)*col, (size.Height+2)*row)))
		}
	}
	pos := fyne.NewPos(0, (size.Height+2)*5)
	button := widget.NewButton("update rows", func() {
		fmt.Println("update rows")
	})
	button.Alignment = widget.ButtonAlignLeading
	button.Importance = widget.HighImportance
	button.Resize(size)
	button.Move(pos)
	entries = append(entries, button)
	grid := container.NewWithoutLayout(entries...)
	myWindow.SetContent(grid)

	screenWidth := (tileSize + 5) * 4
	screenHeight := screenWidth * 2
	screenSize := fyne.NewSize(screenWidth, screenHeight) // pick a default size, the OS will resize as needed
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}

func MakeEntry(text string, size fyne.Size, loc fyne.Position) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetText(text)
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
