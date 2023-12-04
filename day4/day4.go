package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// User input for task selection
	var task int
	fmt.Println("Enter 1 for fist part, 2 for second.")
	_, err := fmt.Scan(&task)
	if err != nil {
		fmt.Println("Error receiving input:", err)
	}

	// Open the input document
	fileContent, err := os.ReadFile("day4/inputs/cards.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	cards := string(fileContent)

	if task == 1 {
		points := calculateWins(cards)
		fmt.Println("Points:", points)
	} else if task == 2 {
		noOfCards := calculateNoOfCards(cards)
		fmt.Println("Cards:", noOfCards)
	}
}

func parseCards(cards string) ([]string, []string) {
	cardLines := strings.Split(strings.TrimSpace(cards), "\n")
	var numbersList []string
	var picksList []string

	for _, line := range cardLines {
		split := strings.Split(line, ":")
		parts := strings.Split(split[1], "|")
		if len(parts) != 2 {
			continue // Skip invalid lines
		}

		numbersList = append(numbersList, parts[0])
		picksList = append(picksList, parts[1])
	}

	return numbersList, picksList
}

func calculateNoOfCards(cards string) int {
	numbers, picks := parseCards(cards)

	sum := 0

	for i := 0; i < len(cards); i++ {
		sum += calculateNoOfCardsRecursive(numbers, picks, i)
	}
	return sum
}

func calculateNoOfCardsRecursive(numbers []string, picks []string, line int) int {
	if line >= len(numbers) {
		return 0
	}

	matches := calculateGameWins(numbers[line], picks[line])

	count := 0
	for i := line + 1; i <= line+matches; i++ {
		count += calculateNoOfCardsRecursive(numbers, picks, i)
	}
	return count + 1
}

func calculateWins(cards string) int {
	points := 0

	for _, game := range strings.Split(strings.TrimSuffix(cards, "\n"), "\n") {

		numbers := strings.Split(game, ":")[1]
		picks := strings.Split(numbers, "|")

		result := calculateGameWins(picks[0], picks[1])
		points += int(math.Pow(2, float64(result-1)))
	}
	return points
}

func calculateGameWins(winningNumbers string, pickedNumbers string) int {
	// Split the string of numbers into slices by Whitespaces.
	winningNumbersSlice := strings.Fields(winningNumbers)
	pickedNumbersSlice := strings.Fields(pickedNumbers)

	// Create a map to store the winning numbers for easy lookup.
	winNumbersMap := make(map[int]bool)

	// Convert winning numbers from string to integer and store in the map.
	for _, num := range winningNumbersSlice {
		number, _ := strconv.Atoi(num)
		winNumbersMap[number] = true
	}

	matchingNumbers := 0

	// Check each picked number if it is in the winning numbers map.
	for _, num := range pickedNumbersSlice {
		number, _ := strconv.Atoi(num)

		if winNumbersMap[number] {
			matchingNumbers += 1
		}
	}

	// Return the total of matching numbers.
	return matchingNumbers
}
