package main

import (
	"tip/apps/keypadonly"
	"tip/utils"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Tip")

	stuff := keypadonly.App8()
	myWindow.SetContent(stuff)
	screenSize := utils.GetScreenSize()
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}
