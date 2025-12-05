package challenges

import (
	"log/slog"

	. "aoc/grid"

	logger "github.com/Bparsons0904/goLogger"
)

func Day4() {
	log := logger.New("Day4")
	grid := New("day4.part1")

	// timer := log.Timer("Both Part Timers")
	// part1Count, part2Count := calculatePaperRolls(grid)
	// timer()
	//
	// slog.Info(
	// 	"Day 4",
	// 	"part1",
	// 	part1Count,
	// 	"part2",
	// 	part2Count,
	// )

	// timer := log.Timer("Both Part Timers Optimized")
	// part1Count, part2Count := calculatePaperRollsOptimized(grid)
	// timer()
	//
	// slog.Info(
	// 	"Day 4 Optimized",
	// 	"part1",
	// 	part1Count,
	// 	"part2",
	// 	part2Count,
	// )

	timer := log.Timer("Both Part Timers Queue")
	part1CountQueue, part2CountQueue := calculatePaperRollsQueue(grid)
	timer()

	slog.Info(
		"Day 4 Queue",
		"part1",
		part1CountQueue,
		"part2",
		part2CountQueue,
	)
}

func calculatePaperRollsQueue(grid *Grid) (int, int) {
	part1Count := 0
	part2Count := 0

	stack := make([]Point, 0, (grid.Height*grid.Width)/2)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.Map[y][x] == PAPER_ROLL {
				point := Point{X: x, Y: y}
				connectedRolls := countPaperRollContacts(grid, point)
				if connectedRolls < 4 {
					part1Count++
					stack = append(stack, point)
				}
			}
		}
	}

	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if grid.Map[point.Y][point.X] != PAPER_ROLL {
			continue
		}

		grid.SetObject(point, EMPTY)
		part2Count++

		for _, dir := range DIRECTIONS {
			neighbor := movePoint(point, dir)
			if grid.PositionContainsObject(neighbor, PAPER_ROLL) {
				connectedRolls := countPaperRollContacts(grid, neighbor)
				if connectedRolls < 4 {
					stack = append(stack, neighbor)
				}
			}
		}
	}

	return part1Count, part2Count
}

func calculatePaperRollsOptimized(grid *Grid) (int, int) {
	part1Count := 0
	part2Count := 0

	mappedPaperRolls := make(map[Point]bool)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.Map[y][x] == PAPER_ROLL {
				mappedPaperRolls[Point{X: x, Y: y}] = true
				connectedRolls := countPaperRollContacts(grid, Point{X: x, Y: y})
				if connectedRolls < 4 {
					part1Count++
				}
			}
		}
	}

	lastPassCount := -1
	for {
		if lastPassCount == 0 {
			break
		}
		lastPassCount = 0

		for point := range mappedPaperRolls {
			connectedRolls := countPaperRollContacts(grid, point)
			if connectedRolls < 4 {
				grid.SetObject(point, EMPTY)
				delete(mappedPaperRolls, point)
				lastPassCount++
				part2Count++
			}
		}
	}

	return part1Count, part2Count
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

	for _, direction := range DIRECTIONS {
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
