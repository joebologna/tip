package custom

import (
	"fmt"
	"tip/utils"

	"fyne.io/fyne/v2/widget"
)

type Summary struct {
	Summary *widget.Label
}

func NewSummary() *Summary {
	return &Summary{
		Summary: widget.NewLabel(""),
	}
}

func (s *Summary) Calculate(total utils.BS, ts *TipPercentSelector) {
	totalValue := utils.ParseFloat32(total.GetS())
	tipValue := totalValue * ts.CurTipFactor
	totalWithTip := totalValue + tipValue

	s.Summary.SetText(fmt.Sprintf("Total: %.2f, Tip: %.2f, Total with Tip: %.2f", totalValue, tipValue, totalWithTip))
}
