package pairs

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

type DisplayedState struct {
	Tapped, Displayed bool
}

type TappedLabels map[string]DisplayedState
type Substance string
type SubstanceMakeup struct {
	Labels     []string
	Discovered bool
}
type SubstanceMap map[Substance]SubstanceMakeup

func (substanceMap SubstanceMap) Found() (ratio_found string) {
	counts := []int{0, len(substanceMap)}
	for _, makeup := range substanceMap {
		if makeup.Discovered {
			counts[0]++
		}
	}
	return fmt.Sprintf("%d/%d", counts[0], counts[1])
}

func NewTappedLabels(pairs []string) (tappedLabels TappedLabels) {
	tappedLabels = make(TappedLabels, 0)
	for _, v := range pairs {
		tappedLabels[v] = DisplayedState{}
	}
	return tappedLabels
}

func (l TappedLabels) GetTrueLabels() (labels []string) {
	for label, tapped := range l {
		if tapped.Tapped {
			labels = append(labels, label)
		}
	}
	return labels
}

func (l TappedLabels) PairIsTapped(labels []string) (isTapped bool) {
	return l[labels[0]].Tapped && l[labels[1]].Tapped
}

func (l TappedLabels) GetLabels() (labels []string) {
	for label := range l {
		labels = append(labels, label)
	}
	return labels
}

func (l TappedLabels) ClearTappedLabels() {
	for label := range l {
		state := l[label]
		state.Tapped = false
		l[label] = state
	}
}

func arraysEqual(arr1, arr2 []string) bool {
	if len(arr1) == 0 || len(arr2) == 0 || len(arr1) != len(arr2) {
		return false
	}

	counts := make(map[string]int)

	for _, v := range arr1 {
		counts[v]++
	}

	for _, v := range arr2 {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}

func (substanceMap SubstanceMap) GetSubstance(pair []string) (substance Substance, exists bool) {
	for s, makeup := range substanceMap {
		if arraysEqual(makeup.Labels, pair) {
			return s, true
		}
	}
	return "", false
}

/*
substance:
	components:
		- component1
		- component2
*/

//go:embed substances.yaml
var embeddedFiles embed.FS

func (substanceMap SubstanceMap) LoadYAML(filePath string) (err error) {
	data, err := embeddedFiles.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded file: %w", err)
	}

	var raw map[string]struct {
		Components []string `yaml:"components"`
	}

	if err := yaml.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	for substance, details := range raw {
		substanceMap[Substance(substance)] = SubstanceMakeup{
			Labels:     details.Components,
			Discovered: false,
		}
	}

	return nil
}
