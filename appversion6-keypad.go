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

	billTitle, tipTitle, totalTitle := widget.NewLabel("Bill"), widget.NewLabel("Tip"), widget.NewLabel("Total")
	billTitle.Alignment, tipTitle.Alignment, totalTitle.Alignment = fyne.TextAlignCenter, fyne.TextAlignCenter, fyne.TextAlignCenter

	theBill, theTip, theTotal := widget.NewLabelWithData(bill), widget.NewLabelWithData(tip), widget.NewLabelWithData(total)
	theBill.Alignment, theTip.Alignment, theTotal.Alignment = fyne.TextAlignTrailing, fyne.TextAlignTrailing, fyne.TextAlignTrailing
	theBillLens, theTipLens, theTotalLens := canvas.NewRectangle(color.Transparent), canvas.NewRectangle(color.Transparent), canvas.NewRectangle(color.Transparent)
	theBillLens.StrokeWidth, theTipLens.StrokeWidth, theTotalLens.StrokeWidth = 1, 1, 1

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
				pending(true, theBillLens, theTipLens, theTotalLens)
			} else if key == "AC" {
				entryString.Set("")
				calc(bill, 0.0, tip, total)
				pending(true, theBillLens, theTipLens, theTotalLens)
			} else if key == "Calc" {
				sum := float32(0)
				for _, v := range strings.Split(entryString.GetS(), ",") {
					sum += utils.ParseFloat32(v)
				}
				calc(bill, sum, tip, total)
				pending(false, theBillLens, theTipLens, theTotalLens)
			} else {
				pending(true, theBillLens, theTipLens, theTotalLens)
				s := entryString.GetS()
				s += key
				entryString.Set(s)
			}
		})
		keys = append(keys, b)
	}
	stuff.Add(container.NewGridWithColumns(3, keys...))
	stuff.Add(container.NewGridWithColumns(3, billTitle, tipTitle, totalTitle))
	stuff.Add(container.NewGridWithColumns(3, container.NewStack(theBillLens, theBill), container.NewStack(theTipLens, theTip), container.NewStack(theTotalLens, theTotal)))
	calc(bill, 0.0, tip, total)
	pending(false, theBillLens, theTipLens, theTotalLens)
	return stuff, widget.NewButton("Bye", func() { os.Exit(0) })
}

func calc(bill utils.BS, sum float32, tip utils.BS, total utils.BS) {
	bill.Set(fmt.Sprintf("%.2f", sum))
	tip.Set(fmt.Sprintf("%.2f", sum*0.20))
	total.Set(fmt.Sprintf("%.2f", sum*1.20))
}

func pending(isPending bool, theBillLens, theTipLens, theTotalLens *canvas.Rectangle) {
	if isPending {
		theBillLens.StrokeColor, theTipLens.StrokeColor, theTotalLens.StrokeColor = RED, RED, RED
	} else {
		theBillLens.StrokeColor, theTipLens.StrokeColor, theTotalLens.StrokeColor = GREEN, GREEN, GREEN
	}
}
