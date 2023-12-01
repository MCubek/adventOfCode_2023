package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Current directory is: ", dir)

	// User input for task selection
	var task int
	fmt.Println("Enter 1 for fist part, 2 for second.")
	_, err = fmt.Scan(&task)
	if err != nil {
		fmt.Println("Error receiving input:", err)
	}

	// Open the calibration document
	file, err := os.Open("day1/inputs/calibration.txt")
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

	var sum int
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()

		var value int
		if task == 1 {
			value = extractCalibrationValue(line)
		} else if task == 2 {
			value = extractSpelledOutDigits(line)
		}

		sum += value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total sum of calibration values:", sum)
}

// extractCalibrationValue extracts the calibration value from a line of text (task 1)
func extractCalibrationValue(line string) int {
	return extractDigits(line)
}

// extractSpelledOutDigits extracts calibration value considering spelled-out digits (task 2)
func extractSpelledOutDigits(line string) int {
	digitMap := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
		"six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	var builder strings.Builder
	for i := 0; i < len(line); {
		matched := false
		for word, digit := range digitMap {
			if strings.HasPrefix(line[i:], word) {
				builder.WriteString(digit)
				//i += len(word)
				matched = true
				break
			}
		}
		if !matched {
			if line[i] >= '0' && line[i] <= '9' {
				builder.WriteByte(line[i])
			}
		}
		i++
	}

	return extractDigits(builder.String())
}

// extractDigits finds the first and last numeric digit in a string
func extractDigits(s string) int {
	var firstDigit, lastDigit string
	foundFirst := false

	for _, char := range s {
		if char >= '0' && char <= '9' {
			if !foundFirst {
				firstDigit = string(char)
				foundFirst = true
			}
			lastDigit = string(char)
		}
	}

	if firstDigit != "" && lastDigit != "" {
		value, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return 0
		}
		return value
	}

	return 0
}
