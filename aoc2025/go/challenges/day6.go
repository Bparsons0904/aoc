package challenges

import (
	"log"
	"log/slog"
	"strconv"
	"strings"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

func Day6() {
	log := logger.New("Day6")
	worksheet := processDay6File()

	timer := log.Timer("Timer")
	part1Total, part2Total := calculateWorksheets(worksheet)
	timer()

	slog.Info("part1", "Part 1", part1Total, "Part 2", part2Total)
}

func calculateWorksheets(worksheet WorksheetMap) (int, int) {
	total := 0
	for _, ws := range worksheet {
		subtotal := 0
		for _, value := range ws.values {
			switch ws.operator {
			case "+":
				subtotal += value
			case "*":
				if subtotal == 0 {
					subtotal = 1
				}
				subtotal *= value
			}
		}
		total += subtotal
	}

	cephalopodTotal := 0
	for _, ws := range worksheet {
		cephalopodSubtotal := 0
		for _, cephalopodValue := range ws.cephalopodValues {
			switch ws.operator {
			case "+":
				cephalopodSubtotal += cephalopodValue

			case "*":
				if cephalopodSubtotal == 0 {
					cephalopodSubtotal = 1
				}
				cephalopodSubtotal *= cephalopodValue
			}
		}
		cephalopodTotal += cephalopodSubtotal
	}

	return total, cephalopodTotal
}

type Worksheet struct {
	values           []int
	cephalopodValues []int
	operator         string
}

type WorksheetMap map[int]Worksheet

func processDay6File() WorksheetMap {
	file := utilities.ReadFile("day6.part1")

	workSheetMap := buildOperators(file)
	file = file[:len(file)-1]
	workSheetMap.processValues(file)
	workSheetMap.processCephalopodValues(file)

	return workSheetMap
}

func (ws WorksheetMap) processCephalopodValues(file []string) {
	tempArray := make([][]string, 0)
	for _, row := range file {
		tempArray = append(tempArray, strings.Split(row, ""))
	}
	length := len(tempArray[0])
	mapIndex := len(ws) - 1
	for j := length - 1; j >= 0; j-- {
		var sb strings.Builder
		for i := 0; i < len(tempArray); i++ {
			sb.WriteString(tempArray[i][j])
		}

		value := strings.TrimSpace(sb.String())
		if value == "" {
			mapIndex--
			continue
		}

		tempWs := ws[mapIndex]
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		tempWs.cephalopodValues = append(tempWs.cephalopodValues, valueInt)
		ws[mapIndex] = tempWs
	}
}

func (ws WorksheetMap) processValues(file []string) {
	for _, row := range file {
		values := strings.Fields(row)
		for j, valueString := range values {
			valueInt, err := strconv.Atoi(valueString)
			if err != nil {
				log.Fatal(err)
			}
			tempWs := ws[j]
			tempWs.values = append(tempWs.values, valueInt)
			ws[j] = tempWs
		}
	}
}

func buildOperators(file []string) WorksheetMap {
	workSheetMap := make(WorksheetMap)

	operators := strings.Fields(file[len(file)-1:][0])
	for i, operatorString := range operators {
		workSheetMap[i] = Worksheet{
			values:   make([]int, 0, len(file)),
			operator: operatorString,
		}
	}

	return workSheetMap
}
