package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeMap map[int]int

func readAlmanac(input string) ([]int, []RangeMap) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var seeds []int
	maps := []RangeMap{{}}

	// Process seeds
	seedString := strings.Fields(lines[0])
	for _, s := range seedString[1:] {
		seed, err := strconv.Atoi(s)
		if err != nil {
			return nil, nil
		}
		seeds = append(seeds, seed)
	}

	// Process maps
	for _, line := range lines[1:] {
		if line == "" {
			maps = append(maps, RangeMap{})
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}

		startDest, _ := strconv.Atoi(parts[0])
		startSrc, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])

		for i := 0; i < length; i++ {
			maps[len(maps)-1][startSrc+i] = startDest + i
		}
	}

	return seeds, maps
}

func convertThroughMaps(value int, maps []RangeMap) int {
	for _, m := range maps {
		if newVal, ok := m[value]; ok {
			value = newVal
		}
	}
	return value
}

func main() {
	// User input for task selection
	var task int
	fmt.Println("Enter 1 for fist part, 2 for second.")
	_, err := fmt.Scan(&task)
	if err != nil {
		fmt.Println("Error receiving input:", err)
	}

	// Open the input document
	fileContent, err := os.ReadFile("day5/inputs/almanac.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	almanac := string(fileContent)

	seeds, maps := readAlmanac(almanac)

	lowestLocation := findLowestLocation(seeds, maps)

	fmt.Println("Lowest location number:", lowestLocation)
}

func findLowestLocation(seeds []int, maps []RangeMap) int {
	lowestLocation := -1
	for _, seed := range seeds {
		location := convertThroughMaps(seed, maps)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}
