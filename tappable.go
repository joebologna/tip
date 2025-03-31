package main

import (
	"image/color"

	"fyne.io/fyne/v2/widget"
)

type TappableIcon struct {
	*widget.Icon
	Label         string
	OriginalColor color.Color
	Tap           func(label string)
	Discovered    bool
}

type TappableIconMap map[string]*TappableIcon

func (d TappableIconMap) Exists(label string) bool {
	_, ok := d[label]
	return ok
}
