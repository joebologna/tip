package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type customEntry struct {
	widget.Entry
	id  string
	sum *widget.Label
}

func newCustomEntryWithData(text binding.String, id string, sum *widget.Label) *customEntry {
	e := &customEntry{id: id, sum: sum}
	e.Bind(text)
	e.ExtendBaseWidget(e)
	return e
}

func (e *customEntry) FocusLost() {
	fmt.Println(e.id + " Focus Lost")
	e.sum.SetText(add(e, 1))
}

func add(e *customEntry, delta int) string {
	i, _ := strconv.ParseFloat(e.Text, 64)
	i = i + float64(delta)
	s := fmt.Sprintf("%.2f", i)
	return s
}

// App4 creates the UI with a custom entry and a button
func App4() (*fyne.Container, *widget.Button) {
	text := binding.NewString()
	text.Set("1")

	sum := widget.NewLabel("sum goes here")
	e := newCustomEntryWithData(text, "1", sum)
	e.Validator = nil
	button := widget.NewButton("push", func() {
		value, _ := text.Get()
		fmt.Println("Input value:", value)
		text.Set("1")
	})

	return container.NewBorder(e, sum, nil, nil), button
}
