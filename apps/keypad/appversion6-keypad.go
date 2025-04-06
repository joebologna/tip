package keypad

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
	var tipSelector, splitSelector *widget.RadioGroup

	entryString, bill, tip, total := utils.NewBS(), utils.NewBS(), utils.NewBS(), utils.NewBS()
	billEach, tipEach, totalEach := utils.NewBS(), utils.NewBS(), utils.NewBS()
	entry := widget.NewEntryWithData(entryString)
	entry.Validator = nil

	billTitle, tipTitle, totalTitle := NewThemedLabel("Bill"), NewThemedLabel("Tip"), NewThemedLabel("Total")
	billTitle.Alignment, tipTitle.Alignment, totalTitle.Alignment = fyne.TextAlignCenter, fyne.TextAlignCenter, fyne.TextAlignCenter

	theBill, theTip, theTotal := NewThemedLabelWithData(bill), NewThemedLabelWithData(tip), NewThemedLabelWithData(total)
	theBill.Alignment, theTip.Alignment, theTotal.Alignment = fyne.TextAlignTrailing, fyne.TextAlignTrailing, fyne.TextAlignTrailing

	theBillEach, theTipEach, theTotalEach := NewThemedLabelWithData(billEach), NewThemedLabelWithData(tipEach), NewThemedLabelWithData(totalEach)
	theBillEach.Alignment, theTipEach.Alignment, theTotalEach.Alignment = fyne.TextAlignTrailing, fyne.TextAlignTrailing, fyne.TextAlignTrailing

	stuff := container.NewVBox(entry)

	keys := make([]fyne.CanvasObject, 0)
	for _, key := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "0", "/", "AC", "Calc", "DEL"} {
		b := widget.NewButton(" "+key+" ", func() {
			if key == "DEL" {
				s := entryString.GetS()
				n := len(s)
				if n > 0 {
					s = s[:n-1]
					entryString.Set(s)
				}
				pending(true, theBill, theTip, theTotal, theBillEach, theTipEach, theTotalEach)
			} else if key == "AC" {
				entryString.Set("")
				doCalc(entryString, bill, tipSelector, splitSelector, tip, total, billEach, tipEach, totalEach)
				pending(false, theBill, theTip, theTotal, theBillEach, theTipEach, theTotalEach)
			} else if key == "Calc" {
				doCalc(entryString, bill, tipSelector, splitSelector, tip, total, billEach, tipEach, totalEach)
				pending(false, theBill, theTip, theTotal, theBillEach, theTipEach, theTotalEach)
			} else {
				pending(true, theBill, theTip, theTotal, theBillEach, theTipEach, theTotalEach)
				s := entryString.GetS()
				s += key
				entryString.Set(s)
			}
		})
		keys = append(keys, b)
	}

	tipSelector = widget.NewRadioGroup([]string{"10%", "15%", "20%", "25%"}, func(cur_selection string) {
		doCalc(entryString, bill, tipSelector, splitSelector, tip, total, billEach, tipEach, totalEach)
		pending(false, theBill, theTip, theTotal, theBillEach, theTipEach, theTotalEach)
	})

	splitSelector = widget.NewRadioGroup([]string{"1", "2", "3", "4", "5", "6"}, func(cur_selection string) {
		doCalc(entryString, bill, tipSelector, splitSelector, tip, total, billEach, tipEach, totalEach)
		pending(false, theBill, theTip, theTotal, theBillEach, theTipEach, theTotalEach)
	})
	splitSelector.Horizontal = true
	splitSelector.SetSelected("1")
	split := utils.ParseFloat32("1")
	tipSelector.Horizontal = true
	tipSelector.SetSelected("20%")

	stuff.Add(container.NewCenter(tipSelector))
	stuff.Add(container.NewCenter(splitSelector))
	stuff.Add(container.NewGridWithColumns(3, keys...))
	stuff.Add(container.NewGridWithColumns(3, billTitle.Stack(), tipTitle.Stack(), totalTitle.Stack()))
	stuff.Add(container.NewGridWithColumns(3, theBill.Stack(), theTip.Stack(), theTotal.Stack()))
	stuff.Add(container.NewGridWithColumns(3, theBillEach.Stack(), theTipEach.Stack(), theTotalEach.Stack()))

	calc(bill, TipFactor(tipSelector.Selected), 0.0, split, tip, total, billEach, tipEach, totalEach)
	pending(false, theBill, theTip, theTotal, theBillEach, theTipEach, theTotalEach)

	bg := canvas.NewRectangle(color.Transparent)
	bg.StrokeWidth = 2
	bg.StrokeColor = color.RGBA{128, 128, 128, 128}

	return container.NewStack(stuff, bg), widget.NewButton("Bye", func() { os.Exit(0) })
}

func doCalc(entryString utils.BS, bill utils.BS, tipSelector, splitSelector *widget.RadioGroup, tip, total, billEach, tipEach, totalEach utils.BS) {
	sum := float32(0)
	parts := strings.Split(entryString.GetS(), "/")
	if len(parts) == 2 {
		sum = utils.ParseFloat32(parts[0])
		splitSelector.SetSelected(parts[1])
		entryString.Set(parts[0])
	} else {
		sum = utils.ParseFloat32(entryString.GetS())
	}

	split := utils.ParseFloat32(splitSelector.Selected)
	calc(bill, TipFactor(tipSelector.Selected), sum, split, tip, total, billEach, tipEach, totalEach)
}

func calc(bill utils.BS, percent, sum, split float32, tip, total, billEach, tipEach, totalEach utils.BS) {
	bill.Set(fmt.Sprintf("%.2f", sum))
	tip.Set(fmt.Sprintf("%.2f", sum*percent))
	total.Set(fmt.Sprintf("%.2f", sum*(1+percent)))
	billEach.Set(fmt.Sprintf("%.2f", sum/split))
	tipEach.Set(fmt.Sprintf("%.2f", sum*percent/split))
	totalEach.Set(fmt.Sprintf("%.2f", sum*(1+percent)/split))
}

func pending(isPending bool, theBillLens, theTipLens, theTotalLens, theBillEachLens, theTipEachLens, theTotalEachLens *ThemedLabel) {
	if isPending {
		theBillLens.overlay.StrokeColor, theTipLens.overlay.StrokeColor, theTotalLens.overlay.StrokeColor = RED, RED, RED
		theBillEachLens.overlay.StrokeColor, theTipEachLens.overlay.StrokeColor, theTotalEachLens.overlay.StrokeColor = RED, RED, RED
	} else {
		theBillLens.overlay.StrokeColor, theTipLens.overlay.StrokeColor, theTotalLens.overlay.StrokeColor = GREEN, GREEN, GREEN
		theBillEachLens.overlay.StrokeColor, theTipEachLens.overlay.StrokeColor, theTotalEachLens.overlay.StrokeColor = GREEN, GREEN, GREEN
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
