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
		body, openFileError = os.ReadFile("src/day05/example.txt")
	} else {
		body, openFileError = os.ReadFile("src/day05/input.txt")
	}

	if openFileError != nil {
		log.Fatalf("unable to read file: %v", openFileError)
	}

	processPartOne(body)
	processPartTwo(body)
}

type RangeConfig struct {
	destinationRange int
	sourceRange      int
	rangeLength      int
}

type Map struct {
	source      string
	destination string
	mapper      []RangeConfig
	sourceToMap map[int]int
}

func parseInput(input string) ([]int, []Map) {
	maps := make([]Map, 0)

	seedsAndMapsConfigs := strings.Split(input, "\n\n")

	seeds := make([]int, 0)

	for _, seed := range strings.Split(seedsAndMapsConfigs[0], " ") {
		numberValue, err := strconv.Atoi(seed)
		if err == nil {
			seeds = append(seeds, numberValue)
		}
	}

	for _, mapper := range seedsAndMapsConfigs[1:] {
		var mapInstance Map
		mapInstance.sourceToMap = make(map[int]int)

		for index, line := range strings.Split(mapper, "\n") {
			if index == 0 {
				infos := strings.Split(strings.Split(line, " ")[0], "-")
				mapInstance.source = infos[0]
				mapInstance.destination = infos[2]
			} else {
				var rangeconfig RangeConfig
				configs := strings.Split(line, " ")

				for configIndex, configValue := range configs {
					configNumberValue, err := strconv.Atoi(configValue)
					if err != nil {
						log.Fatalf("Unable to parse value '%v'", err)
					}

					switch configIndex {
					case 0:
						rangeconfig.destinationRange = configNumberValue
					case 1:
						rangeconfig.sourceRange = configNumberValue
					case 2:
						rangeconfig.rangeLength = configNumberValue
					}
				}

				mapInstance.mapper = append(mapInstance.mapper, rangeconfig)
			}
		}

		maps = append(maps, mapInstance)
	}

	return seeds, maps
}

func processPartOne(input []byte) {
	seeds, maps := parseInput(string(input))
	lowerSeedLocation := -1

	for _, seed := range seeds {
		currentValue := seed

		for _, mapper := range maps {
			for _, checkRange := range mapper.mapper {
				if currentValue >= checkRange.sourceRange &&
					currentValue < checkRange.sourceRange+checkRange.rangeLength {
					currentValue = checkRange.destinationRange + int(math.Abs(float64(currentValue)-float64(checkRange.sourceRange)))
					break
				}
			}
			if mapper.destination == "location" {
				if lowerSeedLocation > currentValue || lowerSeedLocation == -1 {
					lowerSeedLocation = currentValue
				}
				break
			}
		}
	}

	fmt.Printf("The lowest location found was %v\n", lowerSeedLocation)
}

func processPartTwo(input []byte) {
	seedsConfig, maps := parseInput(string(input))

	slices.Reverse(maps)

	initialValue := 0
	for {
		seedFound := false
		currentSeedCheck := -1
		currentValue := initialValue

		for _, mapper := range maps {
			for _, checkRange := range mapper.mapper {
				if currentValue >= checkRange.destinationRange &&
					currentValue < checkRange.destinationRange+checkRange.rangeLength {
					currentValue = checkRange.sourceRange + int(math.Abs(float64(currentValue)-float64(checkRange.destinationRange)))
					break
				}
			}
			if mapper.source == "seed" {
				currentSeedCheck = currentValue
				break
			}
		}

		seedRangeStart := 0
		for index, seedConfig := range seedsConfig {
			if index%2 == 0 {
				seedRangeStart = seedConfig
			} else {
				if currentSeedCheck >= seedRangeStart && currentSeedCheck < seedRangeStart+seedConfig {
					seedFound = true
				}
			}
		}

		fmt.Printf("Location %v has seed %v \n", initialValue, currentSeedCheck)
		if seedFound || currentValue == 100 {
			fmt.Printf("The lowest location is %v with seed %v\n", initialValue, currentSeedCheck)
			break
		} else {
			initialValue++
		}
	}
}
