package custom

import (
	"tip/utils"

	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type TotalEntry struct {
	widget.Entry
	Summary *Summary
	TipSel  *TipPercentSelector
}

func (te *TotalEntry) Keyboard() mobile.KeyboardType {
	return mobile.SingleLineKeyboard
}

func (te *TotalEntry) FocusLost() {
	// This is needed to hide the cursor and remove highlight
}

func NewTotalEntryWithData(text utils.BS, summary *Summary) *TotalEntry {
	e := &TotalEntry{Summary: summary}
	e.Bind(text)
	e.PlaceHolder = "Amounts w/commas"
	e.Validator = nil
	e.ExtendBaseWidget(e)
	return e
}
