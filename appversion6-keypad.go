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
				calc(bill, 0.0, tip, total)
				pending(true, theBill, theTip, theTotal)
			} else if key == "Calc" {
				sum := float32(0)
				for _, v := range strings.Split(entryString.GetS(), ",") {
					sum += utils.ParseFloat32(v)
				}
				calc(bill, sum, tip, total)
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
	stuff.Add(container.NewGridWithColumns(3, keys...))
	stuff.Add(container.NewGridWithColumns(3, billTitle.Stack(), tipTitle.Stack(), totalTitle.Stack()))
	stuff.Add(container.NewGridWithColumns(3, theBill.Stack(), theTip.Stack(), theTotal.Stack()))
	calc(bill, 0.0, tip, total)
	pending(false, theBill, theTip, theTotal)
	bg := canvas.NewRectangle(color.Transparent)
	bg.StrokeWidth = 2
	bg.StrokeColor = color.RGBA{128, 128, 128, 128}
	return container.NewStack(stuff, bg), widget.NewButton("Bye", func() { os.Exit(0) })
}

func calc(bill utils.BS, sum float32, tip utils.BS, total utils.BS) {
	bill.Set(fmt.Sprintf("%.2f", sum))
	tip.Set(fmt.Sprintf("%.2f", sum*0.20))
	total.Set(fmt.Sprintf("%.2f", sum*1.20))
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
