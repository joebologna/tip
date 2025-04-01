package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Tip")

	stuff, button := AppVersion1.makeStuff()

	myWindow.SetContent(container.NewBorder(stuff, button, nil, nil))

	screenSize := GetScreenSize()
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}
