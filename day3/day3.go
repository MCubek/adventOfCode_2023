package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isSymbol(ch rune) bool {
	return !unicode.IsDigit(ch) && ch != '.'
}

func sumPartNumbers(schematic string) int {
	lines := strings.Split(schematic, "\n")
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	sum := 0
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			ch := rune(line[x])

			if unicode.IsDigit(ch) {
				num, xx := parseFullNumberAndLastPosition(x, line)

				// Check adjacent cells for symbols
				isAdjacentToSymbol := false

				for xc := x; xc < xx; xc++ {
					for d := 0; d < 8; d++ {
						ny := y + dy[d]
						nx := xc + dx[d]

						if ny >= 0 && ny < len(lines) && nx >= 0 && nx < len(lines[ny]) && isSymbol(rune(lines[ny][nx])) {
							isAdjacentToSymbol = true
							break
						}
					}
					if isAdjacentToSymbol {
						break
					}
				}

				if isAdjacentToSymbol {
					sum += num
				}
				x = xx - 1 // Skip the rest of the digits already processed
			}
		}
	}
	return sum
}

func parseFullNumberAndLastPosition(x int, line string) (int, int) {
	for x > 0 && unicode.IsDigit(rune(line[x-1])) {
		x -= 1
	}

	numStr := string(rune(line[x]))
	// Scan for the full number
	xx := x + 1
	for xx < len(line) && unicode.IsDigit(rune(line[xx])) {
		numStr += string(line[xx])
		xx++
	}

	num, _ := strconv.Atoi(numStr)

	return num, xx
}

func sumGearRatios(schematic string) int {
	lines := strings.Split(schematic, "\n")
	numberProcessed := make([][]bool, len(lines))
	for i := range numberProcessed {
		numberProcessed[i] = make([]bool, len(lines[0]))
	}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	totalGearRatio := 0

	for y, line := range lines {
		for x, ch := range line {
			if ch == '*' {
				partNumbers := make([]int, 0)

				// Check adjacent cells for part numbers
				for d := 0; d < 8; d++ {
					ny := y + dy[d]
					nx := x + dx[d]

					if ny >= 0 && ny < len(lines) && nx >= 0 && nx < len(lines[ny]) && !numberProcessed[ny][nx] {
						num, lastPos := parseFullNumberAndLastPosition(nx, lines[ny])
						if num != 0 && lastPos > nx {
							partNumbers = append(partNumbers, num)
							// Mark the number as processed
							for pos := nx; pos < lastPos; pos++ {
								numberProcessed[ny][pos] = true
							}
						}
					}
				}

				if len(partNumbers) == 2 {
					totalGearRatio += partNumbers[0] * partNumbers[1]
				}
			}
		}
	}

	return totalGearRatio
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
	fileContent, err := os.ReadFile("day3/input/parts.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	schematic := string(fileContent)

	if task == 1 {
		sum := sumPartNumbers(schematic)
		fmt.Println("Sum of part numbers:", sum)
	} else if task == 2 {
		sum := sumGearRatios(schematic)
		fmt.Println("Sum of gear ratios:", sum)
	}
}
