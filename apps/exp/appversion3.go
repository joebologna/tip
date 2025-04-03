package exp

import (
	"fmt"
	"strings"
	"tip/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type entryCell2 struct {
	widget.Entry
	id                   int
	sum, tip, sumWithTip utils.BS
	boundStrings         *[]utils.BS
	tips                 *widget.RadioGroup
	calculate            func(e *entryCell2, boundStrings []utils.BS, tips *widget.RadioGroup)
}

func newEntryCell2WithData(text binding.String, id int, sum, tip, sumWithTip utils.BS, boundStrings *[]utils.BS, tips *widget.RadioGroup, calculate func(e *entryCell2, boundStrings []utils.BS, tips *widget.RadioGroup)) *entryCell2 {
	e := &entryCell2{id: id, sum: sum, tip: tip, sumWithTip: sumWithTip, calculate: calculate, boundStrings: boundStrings, tips: tips}
	e.Bind(text)
	e.Validator = nil
	e.ExtendBaseWidget(e)
	return e
}

func (e *entryCell2) FocusLost() {
	fmt.Println(e.id, "Focus Lost")
	e.calculate(e, *e.boundStrings, e.tips)
	// this is needed to trigger the hide the cursor and remove highlight
	e.Entry.FocusLost()
}

// Stub to generate a grid with entries and a button which updates them, resolving issues with AdaptiveGrid
func App3() (*fyne.Container, fyne.CanvasObject) {
	boundStrings := make([]utils.BS, 0)
	entries := make([]*entryCell2, 0)

	selected := binding.NewString()
	totalBill, totalTip, totalBillWithTip := utils.NewBS(), utils.NewBS(), utils.NewBS()
	summary := makeSummary(totalBill, totalTip, totalBillWithTip)
	tips := makeTips(selected, totalBill, totalTip, totalBillWithTip, updateSummary)

	// rows:=2
	rows := 0
	cols := 4

	row := 0
	row++
	grid2 := container.NewGridWithColumns(4)
	for col := 0; col < cols; col++ {
		entryString := utils.NewBS()
		boundStrings = append(boundStrings, entryString)
		e := newEntryCell2WithData(entryString, col, totalBill, totalTip, totalBillWithTip, &boundStrings, tips, calc1)
		entries = append(entries, e)
		grid2.Add(e)
	}
	reset(rows, cols, boundStrings, totalBill, totalTip, totalBillWithTip, entries)

	stuff := container.NewVBox(tips, summary, grid2)
	button := widget.NewButton("Reset", func() {
		reset(rows, cols, boundStrings, totalBill, totalTip, totalBillWithTip, entries)
	})
	return stuff, container.NewGridWithColumns(2, widget.NewButton("Calculate Tip", func() {}), button)
}

func makeTips(selected, totalBill, totalTip, totalWithTip binding.String, updateSummary func(totalBill binding.String, totalBillValue string, totalTip binding.String, totalTipValue string, totalWithTip binding.String, totalWithTipValue string)) (tips *widget.RadioGroup) {
	tips = widget.NewRadioGroup([]string{"0.1", "0.15", "0.2", "0.25"}, func(changed string) {
		selected.Set(changed)
		updateSummary(totalBill, "0.00", totalTip, "0.00", totalWithTip, "0.00")
	})
	tips.SetSelected("0.2")
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

func reset(_, cols int, strings []utils.BS, totalBill, totalTip, totalWithTip binding.String, _ []*entryCell2) {
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

func calc1(e *entryCell2, boundStrings []utils.BS, tips *widget.RadioGroup) {
	sums := make([]string, 0)
	for i := 1; i < len(boundStrings); i++ {
		t := boundStrings[i].GetS()
		if t != "" {
			sums = append(sums, t)
		}
	}
	sum := fmt.Sprintf("(%s)", strings.Join(sums, "+"))
	e.sum.Set(fmt.Sprintf("%.2f", utils.EvalFloat(sum)))
	e.tip.Set(fmt.Sprintf("%.2f", utils.EvalFloat(fmt.Sprintf("%s * %s", sum, tips.Selected))))
	e.sumWithTip.Set(fmt.Sprintf("%.2f", utils.EvalFloat(fmt.Sprintf("%s * (1+%s)", sum, tips.Selected))))
}
