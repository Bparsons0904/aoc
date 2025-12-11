package challenges

import (
	"log"
	"strconv"
	"strings"
	"unicode/utf8"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

type Indicator string

const (
	Off Indicator = "."
	On  Indicator = "#"
)

type MachineInstruction struct {
	Indicator           Indicator
	ToggleButtons       [][]int
	JoltageRequirements []int
}

type MachineInstructions struct {
	MachineInstructions []MachineInstruction
}

func Day10() {
	log := logger.New("Day10")

	var machineInstructions MachineInstructions
	machineInstructions.processMachineInstructions("day10.part1")

	timer := log.Timer("Part 1 Timer")
	part1Result := getPart1MachineInitializationResults(machineInstructions)
	timer()

	log.Info("day10", "part1", part1Result, "part2", 0)
}

func getPart1MachineInitializationResults(machineInstructions MachineInstructions) int {
	result := 0
	for _, mi := range machineInstructions.MachineInstructions {
		sb := strings.Builder{}
		indicatorRuneCount := utf8.RuneCountInString(string(mi.Indicator))
		for range indicatorRuneCount {
			sb.WriteString(".")
		}

		baseIndicator := Indicator(sb.String())
		result += getLowestPresses(mi, baseIndicator, 0)
	}
	return result
}

func getLowestPresses(mi MachineInstruction, indicator Indicator, iterations int) int {
	presses := 1
	indicators := []Indicator{indicator}

	for {
		newIndicators := []Indicator{indicator}
		for _, indicator := range indicators {
			for _, buttons := range mi.ToggleButtons {
				newIndicator, found := clickAndCheckIfInitialized(mi, buttons, indicator)
				if found {
					return presses
				}
				newIndicators = append(newIndicators, newIndicator)
			}
		}

		presses++
		indicators = newIndicators
	}
}

func clickAndCheckIfInitialized(
	mi MachineInstruction,
	button []int,
	currentState Indicator,
) (Indicator, bool) {
	newState := []byte(currentState)
	for _, buttonIndex := range button {
		if newState[buttonIndex] == Off[0] {
			newState[buttonIndex] = On[0]
		} else {
			newState[buttonIndex] = Off[0]
		}
	}

	newStateStr := string(newState)
	if newStateStr == string(mi.Indicator) {
		return Indicator(newStateStr), true
	}
	return Indicator(newStateStr), false
}

func (mis *MachineInstructions) processMachineInstructions(filename string) {
	file := utilities.ReadFile(filename)

	for _, row := range file {
		var mi MachineInstruction
		instructions := strings.Split(row, " ")
		// Strip the square brackets from the indicator
		indicatorStr := instructions[0]
		mi.Indicator = Indicator(indicatorStr[1 : len(indicatorStr)-1])
		mi.calculateJoltageRequirements(instructions[len(instructions)-1])
		mi.calculateToggleButtons(instructions[1 : len(instructions)-1])

		mis.MachineInstructions = append(mis.MachineInstructions, mi)
	}
}

func (mi *MachineInstruction) calculateToggleButtons(toggleButtonsString []string) {
	for _, toggleButtonGroup := range toggleButtonsString {
		toggleButtonsIndexes := toggleButtonGroup[1 : len(toggleButtonGroup)-1]
		buttons := strings.Split(toggleButtonsIndexes, ",")
		var toggleButtons []int
		for _, button := range buttons {
			buttonInt, err := strconv.Atoi(button)
			if err != nil {
				log.Fatal(err)
			}
			toggleButtons = append(toggleButtons, buttonInt)
		}

		mi.ToggleButtons = append(mi.ToggleButtons, toggleButtons)
	}
}

func (mi *MachineInstruction) calculateJoltageRequirements(joltageRequirementsString string) {
	joltageRequirementsInt := joltageRequirementsString[1 : len(joltageRequirementsString)-1]

	joltages := strings.Split(joltageRequirementsInt, ",")
	var joltageRequirements []int
	for _, joltage := range joltages {
		joltageInt, err := strconv.Atoi(joltage)
		if err != nil {
			log.Fatal(err)
		}
		joltageRequirements = append(joltageRequirements, joltageInt)
	}

	mi.JoltageRequirements = joltageRequirements
}
