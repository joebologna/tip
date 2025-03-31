package main

import (
	"fmt"
)

type IconRec struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Path string `yaml:"path"`
	Seed bool   `yaml:"seed"`
}

type Icons []IconRec

func (iconRecs Icons) GetIconRec(substanceName string) (iconRec IconRec, err error) {
	for _, rec := range iconRecs {
		if rec.Name == substanceName {
			return rec, nil
		}
	}
	return IconRec{}, fmt.Errorf("icon record not found for substance: %s", substanceName)
}
