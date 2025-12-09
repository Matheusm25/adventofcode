package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var body []byte
	var openFileError error

	if len(os.Args) > 1 && os.Args[1] == "test" {
		body, openFileError = os.ReadFile("src/day06/example.txt")
	} else {
		body, openFileError = os.ReadFile("src/day06/input.txt")
	}

	if openFileError != nil {
		log.Fatalf("unable to read file: %v", openFileError)
	}

	processPartOne(body)
	processPartTwo(body)
}

type RaceConfig struct {
	time     int
	distance int
}

func parseInput(input string, kerning bool) []RaceConfig {
	races := make([]RaceConfig, 0)

	for index, line := range strings.Split(input, "\n") {
		raceIndex := 0

		if kerning {
			line = strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")
		}

		for _, value := range strings.Split(line, " ") {
			if len(races) <= raceIndex {
				var race RaceConfig
				races = append(races, race)
			}
			numberValue, err := strconv.Atoi(value)

			if err == nil {
				if index == 0 {
					races[raceIndex].time = numberValue
				} else if index == 1 {
					races[raceIndex].distance = numberValue
				}
				raceIndex++
			}
		}
	}

	return races
}

func processPartOne(input []byte) {
	races := parseInput(string(input), false)

	racesPossibilitiesResult := 1

	for _, race := range races {
		possibilities := 0
		for index := 0; index < race.time; index++ {
			if (race.time-index)*index > race.distance {
				possibilities++
			}
		}

		if possibilities > 0 {
			racesPossibilitiesResult *= possibilities
		}
	}

	fmt.Printf("The multiplied value of the possibilities is %v\n", racesPossibilitiesResult)
}

func processPartTwo(input []byte) {
	races := parseInput(string(input), true)

	racesPossibilities := 0

	for index := 0; index < races[0].time; index++ {
		if (races[0].time-index)*index > races[0].distance {
			racesPossibilities++
		}
	}

	fmt.Printf("The possibilities for the race be won is %v\n", racesPossibilities)
}
