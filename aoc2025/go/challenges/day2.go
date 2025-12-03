package challenges

import (
	"log"
	"math"
	"strconv"
	"strings"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

type ProductIDRange struct {
	Min int
	Max int
}

func Day2() {
	log := logger.New("Day2")

	productIDRanges := processDay2()

	part1Timer := log.Timer("Part 1 Timer")
	part1Count := calculatePart1(productIDRanges)
	part1Timer()

	part2Timer := log.Timer("Part 2 Timer")
	part2Count := calculatePart2(productIDRanges)
	part2Timer()

	part2_1Timer := log.Timer("Part 2.1 Timer")
	part2_1Count := calculatePart2_1(productIDRanges)
	part2_1Timer()

	log.Info("Day 2", "part1", part1Count, "part2", part2Count, "part2_1", part2_1Count)
}

func calculatePart2_1(productIDRanges []ProductIDRange) (result int) {
	for _, productIDRange := range productIDRanges {
		for value := productIDRange.Min; value <= productIDRange.Max; value++ {
			iString := strconv.Itoa(value)
			for j := 0; j <= len(iString)/2; j++ {
				if strings.ReplaceAll(iString[j+1:], string(iString[:j+1]), "") == "" {
					result += value
					break
				}
			}
		}
	}

	return
}

func calculatePart2(productIDRanges []ProductIDRange) int {
	result := 0

	for _, productIDRange := range productIDRanges {
		for i := productIDRange.Min; i <= productIDRange.Max; i++ {
			iString := strconv.Itoa(i)
			for j := 0; j <= len(iString)-2; j++ {
				toCheck := string(iString[:j+1])
				expectedCount := int(math.Ceil(float64(len(iString)) / float64(len(toCheck))))
				count := strings.Count(iString, toCheck)
				if count >= 2 && count == expectedCount {
					result += i
					break
				}

			}
		}
	}

	return result
}

func calculatePart1(productIDRanges []ProductIDRange) int {
	result := 0
	for _, productIDRange := range productIDRanges {
		for i := productIDRange.Min; i <= productIDRange.Max; i++ {
			if len(strconv.Itoa(i))%2 != 0 {
				continue
			}

			half := len(strconv.Itoa(i)) / 2
			if strconv.Itoa(i)[:half] == strconv.Itoa(i)[half:] {
				result += i
			}
		}
	}

	return result
}

func processDay2() []ProductIDRange {
	row := utilities.ReadFile("day2.part1")[0]

	ranges := strings.Split(row, ",")
	var productIDRanges []ProductIDRange
	for _, value := range ranges {
		idRange := strings.Split(value, "-")
		min, err := strconv.Atoi(idRange[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(idRange[1])
		if err != nil {
			log.Fatal(err)
		}

		productIDRanges = append(productIDRanges, ProductIDRange{
			Min: min,
			Max: max,
		})
	}

	return productIDRanges
}
