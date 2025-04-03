package exp

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Stub to generate a grid with entries and a button which updates them, resolving issues with AdaptiveGrid
func App2() (*fyne.Container, *widget.Button) {
	strings := make([]binding.String, 0)

	// rows:=2
	cols := 4

	row := 0
	grid1 := container.NewGridWithColumns(4)
	for col := 0; col < cols; col++ {
		grid1.Add(LabelCell(row, col))
	}

	row++
	grid2 := container.NewGridWithColumns(4)
	for col := 0; col < cols; col++ {
		entryString := binding.NewString()
		strings = append(strings, entryString)
		grid2.Add(EntryCell(row, col, entryString))
	}

	stuff := container.NewVBox(grid1, grid2)
	button := widget.NewButton("update grid2", func() {
		for col, s := range strings {
			s.Set(fmt.Sprintf("Cell: %d,%d", 2, col+1))
		}
	})
	return stuff, button
}

func LabelCell(row, col int) fyne.CanvasObject {
	return widget.NewLabel(fmt.Sprintf("Cell: %d,%d", row+1, col+1))
}

func EntryCell(row, col int, text binding.String) fyne.CanvasObject {
	e := widget.NewEntryWithData(text)
	e.Validator = nil
	text.Set(fmt.Sprintf("Cell: %d,%d", row+1, col+1))
	return e
}
