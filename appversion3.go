package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Stub to generate a grid with entries and a button which updates them, resolving issues with AdaptiveGrid
func App3() (*fyne.Container, *widget.Button) {
	strings := make([]binding.String, 0)

	selected := binding.NewString()
	tips := MakeTips(selected)
	summary := MakeSummary()

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

	stuff := container.NewVBox(tips, summary, grid1, grid2)
	button := widget.NewButton("update grid2", func() {
		for col, s := range strings {
			s.Set(fmt.Sprintf("Cell: %d,%d", 2, col+1))
		}
	})
	return stuff, button
}

func MakeTips(selected binding.String) (tips *widget.RadioGroup) {
	tips = widget.NewRadioGroup([]string{"10%", "15%", "20%", "25%"}, func(changed string) { selected.Set(changed) })
	tips.SetSelected("20%")
	tips.Horizontal = true
	return tips
}

func MakeSummary() (summary fyne.CanvasObject) {
	totalBillLabel := widget.NewLabel("Total Bill:")
	totalBill := binding.NewString()
	totalBill.Set("0.00")
	totalBillValue := widget.NewLabelWithData(totalBill)

	totalTipLabel := widget.NewLabel("Total Tip:")
	totalTip := binding.NewString()
	totalTip.Set("0.00")
	totalTipValue := widget.NewLabelWithData(totalTip)

	totalWithTipLabel := widget.NewLabel("Total with Tip:")
	totalWithTip := binding.NewString()
	totalWithTip.Set("0.00")
	totalWithTipValue := widget.NewLabelWithData(totalWithTip)

	splitEvenlyLabel := widget.NewCheck("Split Evenly:", func(onOff bool) { fmt.Println(onOff) })
	summary = container.NewVBox(
		container.NewHBox(totalBillLabel, totalBillValue),
		container.NewHBox(totalTipLabel, totalTipValue),
		container.NewHBox(totalWithTipLabel, totalWithTipValue),
		splitEvenlyLabel,
	)
	return summary
}
