package main

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
