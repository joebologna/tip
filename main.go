package main

import (
	"tip/pairs"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v2"
)

var substanceMap pairs.SubstanceMap

//go:generate go run assets/generate_icons_yaml.go

func main() {
	// Read icons.yaml from embedded resources
	data, err := embeddedFiles.ReadFile("icons.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var icons Icons
	err = yaml.Unmarshal(data, &icons)
	if err != nil {
		log.Fatal(err)
	}

	substanceMap = make(pairs.SubstanceMap)
	err = substanceMap.LoadYAML("substances.yaml")
	if err != nil {
		log.Fatal(err)
	}

	myApp := app.New()
	myApp.Settings().SetTheme(&CustomTheme{theme.DefaultTheme()})
	myWindow := myApp.NewWindow("Alchemy")

	status := NewStatusLabel(substanceMap)
	status.UpdateStatus(StatusHelp, substanceMap.Found())

	var scroller *container.Scroll
	var grid *fyne.Container

	tiles := make([]fyne.CanvasObject, 0)

	tileSize := float32(92)
	labelSize := fyne.NewSize(100, 20)
	tappedTiles := []string{}
	var tap func(string)
	tap = func(label string) {
		if len(tappedTiles) == 0 {
			tappedTiles = append(tappedTiles, label)
			status.UpdateStatus(StatusSelected, tappedTiles[0], substanceMap.Found())
		} else if len(tappedTiles) == 1 {
			tappedTiles = append(tappedTiles, label)
			substanceName, exists := substanceMap.GetSubstance(tappedTiles)
			if exists {
				if substanceMap[substanceName].Discovered {
					status.UpdateStatus(StatusExistingSubstance, tappedTiles[0], tappedTiles[1], string(substanceName), substanceMap.Found())
				} else {
					substance := substanceMap[substanceName]
					substance.Discovered = true
					substanceMap[substanceName] = substance
					status.UpdateStatus(StatusNewSubstance, tappedTiles[0], tappedTiles[1], string(substanceName), substanceMap.Found())
					iconRec, _ := icons.GetIconRec(string(substanceName))
					tile := NewTile(iconRec, tileSize, labelSize, tap)
					tiles = append(tiles, tile)
					grid.Add(tile)
				}
			} else {
				status.UpdateStatus(StatusNoReaction, tappedTiles[0], tappedTiles[1], substanceMap.Found())
			}
			tappedTiles = []string{}
		} else {
			panic("should not get here")
		}
	}

	for _, iconRec := range icons {
		if iconRec.Seed {
			tiles = append(tiles, NewTile(iconRec, tileSize, labelSize, tap))
		}
	}

	grid = container.NewGridWrap(fyne.NewSquareSize(tileSize), tiles...)
	scroller = container.NewVScroll(grid)

	var reset *widget.Button
	var resetRect *canvas.Rectangle
	var resetButton *fyne.Container

	reset = widget.NewButton("Reset", func() {
		substanceMap = make(pairs.SubstanceMap)
		err = substanceMap.LoadYAML("substances.yaml")
		if err != nil {
			log.Fatal(err)
		}

		tiles = make([]fyne.CanvasObject, 0)

		for _, iconRec := range icons {
			if iconRec.Seed {
				tiles = append(tiles, NewTile(iconRec, tileSize, labelSize, tap))
			}
		}

		grid = container.NewGridWrap(fyne.NewSquareSize(tileSize), tiles...)
		scroller = container.NewVScroll(grid)

		status.UpdateStatus(StatusHelp, substanceMap.Found())
		myWindow.SetContent(container.NewBorder(resetButton, container.NewCenter(status), nil, nil, scroller))
	})

	resetRect = canvas.NewRectangle(color.RGBA{64, 64, 64, 128})
	resetRect.SetMinSize(fyne.NewSize(100, 20))
	resetButton = container.NewStack(reset, resetRect)
	myWindow.SetContent(container.NewBorder(resetButton, container.NewCenter(status), nil, nil, scroller))

	screenWidth := (tileSize + 5) * 4
	screenHeight := screenWidth * 2
	screenSize := fyne.NewSize(screenWidth, screenHeight) // pick a default size, the OS will resize as needed
	myWindow.Resize(screenSize)
	myWindow.ShowAndRun()
}
