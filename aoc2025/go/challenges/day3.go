package challenges

import (
	"log"
	"log/slog"
	"strconv"
	"strings"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

type BatteryPack []int

func Day3() {
	log := logger.New("Day3")
	slog.Info("day3")

	batteryPacks := processDay3File()

	timer := log.Timer("Part 1 Timer")
	part1Result := calculate12PackJoltage(batteryPacks, 2)
	timer()

	timer = log.Timer("Part 2 Timer")
	part2Result := calculate12PackJoltage(batteryPacks, 12)
	timer()

	log.Info("Results", "Part 1", part1Result, "Part 2", part2Result)
}

func calculate12PackJoltage(batteryPacks []BatteryPack, batterySize int) int {
	maxJoltage := 0
	for _, batteryPack := range batteryPacks {
		maxJoltage += batteryPack.getLargestPackJoltage(batterySize)
	}

	return maxJoltage
}

func (b BatteryPack) getLargestPackJoltage(batterySize int) int {
	maxJoltage := make([]int, 0, batterySize)

	index := 0
	iteration := 0
	for {
		if len(maxJoltage) == batterySize {
			break
		}
		maxFoundJoltage, endIndex := iterateBatteryPackCheck(b, index, iteration, batterySize)
		maxJoltage = append(maxJoltage, maxFoundJoltage)
		index = endIndex + 1
		iteration++
	}

	var string strings.Builder
	for _, joltage := range maxJoltage {
		string.WriteString(strconv.Itoa(joltage))
	}

	result, err := strconv.Atoi(string.String())
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func iterateBatteryPackCheck(
	battery BatteryPack,
	startingIndex int,
	currentBatterySize int,
	batterySize int,
) (int, int) {
	maxFoundJoltage := 0
	endIndex := startingIndex
	stop := len(battery) + currentBatterySize - (batterySize - 1)
	for i := startingIndex; i < stop; i++ {
		if battery[i] > maxFoundJoltage {
			maxFoundJoltage = battery[i]
			endIndex = i
		}
	}

	return maxFoundJoltage, endIndex
}

func processDay3File() []BatteryPack {
	file := utilities.ReadFile("day3.part1")
	var batteryPacks []BatteryPack
	for _, file := range file {
		var batteryPack BatteryPack
		for _, batteryString := range file {
			battery, err := strconv.Atoi(string(batteryString))
			if err != nil {
				log.Fatal(err)
			}
			batteryPack = append(batteryPack, battery)
		}

		batteryPacks = append(batteryPacks, batteryPack)
	}

	return batteryPacks
}
