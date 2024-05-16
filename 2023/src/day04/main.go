package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var body []byte
	var openFileError error

	if len(os.Args) > 1 && os.Args[1] == "test" {
		body, openFileError = os.ReadFile("src/day04/example.txt")
	} else {
		body, openFileError = os.ReadFile("src/day04/input.txt")
	}

	if openFileError != nil {
		log.Fatalf("unable to read file: %v", openFileError)
	}

	processPartOne(body)
	processPartTwo(body)
}

type Card struct {
	winningNumbers []int
	ownNumbers     []int
	id             string
	instances      int
}

func parseInput(input string) []Card {
	cards := make([]Card, 0)

	for _, line := range strings.Split(input, "\n") {
		cardInfo := strings.Split(line, ":")

		var card Card
		card.instances = 1
		card.id = strings.Split(cardInfo[0], "Card ")[1]

		for index, numbers := range strings.Split(cardInfo[1], "|") {
			for _, n := range strings.Split(numbers, " ") {
				if n == "" {
					continue
				}

				numberValue, err := strconv.Atoi(n)
				if err != nil {
					log.Fatalf("Unable to parse value %v", err)
				}

				if index == 1 {
					card.winningNumbers = append(card.winningNumbers, numberValue)
				} else {
					card.ownNumbers = append(card.ownNumbers, numberValue)
				}
			}
		}

		cards = append(cards, card)
	}

	return cards
}

func processPartOne(input []byte) {
	points := 0

	cards := parseInput(string(input))

	for _, card := range cards {
		winsCounter := 0

		for _, n := range card.ownNumbers {
			if slices.Contains(card.winningNumbers, n) {
				winsCounter++
			}
		}

		if winsCounter > 0 {
			points += int(math.Pow(2, float64(winsCounter)-1))
		}
	}

	fmt.Printf("The total points is %v\n", points)

}

func processPartTwo(input []byte) {
	cards := parseInput(string(input))

	sumInstances := 0

	for cardIndex, card := range cards {
		winsCounter := 0

		for _, n := range card.ownNumbers {
			if slices.Contains(card.winningNumbers, n) {
				winsCounter++
			}
		}

		for i := 1; i <= winsCounter; i++ {
			cards[cardIndex+i].instances += card.instances
		}
	}

	for _, card := range cards {
		sumInstances += card.instances
	}

	fmt.Printf("The total number of cards is %v\n", sumInstances)
}
