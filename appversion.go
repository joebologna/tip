package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AppVersion int

const (
	AppVersion1 AppVersion = iota
	AppVersion2
)

func (app AppVersion) String() string {
	switch app {
	case AppVersion1:
		return "V1"
	case AppVersion2:
		return "V2"
	default:
		return "Unknown"
	}
}

func (v AppVersion) makeStuff() (stuff *fyne.Container, button *widget.Button) {
	switch v {
	case AppVersion1:
		return App1()
	default:
		panic("unsupported version")
	}
}
