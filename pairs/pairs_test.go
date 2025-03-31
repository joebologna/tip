package pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairIsTapped(t *testing.T) {
	pairs := []string{"oxygen", "hydrogen", "salt", "water"}
	tappedLabels := NewTappedLabels(pairs)

	// Test when both labels are not tapped
	assert.False(t, tappedLabels.PairIsTapped([]string{"oxygen", "hydrogen"}), "Expected PairIsTapped to be false when both labels are not tapped")

	// Test when one label is tapped
	tappedLabels["oxygen"] = DisplayedState{Tapped: true}
	assert.False(t, tappedLabels.PairIsTapped([]string{"oxygen", "hydrogen"}), "Expected PairIsTapped to be false when one label is tapped")

	// Test when both labels are tapped
	tappedLabels["hydrogen"] = DisplayedState{Tapped: true}
	assert.True(t, tappedLabels.PairIsTapped([]string{"oxygen", "hydrogen"}), "Expected PairIsTapped to be true when both labels are tapped")

	// Test when labels do not exist
	assert.False(t, tappedLabels.PairIsTapped([]string{"x", "y"}), "Expected PairIsTapped to be false when labels do not exist")

	// Test for brine
	tappedLabels.ClearTappedLabels()
	assert.False(t, tappedLabels.PairIsTapped([]string{"water", "salt"}), "Expected PairIsTapped to be false when both labels are not tapped for brine")
	tappedLabels["water"] = DisplayedState{Tapped: true}
	assert.False(t, tappedLabels.PairIsTapped([]string{"water", "salt"}), "Expected PairIsTapped to be false when one label is tapped for brine")
	tappedLabels["salt"] = DisplayedState{Tapped: true}
	assert.True(t, tappedLabels.PairIsTapped([]string{"water", "salt"}), "Expected PairIsTapped to be true when both labels are tapped for brine")

	// Test for water
	tappedLabels.ClearTappedLabels()
	assert.False(t, tappedLabels.PairIsTapped([]string{"oxygen", "hydrogen"}), "Expected PairIsTapped to be false when both labels are not tapped for water")
	tappedLabels["oxygen"] = DisplayedState{Tapped: true}
	assert.False(t, tappedLabels.PairIsTapped([]string{"oxygen", "hydrogen"}), "Expected PairIsTapped to be false when one label is tapped for water")
	tappedLabels["hydrogen"] = DisplayedState{Tapped: true}
	assert.True(t, tappedLabels.PairIsTapped([]string{"oxygen", "hydrogen"}), "Expected PairIsTapped to be true when both labels are tapped for water")

	// Clear all tapped labels and test again
	tappedLabels.ClearTappedLabels()
	assert.False(t, tappedLabels.PairIsTapped([]string{"oxygen", "hydrogen"}), "Expected PairIsTapped to be false after clearing tapped labels")
}

func TestGetLabels(t *testing.T) {
	pairs := []string{"oxygen", "hydrogen"}
	tappedLabels := NewTappedLabels(pairs)

	labels := tappedLabels.GetLabels()
	assert.Len(t, labels, 2, "Expected length of labels to be 2")
	assert.Contains(t, labels, "oxygen", "Expected labels to contain 'oxygen'")
	assert.Contains(t, labels, "hydrogen", "Expected labels to contain 'hydrogen'")

	// Test when no labels are tapped
	tappedLabels.ClearTappedLabels()
	labels = tappedLabels.GetLabels()
	assert.Len(t, labels, 2, "Expected length of labels to be 2 after clearing tapped labels")
	assert.Contains(t, labels, "oxygen", "Expected labels to contain 'oxygen' after clearing tapped labels")
	assert.Contains(t, labels, "hydrogen", "Expected labels to contain 'hydrogen' after clearing tapped labels")
}

func TestClearTappedLabels(t *testing.T) {
	pairs := []string{"oxygen", "hydrogen"}
	tappedLabels := NewTappedLabels(pairs)

	// Tap some labels
	tappedLabels["oxygen"] = DisplayedState{Tapped: true}
	tappedLabels["hydrogen"] = DisplayedState{Tapped: true}

	// Clear all tapped labels
	tappedLabels.ClearTappedLabels()

	// Test that all labels are cleared
	for _, tapped := range tappedLabels {
		assert.False(t, tapped.Tapped, "Expected all labels to be cleared")
	}

	// Test when no labels are tapped
	tappedLabels.ClearTappedLabels()
	for _, tapped := range tappedLabels {
		assert.False(t, tapped.Tapped, "Expected all labels to be cleared when no labels are tapped")
	}
}

