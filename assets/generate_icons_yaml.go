package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Icon struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Path string `yaml:"path"`
	Seed bool   `yaml:"seed"`
}

type Icons []Icon

func main() {
	iconDir := "assets/icons"
	files, err := os.ReadDir(iconDir)
	if err != nil {
		log.Fatal(err)
	}

	var icons Icons
	for _, file := range files {
		if !file.IsDir() {
			icon := Icon{
				Path: filepath.Join(iconDir, file.Name()),
				Type: getType(file.Name()),
			}
			icon.Name, _ = strings.CutSuffix(file.Name(), icon.Type)
			icons = append(icons, icon)
		}
	}

	baseMaterials := []string{"oxygen", "hydrogen", "fire", "aluminum", "gold", "mercury", "cold", "earth", "electricity", "pressure", "sand", "iron", "sodium", "air", "chlorine", "phosphorus", "nitrogen", "calcium", "carbon"}
	for _, baseMaterial := range baseMaterials {
		icons.SetSeed(baseMaterial)
	}

	data, err := yaml.Marshal(&icons)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("icons.yaml", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("icons.yaml generated successfully")
}

func getType(name string) string {
	if strings.HasSuffix(name, ".svg") {
		return ".svg"
	}
	return ".png"
}

func (icons Icons) SetSeed(name string) {
	for i := 0; i < len(icons); i++ {
		icons[i].Seed = icons[i].Seed || icons[i].Name == name
	}
}
