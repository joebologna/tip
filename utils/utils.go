package utils

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/Knetic/govaluate"
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
	return fyne.NewSize(768, 1024)
}

type O fyne.DeviceOrientation

func (o O) String() string {
	if o == O(fyne.OrientationHorizontalLeft) {
		return "OrientationHorizontalLeft"
	}
	if o == O(fyne.OrientationHorizontalRight) {
		return "OrientationHorizontalRight"
	}
	if o == O(fyne.OrientationVertical) {
		return "OrientationVertical"
	}
	if o == O(fyne.OrientationVerticalUpsideDown) {
		return "OrientationVerticalUpsideDown"
	}
	return ""
}

func EvalFloat(e string) float32 {
	ee, err := govaluate.NewEvaluableExpression(e)
	if err != nil {
		return 0.0
	}
	result, err := ee.Evaluate(nil)
	if err != nil {
		return 0.0
	}
	// Handle different possible types of the result
	switch v := result.(type) {
	case float64:
		return float32(v)
	case int:
		return float32(v)
	default:
		fmt.Println("Unexpected result type:", v)
		return 0.0
	}
}

type BS struct{ binding.String }

func NewBS() BS { return BS{binding.NewString()} }

func (s BS) GetS() string { t, _ := s.Get(); return t }