func TestGetSubstance(t *testing.T) {
	substanceMap := SubstanceMap{
		"water": {Labels: []string{"oxygen", "hydrogen"}},
		"brine": {Labels: []string{"water", "salt"}},
	}

	tests := []struct {
		pair     []string
		expected Substance
		exists   bool
	}{
		{[]string{"oxygen", "hydrogen"}, "water", true},
		{[]string{"hydrogen", "oxygen"}, "water", true},
		{[]string{"carbon", "oxygen"}, "", false},
		{[]string{}, "", false}, // Test with an empty pair
	}

	for _, test := range tests {
		substance, exists := substanceMap.GetSubstance(test.pair)
		assert.Equal(t, test.expected, substance, "they should be equal")
		assert.Equal(t, test.exists, exists, "they should be equal")
	}
}

func TestArraysEqual(t *testing.T) {
	tests := []struct {
		arr1     []string
		arr2     []string
		expected bool
	}{
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "d"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{[]string{}, []string{"a", "b", "c", "d"}, false},
		{[]string{}, []string{}, false},
		{[]string{"a", "b", "c"}, []string{}, false}, // Test with one empty array and one non-empty array
	}

	for _, test := range tests {
		result := arraysEqual(test.arr1, test.arr2)
		assert.Equal(t, test.expected, result, "arraysEqual(%v, %v) should be %v", test.arr1, test.arr2, test.expected)
	}
}

func TestGetTrueLabels(t *testing.T) {
	pairs := []string{"oxygen", "hydrogen", "salt", "water"}
	tappedLabels := NewTappedLabels(pairs)

	// Initially, no labels are tapped
	trueLabels := tappedLabels.GetTrueLabels()
	assert.Empty(t, trueLabels, "Expected no true labels initially")

	// Tap some labels
	tappedLabels["oxygen"] = DisplayedState{Tapped: true}
	tappedLabels["water"] = DisplayedState{Tapped: true}

	// Get true labels
	trueLabels = tappedLabels.GetTrueLabels()
	assert.Len(t, trueLabels, 2, "Expected length of true labels to be 2")
	assert.Contains(t, trueLabels, "oxygen", "Expected true labels to contain 'oxygen'")
	assert.Contains(t, trueLabels, "water", "Expected true labels to contain 'water'")

	// Tap another label
	tappedLabels["salt"] = DisplayedState{Tapped: true}

	// Get true labels
	trueLabels = tappedLabels.GetTrueLabels()
	assert.Len(t, trueLabels, 3, "Expected length of true labels to be 3")
	assert.Contains(t, trueLabels, "salt", "Expected true labels to contain 'salt'")

	// Test when no labels are tapped
	tappedLabels.ClearTappedLabels()
	trueLabels = tappedLabels.GetTrueLabels()
	assert.Empty(t, trueLabels, "Expected no true labels after clearing tapped labels")
}

func TestFound(t *testing.T) {
	substanceMap := SubstanceMap{
		"water": {Labels: []string{"oxygen", "hydrogen"}, Discovered: true},
		"brine": {Labels: []string{"water", "salt"}, Discovered: false},
	}

	ratio := substanceMap.Found()
	assert.Equal(t, "1/2", ratio, "Expected Found ratio to be '1/2'")
}

func TestNewTappedLabels(t *testing.T) {
	pairs := []string{"oxygen", "hydrogen", "salt", "water"}
	tappedLabels := NewTappedLabels(pairs)

	assert.Len(t, tappedLabels, 4, "Expected length of tappedLabels to be 4")
	for _, label := range pairs {
		assert.Contains(t, tappedLabels, label, "Expected tappedLabels to contain '%s'", label)
		assert.False(t, tappedLabels[label].Tapped, "Expected '%s' to be initially not tapped", label)
	}
}
