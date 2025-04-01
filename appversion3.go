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
	totalBill, totalTip, totalBillWithTip := binding.NewString(), binding.NewString(), binding.NewString()
	summary := makeSummary(totalBill, totalTip, totalBillWithTip)
	tips := makeTips(selected, totalBill, totalTip, totalBillWithTip, updateSummary)

	// rows:=2
	rows := 0
	cols := 4

	row := 0
	row++
	grid2 := container.NewGridWithColumns(4)
	for col := 0; col < cols; col++ {
		entryString := binding.NewString()
		strings = append(strings, entryString)
		grid2.Add(entryCell2(entryString))
	}
	updateCells(rows, cols, strings)

	stuff := container.NewVBox(tips, summary, grid2)
	button := widget.NewButton("update grid2", func() {
		updateCells(rows, cols, strings)
	})
	return stuff, button
}

func makeTips(selected, totalBill, totalTip, totalWithTip binding.String, updateSummary func(totalBill, totalTip, totalWithTip binding.String)) (tips *widget.RadioGroup) {
	tips = widget.NewRadioGroup([]string{"10%", "15%", "20%", "25%"}, func(changed string) {
		selected.Set(changed)
		updateSummary(totalBill, totalTip, totalWithTip)
	})
	tips.SetSelected("20%")
	tips.Horizontal = true
	return tips
}

func makeSummary(totalBill, totalTip, totalWithTip binding.String) (summary fyne.CanvasObject) {
	totalBillLabel := widget.NewLabel("Total Bill:")
	totalBillValue := widget.NewLabelWithData(totalBill)

	totalTipLabel := widget.NewLabel("Total Tip:")
	totalTipValue := widget.NewLabelWithData(totalTip)

	totalWithTipLabel := widget.NewLabel("Total with Tip:")
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

func entryCell2(text binding.String) fyne.CanvasObject {
	e := widget.NewEntryWithData(text)
	e.Validator = nil
	return e
}

func updateCells(rows, cols int, strings []binding.String) {
	// for row := 0; row < rows; row++ {
	for col := 0; col < cols; col++ {
		if col == 0 {
			strings[col].Set(fmt.Sprintf("Payee %d", col+1))
		} else {
			strings[col].Set("")
		}
	}
	// }
}

func updateSummary(totalBill, totalTip, totalWithTip binding.String) {
	totalBill.Set("Update Bill")
	totalTip.Set("Update Tip")
	totalWithTip.Set("Update Bill with Tip")
}
