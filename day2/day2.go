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
	file, err := os.Open("day2/inputs/games.txt")
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

	if task == 1 {
		redLimit, greenLimit, blueLimit := 12, 13, 14
		sumOfGamesIDs := 0

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			game := scanner.Text()

			parts := strings.Split(game, ":")
			gameID, _ := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
			gameData := parts[1]

			if isGamePossible(gameData, redLimit, greenLimit, blueLimit) {
				sumOfGamesIDs += gameID
			}
		}

		fmt.Println("Sum of game IDs:", sumOfGamesIDs)
	} else if task == 2 {
		sumOfGamesPowers := 0

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			game := scanner.Text()

			minRed, minGreen, minBlue := minCubes(game)
			power := minRed * minGreen * minBlue
			sumOfGamesPowers += power
		}

		fmt.Println("Total power:", sumOfGamesPowers)
	}

}

func isGamePossible(game string, redLimit int, greenLimit int, blueLimit int) bool {
	subsets := strings.Split(game, ";")

	for _, subset := range subsets {
		reds, greens, blues := 0, 0, 0

		cubes := strings.Split(subset, ",")
		for _, cube := range cubes {

			cube = strings.TrimSpace(cube)
			parts := strings.Split(cube, " ")
			count, _ := strconv.Atoi(parts[0])
			color := parts[1]

			switch color {
			case "red":
				reds += count
			case "green":
				greens += count
			case "blue":
				blues += count
			}
		}

		if reds > redLimit || greens > greenLimit || blues > blueLimit {
			return false
		}
	}

	return true
}

func minCubes(game string) (int, int, int) {
	gameSplit := strings.Split(game, ":")

	subsets := strings.Split(gameSplit[1], ";")

	minRed, minGreen, minBlue := 0, 0, 0

	for _, subset := range subsets {
		reds, greens, blues := 0, 0, 0

		cubes := strings.Split(subset, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			parts := strings.Split(cube, " ")
			count, _ := strconv.Atoi(parts[0])
			color := parts[1]

			switch color {
			case "red":
				reds += count
			case "green":
				greens += count
			case "blue":
				blues += count
			}
		}

		if reds > minRed {
			minRed = reds
		}
		if greens > minGreen {
			minGreen = greens
		}
		if blues > minBlue {
			minBlue = blues
		}
	}
	return minRed, minGreen, minBlue
}
