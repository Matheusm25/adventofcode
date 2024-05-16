package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	body, err := os.ReadFile("src/day03/input.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	processPartOne(body)
	processPartTwo(body)
}

type SchemaNumber struct {
	indexes []string
	value   int
}

func processPartOne(input []byte) {
	var partsSum int

	schema := make([][]string, 0)

	for lineIndex, line := range strings.Split(string(input), "\n") {
		schema = append(schema, make([]string, 0))
		schema[lineIndex] = append(schema[lineIndex], strings.Split(line, "")...)
	}

	var hasNearbySymbol bool = false
	var currentNumber string = ""
	for lineIndex, line := range schema {
		for columnIndex, value := range line {
			if isInteger(value) {
				currentNumber += value

				if !hasNearbySymbol {
					hasNearbySymbol = checkSurroundings(schema, lineIndex, columnIndex)
				}

			} else if currentNumber != "" && hasNearbySymbol {
				partNumber, err := strconv.Atoi(currentNumber)
				if err != nil {
					log.Fatalf("Unable to parse value %v", err)
				}

				partsSum += partNumber
				currentNumber = ""
				hasNearbySymbol = false
			} else if currentNumber != "" {
				currentNumber = ""
				hasNearbySymbol = false
			} else {
				currentNumber = ""
				hasNearbySymbol = false
			}
		}
	}

	fmt.Printf("The sum of the parts is: %v\n", partsSum)
}

func processPartTwo(input []byte) {
	var partsSum int
	schema := make([][]string, 0)

	for lineIndex, line := range strings.Split(string(input), "\n") {
		schema = append(schema, make([]string, 0))
		schema[lineIndex] = append(schema[lineIndex], strings.Split(line, "")...)
	}

	numbers := make([]SchemaNumber, 0)

	var currentNumber string = ""
	var currentSchemaNumber SchemaNumber
	for lineIndex, line := range schema {
		if currentNumber != "" {
			partNumber, err := strconv.Atoi(currentNumber)
			if err != nil {
				log.Fatalf("Unable to parse %v", err)
			}
			currentSchemaNumber.value = partNumber

			numbers = append(numbers, currentSchemaNumber)
			currentSchemaNumber = SchemaNumber{}
			currentNumber = ""
		}

		for columnIndex, value := range line {
			if isInteger(value) {
				currentNumber += value
				currentSchemaNumber.indexes = append(currentSchemaNumber.indexes, fmt.Sprintf("[%v,%v]", lineIndex, columnIndex))
			} else if currentNumber != "" {
				partNumber, err := strconv.Atoi(currentNumber)
				if err != nil {
					log.Fatalf("Unable to parse %v", err)
				}

				currentSchemaNumber.value = partNumber
				numbers = append(numbers, currentSchemaNumber)
				currentSchemaNumber = SchemaNumber{}
				currentNumber = ""
			}
		}
	}

	for lineIndex, line := range schema {
		for columnIndex, item := range line {
			if item == "*" {
				numberArround := getNumbersSurrounding(numbers, lineIndex, columnIndex)

				if len(numberArround) == 2 {
					partsSum += numberArround[0] * numberArround[1]
				}
			}
		}
	}

	fmt.Printf("The sum of the gears ratio is %v\n", partsSum)

}

func getNumbersSurrounding(numbers []SchemaNumber, lineIndex int, columnIndex int) []int {
	numbersArround := make([]int, 0)

	for _, n := range numbers {
		if slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex-1, columnIndex-1)) ||
			slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex-1, columnIndex)) ||
			slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex-1, columnIndex+1)) ||
			slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex, columnIndex-1)) ||
			slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex, columnIndex+1)) ||
			slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex+1, columnIndex-1)) ||
			slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex+1, columnIndex)) ||
			slices.Contains(n.indexes, fmt.Sprintf("[%v,%v]", lineIndex+1, columnIndex+1)) {
			numbersArround = append(numbersArround, n.value)
		}
	}

	return numbersArround
}

func checkSurroundings(schema [][]string, lineIndex int, columnIndex int) bool {
	var possibilities = make([]string, 0)

	if lineIndex == 0 && columnIndex == 0 {
		possibilities = []string{
			schema[lineIndex][columnIndex+1],
			schema[lineIndex+1][columnIndex+1],
			schema[lineIndex+1][columnIndex],
		}
	} else if lineIndex > 0 && lineIndex < len(schema)-1 && columnIndex > 0 && columnIndex < len(schema[lineIndex])-1 {
		possibilities = []string{
			schema[lineIndex-1][columnIndex-1],
			schema[lineIndex-1][columnIndex],
			schema[lineIndex-1][columnIndex+1],

			schema[lineIndex][columnIndex-1],
			schema[lineIndex][columnIndex+1],

			schema[lineIndex+1][columnIndex-1],
			schema[lineIndex+1][columnIndex],
			schema[lineIndex+1][columnIndex+1],
		}
	} else if lineIndex == 0 && columnIndex < len(schema[lineIndex])-1 {
		possibilities = []string{
			schema[lineIndex][columnIndex-1],
			schema[lineIndex][columnIndex+1],

			schema[lineIndex+1][columnIndex-1],
			schema[lineIndex+1][columnIndex],
			schema[lineIndex+1][columnIndex+1],
		}
	} else if lineIndex == 0 && columnIndex == len(schema[lineIndex])-1 {
		possibilities = []string{
			schema[lineIndex][columnIndex-1],

			schema[lineIndex+1][columnIndex-1],
			schema[lineIndex+1][columnIndex],
		}
	} else if columnIndex == 0 && lineIndex < len(schema)-1 {
		possibilities = []string{
			schema[lineIndex-1][columnIndex],
			schema[lineIndex-1][columnIndex+1],

			schema[lineIndex][columnIndex+1],

			schema[lineIndex+1][columnIndex],
			schema[lineIndex+1][columnIndex+1],
		}
	} else if columnIndex == 0 && lineIndex == len(schema)-1 {
		possibilities = []string{
			schema[lineIndex][columnIndex+1],
			schema[lineIndex-1][columnIndex+1],
			schema[lineIndex-1][columnIndex],
		}
	} else if lineIndex == len(schema)-1 && columnIndex == len(schema[lineIndex])-1 {
		possibilities = []string{
			schema[lineIndex-1][columnIndex],
			schema[lineIndex-1][columnIndex-1],
			schema[lineIndex][columnIndex-1],
		}
	} else if lineIndex == len(schema)-1 && columnIndex < len(schema[lineIndex])-1 {
		possibilities = []string{
			schema[lineIndex][columnIndex-1],
			schema[lineIndex][columnIndex+1],

			schema[lineIndex-1][columnIndex-1],
			schema[lineIndex-1][columnIndex],
			schema[lineIndex-1][columnIndex+1],
		}
	} else if columnIndex == len(schema[lineIndex])-1 {
		possibilities = []string{
			schema[lineIndex-1][columnIndex],
			schema[lineIndex-1][columnIndex-1],
			schema[lineIndex][columnIndex-1],
			schema[lineIndex+1][columnIndex-1],
			schema[lineIndex+1][columnIndex],
		}
	}

	hasSymbol := false

	for _, value := range possibilities {
		if !isInteger(value) && value != "." {
			hasSymbol = true
			break
		}
	}

	return hasSymbol
}

func isInteger(char string) bool {
	options := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	return slices.Contains(options, char)
}
