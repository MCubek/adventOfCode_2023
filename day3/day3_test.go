package main

import (
	"testing"
)

func TestSumPartNumbers(t *testing.T) {
	schematic := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	expected := 4361
	result := sumPartNumbers(schematic)

	if result != expected {
		t.Errorf("sumPartNumbers() = %d; want %d", result, expected)
	}
}
