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

	values := make([]int, 0)

	lines := strings.Split(string(body), "\n")

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

	fmt.Printf("The final number is %v\n", sum)
}

func isInteger(charCode byte) bool {
	if int(charCode) >= 48 && int(charCode) <= 57 {
		return true
	} else {
		return false
	}
}
