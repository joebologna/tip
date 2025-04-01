package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
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
		grid2.Add(entryCell2(entryString, func(value string) { fmt.Println(col, value) }))
	}
	reset(rows, cols, strings, totalBill, totalTip, totalBillWithTip)

	stuff := container.NewVBox(tips, summary, grid2)
	button := widget.NewButton("Reset", func() {
		reset(rows, cols, strings, totalBill, totalTip, totalBillWithTip)
	})
	return stuff, button
}

func makeTips(selected, totalBill, totalTip, totalWithTip binding.String, updateSummary func(totalBill binding.String, totalBillValue string, totalTip binding.String, totalTipValue string, totalWithTip binding.String, totalWithTipValue string)) (tips *widget.RadioGroup) {
	tips = widget.NewRadioGroup([]string{"10%", "15%", "20%", "25%"}, func(changed string) {
		selected.Set(changed)
		updateSummary(totalBill, "0.00", totalTip, "0.00", totalWithTip, "0.00")
	})
	tips.SetSelected("20%")
	tips.Horizontal = true
	return tips
}

func makeSummary(totalBill, totalTip, totalWithTip binding.String) (summary fyne.CanvasObject) {
	align := fyne.TextAlignLeading
	createLabelWithValue := func(labelText string, value binding.String) (*widget.Label, *widget.Label) {
		label := widget.NewLabel(labelText)
		label.Alignment = align
		valueLabel := widget.NewLabelWithData(value)
		valueLabel.Alignment = align
		return label, valueLabel
	}

	totalBillLabel, totalBillValue := createLabelWithValue("Total Bill:", totalBill)
	totalTipLabel, totalTipValue := createLabelWithValue("Total Tip:", totalTip)
	totalWithTipLabel, totalWithTipValue := createLabelWithValue("Total with Tip:", totalWithTip)

	splitEvenlyLabel := widget.NewCheck("Split Evenly", func(onOff bool) { fmt.Println(onOff) })

	summary = container.NewHBox(
		layout.NewSpacer(),
		container.NewVBox(
			container.NewGridWithColumns(2,
				totalBillLabel,
				totalBillValue,
				totalTipLabel,
				totalTipValue,
				totalWithTipLabel,
				totalWithTipValue,
			),
			splitEvenlyLabel,
		),
		layout.NewSpacer(),
	)
	return summary
}

func entryCell2(text binding.String, onChanged func(string)) fyne.CanvasObject {
	e := widget.NewEntryWithData(text)
	e.Validator = nil
	e.OnChanged = onChanged
	return e
}

func reset(rows, cols int, strings []binding.String, totalBill, totalTip, totalWithTip binding.String) {
	// for row := 0; row < rows; row++ {
	for col := 0; col < cols; col++ {
		if col == 0 {
			strings[col].Set(fmt.Sprintf("Payee %d", col+1))
		} else {
			strings[col].Set("")
		}
	}
	// }
	updateSummary(totalBill, "0.00", totalTip, "0.00", totalWithTip, "0.00")
}

func updateSummary(totalBill binding.String, totalBillValue string, totalTip binding.String, totalTipValue string, totalWithTip binding.String, totalWithTipValue string) {
	totalBill.Set(totalBillValue)
	totalTip.Set(totalTipValue)
	totalWithTip.Set(totalWithTipValue)
}
