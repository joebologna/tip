package main

import (
	"tip/apps"
	"tip/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Tip")

	v := apps.AppVersion6
	vl := widget.NewLabel(v.String())
	vl.Alignment = fyne.TextAlignCenter
	stuff, button := v.App()

	myWindow.SetContent(container.NewBorder(container.NewCenter(stuff), container.NewVBox(button, vl), nil, nil))

	screenSize := utils.GetScreenSize()
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}
