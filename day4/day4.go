package main

import (
	"bufio"
	"fmt"
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
	file, err := os.Open("day4/inputs/cards.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)

	if task == 1 {
		points := 0

		for scanner.Scan() {
			game := scanner.Text()

			numbers := strings.Split(game, ":")[1]
			picks := strings.Split(numbers, "|")

			points += calculateWins(picks[0], picks[1])
		}
		fmt.Println("Points:", points)
	} else if task == 2 {
	}
}

func calculateWins(winningNumbers string, pickedNumbers string) int {
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
			if matchingNumbers == 0 {
				matchingNumbers = 1
			} else {
				matchingNumbers *= 2
			}
		}
	}

	// Return the total of matching numbers.
	return matchingNumbers
}
