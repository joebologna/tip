package utils

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

func GetDeviceType() (is_mobile, is_browser, is_desktop bool) {
	is_mobile, is_browser = fyne.CurrentDevice().IsMobile(), fyne.CurrentDevice().IsBrowser()
	is_desktop = !(is_mobile || is_browser)
	return
}

func GetScreenSize() fyne.Size {
	is_mobile, is_browser, _ := GetDeviceType()
	o := fyne.CurrentDevice().Orientation()
	if is_mobile || is_browser {
		if o == fyne.OrientationVertical || o == fyne.OrientationVerticalUpsideDown {
			return fyne.NewSize(768, 1024)
		}
		return fyne.NewSize(1024, 768)
	}
	if o == fyne.OrientationVertical || o == fyne.OrientationVerticalUpsideDown {
		return fyne.NewSize(768, 1024)
	}
	return fyne.NewSize(768/2, 1024)
}

type BS struct{ binding.String }

func NewBS() BS { return BS{binding.NewString()} }

func (s BS) GetS() string { t, _ := s.Get(); return t }

func ParseFloat32(s string) float32 {
	num, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0
	}
	return float32(num)
}
