package main

import (
	"tip/apps/exp"
	"tip/apps/textonly"

	"fyne.io/fyne/v2"
)

type AppVersion int

const (
	AppVersion1 AppVersion = iota
	AppVersion2
	AppVersion3
	AppVersion4
	AppVersion5
	AppVersion6
)

func (app AppVersion) String() string {
	switch app {
	case AppVersion1:
		return "V1"
	case AppVersion2:
		return "V2"
	case AppVersion3:
		return "V3"
	case AppVersion4:
		return "V4"
	case AppVersion5:
		return "V5"
	case AppVersion6:
		return "V6"
	default:
		return "Unknown"
	}
}

func (v AppVersion) App() (stuff *fyne.Container, button fyne.CanvasObject) {
	switch v {
	case AppVersion1:
		return exp.App1()
	case AppVersion2:
		return exp.App2()
	case AppVersion3:
		return exp.App3()
	case AppVersion4:
		return exp.App4()
	case AppVersion5:
		return textonly.App5()
	case AppVersion6:
		return App6()
	default:
		panic("unsupported version")
	}
}
