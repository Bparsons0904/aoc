package challenges

import (
	"fmt"
	"log"
	"log/slog"
	"strconv"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

type BatteryPack []int

func Day3() {
	log := logger.New("Day3")
	slog.Info("day3")

	batteryPacks := processDay3File()

	part1Result := calculate2PackJoltage(batteryPacks)
	log.Info("Results", "Part 1", part1Result)
}

func calculate2PackJoltage(batteryPacks []BatteryPack) int {
	maxJoltage := 0
	for i, batteryPack := range batteryPacks {
		maxJoltage += batteryPack.getLargest2PackJoltage()
		slog.Info("Battery Pack", "Index", i, "Joltage", batteryPack.getLargest2PackJoltage())
	}

	return maxJoltage
}

func (b BatteryPack) getLargest2PackJoltage() int {
	maxJoltageFirst := 0
	maxJoltageSecond := 0
	index := 0

	for i := 0; i < len(b)-1; i++ {
		if b[i] > maxJoltageFirst {
			maxJoltageFirst = b[i]
			index = i
		}
	}

	for i := index + 1; i < len(b); i++ {
		if b[i] > maxJoltageSecond {
			maxJoltageSecond = b[i]
		}
	}

	stringedInt := fmt.Sprintf("%d%d", maxJoltageFirst, maxJoltageSecond)
	maxJoltage, err := strconv.Atoi(stringedInt)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info(
		"Max Joltage",
		"Index",
		index,
		"Joltage First",
		maxJoltageFirst,
		"Joltage Second",
		maxJoltageSecond,
	)

	return maxJoltage
}

func processDay3File() []BatteryPack {
	file := utilities.ReadFile("day3.part1")
	var batteryPacks []BatteryPack
	for _, file := range file {
		batteryPacks = append(batteryPacks, getBatteryPack(file))
	}

	return batteryPacks
}

func getBatteryPack(batteryPackString string) BatteryPack {
	var batteryPack BatteryPack
	for _, batteryString := range batteryPackString {
		battery, err := strconv.Atoi(string(batteryString))
		if err != nil {
			log.Fatal(err)
		}
		batteryPack = append(batteryPack, battery)
	}

	return batteryPack
}
