package challenges

import (
	"fmt"
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
	part1Result := calculate2PackJoltage(batteryPacks)
	timer()

	timer = log.Timer("Part 2 Timer")
	part2Result := calculate12PackJoltage(batteryPacks)
	timer()

	log.Info("Results", "Part 1", part1Result, "Part 2", part2Result)
}

func calculate12PackJoltage(batteryPacks []BatteryPack) int {
	maxJoltage := 0
	for _, batteryPack := range batteryPacks {
		maxJoltage += batteryPack.getLargest12PackJoltage()
	}

	return maxJoltage
}

func (b BatteryPack) getLargest12PackJoltage() int {
	maxJoltage := make([]int, 0, 12)

	index := 0
	iteration := 0
	for {
		if len(maxJoltage) == 12 {
			break
		}
		maxFoundJoltage, endIndex := iterateBatteryPackCheck(b, index, iteration)
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

func iterateBatteryPackCheck(battery BatteryPack, startingIndex int, batterySize int) (int, int) {
	maxFoundJoltage := 0
	endIndex := startingIndex
	stop := len(battery) - 11 + batterySize
	for i := startingIndex; i < stop; i++ {
		if battery[i] > maxFoundJoltage {
			maxFoundJoltage = battery[i]
			endIndex = i
		}
	}

	return maxFoundJoltage, endIndex
}

func calculate2PackJoltage(batteryPacks []BatteryPack) int {
	maxJoltage := 0
	for _, batteryPack := range batteryPacks {
		maxJoltage += batteryPack.getLargest2PackJoltage()
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
