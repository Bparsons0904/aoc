package challenges

import (
	"log/slog"

	. "aoc/grid"

	logger "github.com/Bparsons0904/goLogger"
)

func Day4() {
	log := logger.New("Day4")
	grid := New("day4.part1")

	timer := log.Timer("Both Part Timers")
	part1Count, part2Count := calculatePaperRolls(grid)
	timer()

	slog.Info(
		"Day 4",
		"part1",
		part1Count,
		"part2",
		part2Count,
	)
}

func calculatePaperRolls(grid *Grid) (int, int) {
	part1Count := 0
	part2Count := 0

	firstPass := true
	lastPassCount := -1
	for {
		if lastPassCount == 0 {
			break
		}
		lastPassCount = 0

		for y := 0; y < grid.Height; y++ {
			for x := 0; x < grid.Width; x++ {
				if grid.Map[y][x] != PAPER_ROLL {
					continue
				}
				connectedRolls := countPaperRollContacts(grid, Point{X: x, Y: y})
				if connectedRolls < 4 {
					if firstPass {
						part1Count++
					} else {
						grid.SetObject(Point{X: x, Y: y}, EMPTY)
						lastPassCount++
						part2Count++
					}
				}
			}
		}

		if firstPass {
			firstPass = false
			lastPassCount = -1
		}
	}

	return part1Count, part2Count
}

func countPaperRollContacts(grid *Grid, point Point) int {
	count := 0
	pointsToCheck := []Point{
		LEFT, RIGHT, UP, DOWN, LEFT_DOWN, LEFT_UP, RIGHT_DOWN, RIGHT_UP,
	}

	for _, direction := range pointsToCheck {
		if grid.PositionContainsObject(movePoint(point, direction), PAPER_ROLL) {
			count++
		}
	}

	return count
}

func movePoint(point Point, direction Point) Point {
	return Point{
		X: point.X + direction.X,
		Y: point.Y + direction.Y,
	}
}
