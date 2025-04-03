package custom

import (
	"fmt"
	"strings"
	"tip/utils"

	"fyne.io/fyne/v2/widget"
)

type TipPercentSelector struct {
	*widget.RadioGroup
	TipLabels       []string
	TipLabelDefault string
	TipFactors      []float32
	CurTipFactor    float32
	CalculateFunc   func(ts *TipPercentSelector)
	TotalEntry      *TotalEntry
}

func (ts *TipPercentSelector) Calculate() {
	ts.CalculateFunc(ts)
}

func NewTipPercentSelector(te *TotalEntry, calculate func(ts *TipPercentSelector)) *TipPercentSelector {
	tipLabels := make([]string, 0)
	for tip := 10; tip <= 25; tip += 5 {
		tipLabels = append(tipLabels, fmt.Sprintf("%d%%", tip))
	}
	ts := &TipPercentSelector{
		TipLabels:       tipLabels,
		TipLabelDefault: "20%",
		TipFactors:      make([]float32, 0),
		CalculateFunc:   calculate,
	}
	for _, v := range tipLabels {
		ts.TipFactors = append(ts.TipFactors, TipLabelToFactor(v))
	}
	ts.CurTipFactor = TipLabelToFactor(ts.TipLabelDefault)
	ts.RadioGroup = widget.NewRadioGroup(ts.TipLabels, func(selected string) {
		ts.CurTipFactor = TipLabelToFactor(selected)
		ts.Calculate()
	})
	ts.RadioGroup.Horizontal = true
	ts.RadioGroup.SetSelected(ts.TipLabelDefault)
	ts.TotalEntry = te
	return ts
}

func TipLabelToFactor(s string) float32 {
	return utils.ParseFloat32(strings.ReplaceAll(s, "%", "")) / 100.0
}
