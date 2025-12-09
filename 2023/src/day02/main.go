package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	body, err := os.ReadFile("src/day02/input.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	processPartOne(body)
	processPartTwo(body)
}

func processPartOne(input []byte) {
	var validGamesIdSum int

	colorsLimit := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := strings.Split(string(input), "\n")

	for _, game := range games {
		validGame := true
		gameInfo := strings.Split(game, ":")

		gameId, err := strconv.Atoi(strings.Split(gameInfo[0], " ")[1])
		if err != nil {
			log.Fatalf("Unable to parse value %v", err)
		}

		for _, set := range strings.Split(gameInfo[1], ";") {
			gameResult := make(map[string]int)
			for _, value := range strings.Split(set, ",") {
				info := strings.Split(strings.Trim(value, " "), " ")
				qty, err := strconv.Atoi(info[0])
				if err != nil {
					log.Fatalf("Unable to parse value %v", err)
				}
				gameResult[info[1]] += qty
			}

			if !validateGame(colorsLimit, gameResult) {
				validGame = false
				break
			}
		}

		if validGame {
			validGamesIdSum += gameId
		}
	}

	fmt.Printf("The sum of the ids in the valid games is: %v\n", validGamesIdSum)
}

func processPartTwo(input []byte) {
	var powerSum int

	games := strings.Split(string(input), "\n")

	for _, game := range games {
		gameInfo := strings.Split(game, ":")
		colorLimit := make(map[string]int)

		for _, set := range strings.Split(gameInfo[1], ";") {
			for _, value := range strings.Split(set, ",") {
				info := strings.Split(strings.Trim(value, " "), " ")
				qty, err := strconv.Atoi(info[0])
				if err != nil {
					log.Fatalf("Unable to parse value %v", err)
				}

				if colorLimit[info[1]] < qty {
					colorLimit[info[1]] = qty
				}
			}
		}

		var gamePower int

		for _, value := range colorLimit {
			if gamePower == 0 {
				gamePower = value
			} else {
				gamePower *= value
			}
		}

		powerSum += gamePower
	}

	fmt.Printf("The sum of the power in the games is: %v\n", powerSum)

}

func validateGame(rule map[string]int, gameResult map[string]int) bool {
	isValid := true

	for key := range rule {
		if gameResult[key] > rule[key] {
			isValid = false
		}
	}

	return isValid
}
