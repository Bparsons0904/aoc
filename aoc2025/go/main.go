package main

import (
	"time"

	"aoc/challenges"
)

func main() {
	switch time.Now().Day() {
	case 1:
		challenges.Day1()
		challenges.Day1_1()
	case 2:
		challenges.Day2()
	case 3:
		challenges.Day3()
	case 4:
		// challenges.Day4()
		challenges.Day5()
	case 5:
		challenges.Day5()
	case 6:
		challenges.Day6()
	case 7:
		challenges.Day7()
	case 8:
		challenges.Day8()
	case 9:
		challenges.Day9()
	case 10:
		challenges.Day10()
	case 11:
		challenges.Day11()
	case 12:
		challenges.Day12()
	}
}
