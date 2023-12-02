package main

import (
	"testing"
)

func TestMinCubes(t *testing.T) {
	tests := []struct {
		game          string
		expectedRed   int
		expectedGreen int
		expectedBlue  int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 4, 2, 6},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 1, 3, 4},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 20, 13, 6},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 14, 3, 15},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 6, 3, 2},
	}

	for _, test := range tests {
		gotRed, gotGreen, gotBlue := minCubes(test.game)
		if gotRed != test.expectedRed || gotGreen != test.expectedGreen || gotBlue != test.expectedBlue {
			t.Errorf("minCubes(%q) = %d, %d, %d; want %d, %d, %d",
				test.game, gotRed, gotGreen, gotBlue, test.expectedRed, test.expectedGreen, test.expectedBlue)
		}
	}
}
