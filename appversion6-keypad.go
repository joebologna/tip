package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"
	"tip/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var RED, GREEN = color.RGBA{255, 0, 0, 128}, color.RGBA{0, 255, 0, 128}

func App6() (*fyne.Container, *widget.Button) {
	var tipSelector *widget.RadioGroup

	entryString, bill, tip, total := utils.NewBS(), utils.NewBS(), utils.NewBS(), utils.NewBS()
	entry := widget.NewEntryWithData(entryString)
	entry.Validator = nil

	billTitle, tipTitle, totalTitle := NewThemedLabel("Bill"), NewThemedLabel("Tip"), NewThemedLabel("Total")
	billTitle.Alignment, tipTitle.Alignment, totalTitle.Alignment = fyne.TextAlignCenter, fyne.TextAlignCenter, fyne.TextAlignCenter

	theBill, theTip, theTotal := NewThemedLabelWithData(bill), NewThemedLabelWithData(tip), NewThemedLabelWithData(total)
	theBill.Alignment, theTip.Alignment, theTotal.Alignment = fyne.TextAlignTrailing, fyne.TextAlignTrailing, fyne.TextAlignTrailing

	stuff := container.NewVBox(entry)

	keys := make([]fyne.CanvasObject, 0)
	for _, key := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "0", ",", "AC", "Calc", "DEL"} {
		b := widget.NewButton(" "+key+" ", func() {
			if key == "DEL" {
				s := entryString.GetS()
				n := len(s)
				if n > 0 {
					s = s[:n-1]
					entryString.Set(s)
				}
				pending(true, theBill, theTip, theTotal)
			} else if key == "AC" {
				entryString.Set("")
				doCalc(entryString, bill, tipSelector, tip, total)
				pending(false, theBill, theTip, theTotal)
			} else if key == "Calc" {
				doCalc(entryString, bill, tipSelector, tip, total)
				pending(false, theBill, theTip, theTotal)
			} else {
				pending(true, theBill, theTip, theTotal)
				s := entryString.GetS()
				s += key
				entryString.Set(s)
			}
		})
		keys = append(keys, b)
	}

	tipSelector = widget.NewRadioGroup([]string{"10%", "15%", "20%", "25%"}, func(cur_selection string) {
		doCalc(entryString, bill, tipSelector, tip, total)
		pending(false, theBill, theTip, theTotal)
	})
	tipSelector.Horizontal = true
	tipSelector.SetSelected("20%")

	stuff.Add(container.NewCenter(tipSelector))
	stuff.Add(container.NewGridWithColumns(3, keys...))
	stuff.Add(container.NewGridWithColumns(3, billTitle.Stack(), tipTitle.Stack(), totalTitle.Stack()))
	stuff.Add(container.NewGridWithColumns(3, theBill.Stack(), theTip.Stack(), theTotal.Stack()))

	calc(bill, TipFactor(tipSelector.Selected), 0.0, tip, total)
	pending(false, theBill, theTip, theTotal)

	bg := canvas.NewRectangle(color.Transparent)
	bg.StrokeWidth = 2
	bg.StrokeColor = color.RGBA{128, 128, 128, 128}

	return container.NewStack(stuff, bg), widget.NewButton("Bye", func() { os.Exit(0) })
}

func doCalc(entryString utils.BS, bill utils.BS, tipSelector *widget.RadioGroup, tip utils.BS, total utils.BS) {
	sum := float32(0)
	for _, v := range strings.Split(entryString.GetS(), ",") {
		sum += utils.ParseFloat32(v)
	}
	calc(bill, TipFactor(tipSelector.Selected), sum, tip, total)
}

func calc(bill utils.BS, percent float32, sum float32, tip utils.BS, total utils.BS) {
	bill.Set(fmt.Sprintf("%.2f", sum))
	tip.Set(fmt.Sprintf("%.2f", sum*percent))
	total.Set(fmt.Sprintf("%.2f", sum*(1+percent)))
}

func pending(isPending bool, theBillLens, theTipLens, theTotalLens *ThemedLabel) {
	if isPending {
		theBillLens.overlay.StrokeColor, theTipLens.overlay.StrokeColor, theTotalLens.overlay.StrokeColor = RED, RED, RED
	} else {
		theBillLens.overlay.StrokeColor, theTipLens.overlay.StrokeColor, theTotalLens.overlay.StrokeColor = GREEN, GREEN, GREEN
	}
}

type ThemedLabel struct {
	*widget.Label
	overlay *canvas.Rectangle
}

func NewThemedLabel(text string) *ThemedLabel {
	l := &ThemedLabel{Label: widget.NewLabel(text), overlay: canvas.NewRectangle(GREEN)}
	l.overlay.StrokeWidth = 1
	return l
}

func NewThemedLabelWithData(text utils.BS) *ThemedLabel {
	l := &ThemedLabel{Label: widget.NewLabelWithData(text), overlay: canvas.NewRectangle(color.Transparent)}
	l.overlay.StrokeWidth = 1
	return l
}

func (t *ThemedLabel) Stack() fyne.CanvasObject {
	return container.NewStack(t.overlay, t.Label)
}

func TipFactor(s string) float32 {
	return utils.ParseFloat32(strings.ReplaceAll(s, "%", "")) / 100.0
}
