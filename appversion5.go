package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type TipPercentSelector struct {
	*widget.RadioGroup
	tipLabels       []string
	tipLabelDefault string
	tipFactors      []float32
	curTipFactor    float32
	calculate       func(ts *TipPercentSelector)
	totalEntry      *TotalEntry
}

func (ts *TipPercentSelector) Calculate() {
	ts.calculate(ts)
}

func NewTipSelector(te *TotalEntry, calculate func(ts *TipPercentSelector)) *TipPercentSelector {
	tipLabels := make([]string, 0)
	for tip := 10; tip <= 25; tip += 5 {
		tipLabels = append(tipLabels, fmt.Sprintf("%d%%", tip))
	}
	ts := &TipPercentSelector{tipLabels: tipLabels, tipLabelDefault: "20%", tipFactors: make([]float32, 0), calculate: calculate}
	for _, v := range tipLabels {
		ts.tipFactors = append(ts.tipFactors, TipLabelToFactor(v))
	}
	ts.curTipFactor = TipLabelToFactor(ts.tipLabelDefault)
	ts.RadioGroup = widget.NewRadioGroup(ts.tipLabels, func(selected string) {
		ts.curTipFactor = TipLabelToFactor(selected)
		ts.Calculate()
	})
	ts.RadioGroup.Horizontal = true
	ts.RadioGroup.SetSelected(ts.tipLabelDefault)
	ts.totalEntry = te
	return ts
}

type TotalEntry struct {
	widget.Entry
	summary *Summary
	ts      *TipPercentSelector
}

func NewTotalEntryWithData(text BS, summary *Summary) *TotalEntry {
	e := &TotalEntry{summary: summary}
	e.Bind(text)
	e.Validator = nil
	e.ExtendBaseWidget(e)
	return e
}

func (e *TotalEntry) FocusLost() {
	// this is needed to trigger the hide the cursor and remove highlight
	e.ts.Calculate()
	e.Entry.FocusLost()
}

type Summary struct {
	totalLabel, tipLabel, totalWithTipLabel *widget.Label
	summary                                 fyne.CanvasObject
	total, tip, totalWithTip                float32
}

func NewSummary() *Summary {
	s := &Summary{totalLabel: widget.NewLabel(""), tipLabel: widget.NewLabel(""), totalWithTipLabel: widget.NewLabel("")}
	s.summary = container.NewVBox(
		container.NewGridWithColumns(2, widget.NewLabel("Total:"), s.totalLabel),
		container.NewGridWithColumns(2, widget.NewLabel("Tip:"), s.tipLabel),
		container.NewGridWithColumns(2, widget.NewLabel("Total with Tip:"), s.totalWithTipLabel),
	)
	return s
}

func (s *Summary) Calculate(newTotal float32, ts *TipPercentSelector) {
	s.total = newTotal
	s.tip = newTotal * ts.curTipFactor
	s.totalWithTip = newTotal * (1 + ts.curTipFactor)
	s.totalLabel.SetText(fmt.Sprintf("%.2f", s.total))
	s.tipLabel.SetText(fmt.Sprintf("%.2f", s.tip))
	s.totalWithTipLabel.SetText(fmt.Sprintf("%.2f", s.totalWithTip))
}

func ParseFloat32(s string) float32 {
	num, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0
	}
	return float32(num)
}

func TipLabelToFactor(s string) float32 { return ParseFloat32(strings.ReplaceAll(s, "%", "")) / 100.0 }

func App5() (*fyne.Container, fyne.CanvasObject) {
	total := NewBS()
	summary := NewSummary()
	te := NewTotalEntryWithData(total, summary)
	tipSelector := NewTipSelector(te, func(ts *TipPercentSelector) {
		fmt.Println(ts)
		newTotal := ParseFloat32(total.get())
		summary.Calculate(newTotal, ts)
	})
	te.ts = tipSelector
	return container.NewBorder(
		container.NewVBox(
			tipSelector.RadioGroup,
			container.NewGridWithColumns(2, te, widget.NewLabel("Calculate")),
			summary.summary,
		),
		nil,
		nil,
		nil,
	), container.NewHBox(layout.NewSpacer(), widget.NewButton("Bye", func() { os.Exit(0) }), layout.NewSpacer())
}
