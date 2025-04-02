package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type BString struct{ binding.String }

func NewBString() BString { return BString{binding.NewString()} }

type TipSelector struct {
	*widget.RadioGroup
	tipLabels       []string
	tipLabelDefault string
	tipFactors      []float32
	curTipFactor    float32
	calculate       func(ts *TipSelector)
}

func (ts *TipSelector) Calculate() {
	ts.calculate(ts)
}

func NewTipSelector(calculate func(ts *TipSelector)) *TipSelector {
	tipLabels := make([]string, 0)
	for tip := 10; tip <= 25; tip += 5 {
		tipLabels = append(tipLabels, fmt.Sprintf("%d%%", tip))
	}
	ts := &TipSelector{tipLabels: tipLabels, tipLabelDefault: "20%", tipFactors: make([]float32, 0), calculate: calculate}
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
	return ts
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
	// render tip radio group
	tipSelector := NewTipSelector(func(ts *TipSelector) { fmt.Println(ts) })
	return container.NewBorder(
		tipSelector.RadioGroup,
		nil,
		nil,
		nil,
	), container.NewHBox(layout.NewSpacer(), widget.NewButton("Bye", func() { os.Exit(0) }), layout.NewSpacer())
}
