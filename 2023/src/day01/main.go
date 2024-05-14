package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	body, err := os.ReadFile("src/day01/input.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	processPartOne(body)
	processPartTwo(body)
}

func processPartTwo(input []byte) {
	values := make([]int, 0)

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		initialIndex := 0
		lastIndex := len(line) - 1

		var firstPart string
		var lastPart string

		firstDigit := -1
		lastDigit := -1

		for initialIndex < len(line) {
			if firstDigit == -1 && isInteger(line[initialIndex]) {
				value, err := strconv.Atoi(string(line[initialIndex]))
				if err != nil {
					log.Fatalf("unable to parse value: %v", err)
				}

				firstDigit = value
			}

			if lastDigit == -1 && isInteger(line[lastIndex]) {
				value, err := strconv.Atoi(string(line[lastIndex]))

				if err != nil {
					log.Fatalf("unable to parse value: %v", err)
				}

				lastDigit = value
			}

			firstPart += string(line[initialIndex])
			lastPart = string(line[lastIndex]) + lastPart

			if firstDigit == -1 {
				firstDigit = hasDigit(firstPart)
			}

			if lastDigit == -1 {
				lastDigit = hasDigit(lastPart)
			}

			if firstDigit != -1 && lastDigit != -1 {
				break
			}

			initialIndex++
			lastIndex--
		}

		value, err := strconv.Atoi(fmt.Sprintf("%v%v", firstDigit, lastDigit))

		if err != nil {
			log.Fatalf("unable to parse value: %v", err)
		}

		values = append(values, value)
	}

	sum := 0

	for _, value := range values {
		sum += value
	}

	fmt.Printf("The result of part 1 is:  %v\n", sum)
}

func processPartOne(input []byte) {
	values := make([]int, 0)

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		initialIndex := 0
		lastIndex := len(line) - 1

		var firstDigit string
		var lastDigit string

		for initialIndex < len(line) {
			if firstDigit == "" && isInteger(line[initialIndex]) {
				firstDigit = string(line[initialIndex])
			}

			if lastDigit == "" && isInteger(line[lastIndex]) {
				lastDigit = string(line[lastIndex])
			}

			if firstDigit != "" && lastDigit != "" {
				break
			}

			initialIndex++
			lastIndex--
		}

		value, err := strconv.Atoi(firstDigit + lastDigit)

		if err != nil {
			log.Fatalf("unable to parse value: %v", err)
		}

		values = append(values, value)
	}

	sum := 0

	for _, value := range values {
		sum += value
	}

	fmt.Printf("The result of part 1 is:  %v\n", sum)
}

func hasDigit(line string) int {
	options := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	index := -1

	for i, option := range options {
		if strings.Contains(line, option) {
			index = i
			break
		}
	}

	return index
}

func isInteger(charCode byte) bool {
	if int(charCode) >= 48 && int(charCode) <= 57 {
		return true
	} else {
		return false
	}
}
