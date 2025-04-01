package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AppVersion int

const (
	AppVersion1 AppVersion = iota
	AppVersion2
	AppVersion3
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

func (v AppVersion) app() (stuff *fyne.Container, button *widget.Button) {
	switch v {
	case AppVersion1:
		return App1()
	case AppVersion2:
		return App2()
	case AppVersion3:
		return App3()
	default:
		panic("unsupported version")
	}
}
