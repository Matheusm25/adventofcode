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
	var body []byte
	var openFileError error

	if len(os.Args) > 1 && os.Args[1] == "test" {
		body, openFileError = os.ReadFile("src/day07/example.txt")
	} else {
		body, openFileError = os.ReadFile("src/day07/input.txt")
	}

	if openFileError != nil {
		log.Fatalf("unable to read file: %v", openFileError)
	}

	processPartOne(body)
	processPartTwo(body)
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

type Game struct {
	bid        int
	hand       []string
	handResult int
	rank       int
}

func checkHandValues(hand []string, jokerRule bool) int {
	repeatedCards := make(map[string]int)
	jokers := 0
	result := -1

	for _, card := range hand {
		repeatedCards[card]++
		if card == "J" {
			jokers++
		}
	}

	hasThree := false
	twoCounter := 0
	for _, quantity := range repeatedCards {
		if quantity == 5 {
			result = FIVE_OF_A_KIND
			break
		} else if quantity == 4 {
			result = FOUR_OF_A_KIND
			break
		} else if quantity == 3 && twoCounter == 0 {
			hasThree = true
		} else if quantity == 3 && twoCounter == 1 {
			result = FULL_HOUSE
			break
		} else if quantity == 2 && hasThree {
			result = FULL_HOUSE
			break
		} else if quantity == 2 && twoCounter == 1 {
			result = TWO_PAIR
			break
		} else if quantity == 2 {
			twoCounter++
		}
	}

	if result == -1 && hasThree {
		result = THREE_OF_A_KIND
	} else if result == -1 && twoCounter == 1 {
		result = ONE_PAIR
	} else if result == -1 {
		result = HIGH_CARD
	}

	// preciso conseguir difereciar quando os results
	// são formados por joker ou não
	// caso sejam formados por jokers, tenho que desconsiderar

	if jokerRule {
		if result == FOUR_OF_A_KIND && jokers == 1 {
			result = FIVE_OF_A_KIND
		} else if result == THREE_OF_A_KIND && jokers == 2 {
			result = FIVE_OF_A_KIND
		} else if result == THREE_OF_A_KIND && jokers == 1 {
			result = FOUR_OF_A_KIND
		}

	}

	return result
}

func getCardValue(card string, jokerRule bool) int {
	options := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	if jokerRule {
		options["J"] = 1
	}

	return options[card]
}

func parseInput(input string, jokerRule bool) []Game {
	games := make([]Game, 0)

	for _, gameString := range strings.Split(input, "\n") {
		var game Game
		for index, info := range strings.Split(gameString, " ") {
			if index == 0 {
				game.hand = strings.Split(info, "")
				game.handResult = checkHandValues(game.hand, jokerRule)
			} else if index == 1 {
				numberValue, err := strconv.Atoi(info)
				if err != nil {
					log.Fatalf("Unable to parse value '%v'", err)
				}

				game.bid = numberValue
			}
		}
		games = append(games, game)
	}

	return games
}

func processPartOne(input []byte) {
	games := parseInput(string(input), false)

	gamesResultSum := 0

	slices.SortFunc(games, func(gameA Game, gameB Game) int {
		if gameA.handResult != gameB.handResult {
			return gameA.handResult - gameB.handResult
		} else {
			var result int

			for index := range gameA.hand {
				if gameA.hand[index] != gameB.hand[index] {
					result = getCardValue(gameA.hand[index], false) - getCardValue(gameB.hand[index], false)
					break
				} else {
					continue
				}
			}

			return result
		}
	})

	for index, game := range games {
		gamesResultSum += (index + 1) * game.bid
	}

	fmt.Printf("The sum of results in all games is %v\n", gamesResultSum)
}

func processPartTwo(input []byte) {
}
