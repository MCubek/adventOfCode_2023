package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeMapping struct {
	StartSrc  int
	StartDest int
	Length    int
}

type RangeMap []RangeMapping

func readAlmanac(input string) ([]int, []RangeMap) {
	seeds := readSeeds1(input)
	maps := readMaps(input)

	return seeds, maps
}
func readMaps(input string) []RangeMap {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var maps []RangeMap
	var currentMap RangeMap

	// Process maps
	for _, line := range lines[1:] {
		if line == "" {
			maps = append(maps, currentMap)
			currentMap = RangeMap{}
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}

		startDest, _ := strconv.Atoi(parts[0])
		startSrc, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])

		currentMap = append(currentMap, RangeMapping{StartSrc: startSrc, StartDest: startDest, Length: length})
	}
	maps = append(maps, currentMap) // Don't forget to append the last map

	return maps
}

func readSeeds1(input string) []int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var seeds []int

	// Process seeds
	seedString := strings.Fields(lines[0])
	for _, s := range seedString[1:] {
		seed, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func readSeeds2(input string) []int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var seeds []int

	// Process seeds
	seedString := strings.Fields(lines[0])
	for i := 1; i < len(seedString); i += 2 {
		startingSeed, err := strconv.Atoi(seedString[i])
		if err != nil {
			return nil
		}
		seedRange, err := strconv.Atoi(seedString[i+1])
		if err != nil {
			return nil
		}

		for j := 0; j < seedRange; j++ {
			seeds = append(seeds, startingSeed+j)
		}
	}
	return seeds
}

func convertThroughMaps(value int, maps []RangeMap) int {
	for _, m := range maps {
		converted := false
		for _, mapping := range m {
			if value >= mapping.StartSrc && value < mapping.StartSrc+mapping.Length {
				value = mapping.StartDest + (value - mapping.StartSrc)
				converted = true
				break
			}
		}
		if !converted {
			// If not converted, it remains the same
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

	var lowestLocation int
	if task == 1 {
		seeds, maps := readAlmanac(almanac)

		lowestLocation = findLowestLocation(seeds, maps)
	} else {
		seeds := readSeeds2(almanac)
		maps := readMaps(almanac)

		lowestLocation = findLowestLocation(seeds, maps)
	}

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
