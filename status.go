package main

import (
	"tip/pairs"
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

type StatusMsg int

const (
	StatusHelp = iota
	StatusSelected
	StatusNoReaction
	StatusNewSubstance
	StatusExistingSubstance
)

type StatusLabel struct {
	*widget.Label
}

func NewStatusLabel(substanceMap pairs.SubstanceMap) *StatusLabel {
	l := &StatusLabel{Label: widget.NewLabel("")}
	l.UpdateStatus(StatusHelp, substanceMap.Found())
	return l
}

func (l *StatusLabel) UpdateStatus(s StatusMsg, args ...string) {
	switch s {
	case StatusSelected:
		l.SetText(fmt.Sprintf("%s has been selected. %s", capitalize(args[0]), args[1]))
	case StatusNoReaction:
		l.SetText(fmt.Sprintf("%s + %s = (no reaction). %s", capitalize(args[0]), capitalize(args[1]), args[2]))
	case StatusNewSubstance:
		l.SetText(fmt.Sprintf("%s + %s = %s. %s", capitalize(args[0]), capitalize(args[1]), capitalize(args[2]), args[3]))
	case StatusExistingSubstance:
		l.SetText(fmt.Sprintf("%s + %s = %s (already created). %s", capitalize(args[0]), capitalize(args[1]), capitalize(args[2]), args[3]))
	case StatusHelp:
		fallthrough
	default:
		l.SetText(fmt.Sprintf("Select two tiles to create a new substance. %s", args[0]))
	}
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
