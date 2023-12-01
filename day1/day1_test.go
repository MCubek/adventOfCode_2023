package main

import (
	"testing"
)

func TestExtractSpelledOutDigits(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	var sum int

	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			got := extractSpelledOutDigits(test.line)
			if got != test.expected {
				t.Errorf("ExtractSpelledOutDigits(%q) = %d; want %d", test.line, got, test.expected)
			}
			sum += got
		})
	}
	t.Run("SumCorect", func(t *testing.T) {
		if sum != 281 {
			t.Errorf("Expected sum of 281 but got %d", sum)
		}
	})
}
