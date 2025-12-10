package challenges

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"aoc/grid"
	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

func Day9() {
	log := logger.New("Day9")

	filename := "day9.part1"
	redTiles := getRedTiles(filename)

	timer := log.Timer("Part 1 Timer")
	largestAreaPart1 := getLargestArea(redTiles)
	timer()

	timer = log.Timer("Part 2 Timer getRedAndGreenTiles")
	redAndGreenTiles := getRedAndGreenTiles(redTiles)
	timer()

	timer = log.Timer("Part 2 Timer")
	largestAreaPart2 := getLargestAreaWithRedOrGreen(redTiles, redAndGreenTiles)
	timer()

	log.Info("Results", "Part 1", largestAreaPart1, "Part 2", largestAreaPart2)
}

func getRedAndGreenTiles(redTiles []grid.Point) *grid.Grid {
	log := logger.New("Day9").With("Function", "getRedAndGreenTiles")
	timer := log.Timer("make tile grid")
	tileGrid := grid.MakeGridByPoints(
		redTiles,
		grid.TileDef{Char: '#', Color: grid.ColorRed, Points: redTiles},
		grid.TileDef{Char: 'O', Color: grid.ColorGreen},
	)
	timer()

	timer = log.Timer("set red and green tiles")
	for _, point := range redTiles {
		farthestRight := tileGrid.FindLastObjectToRight(point, '#')
		if farthestRight != (grid.Point{}) {
			for i := point.X; i <= farthestRight.X; i++ {
				tileGrid.SetObject(grid.Point{X: i, Y: point.Y}, 'O')
			}
		}
		farthestDown := tileGrid.FindLastObjectToBottom(point, '#')
		if farthestDown != (grid.Point{}) {
			for i := point.Y; i <= farthestDown.Y; i++ {
				tileGrid.SetObject(grid.Point{X: point.X, Y: i}, 'O')
			}
		}
	}
	timer()

	timer = log.Timer("set outer tiles")
OuterLoop:
	for i, point := range tileGrid.Map {
		var farthestRight grid.Point

		for j, char := range point {
			if char == 'O' && farthestRight == (grid.Point{}) {
				farthestRight = tileGrid.FindLastObjectToRight(grid.Point{X: j, Y: i}, 'O')
			}
			if farthestRight != (grid.Point{}) {
				if j == farthestRight.X {
					continue OuterLoop
				}
				tileGrid.SetObject(grid.Point{X: j, Y: i}, 'O')
			}
		}
	}
	tileGrid.Print()

	timer()

	return tileGrid
}

func getLargestAreaWithRedOrGreen(redTiles []grid.Point, fullGrid *grid.Grid) int {
	log := logger.New("Day9").With("Function", "getLargestAreaWithRedOrGreen")
	timer := log.Timer("get potential tiles")
	type potentialTile struct {
		point1 grid.Point
		point2 grid.Point
		result int
	}
	potentialTiles := make([]potentialTile, 0, len(redTiles)*len(redTiles))

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for _, point1 := range redTiles {
		for _, point2 := range redTiles {
			width := abs(point1.X-point2.X) + 1
			height := abs(point1.Y-point2.Y) + 1
			area := width * height
			potentialTiles = append(
				potentialTiles,
				potentialTile{point1: point1, point2: point2, result: area},
			)
		}
	}

	timer()

	timer = log.Timer("sort potential tiles")
	slices.SortFunc(potentialTiles, func(a, b potentialTile) int {
		return int(b.result - a.result)
	})

	timer()

OuterLoop:
	for _, potential := range potentialTiles {
		// timer = log.Timer("check potential tile")
		minX := potential.point1.X
		maxX := potential.point2.X
		if potential.point2.X < minX {
			minX, maxX = maxX, minX
		}
		minY := potential.point1.Y
		maxY := potential.point2.Y
		if potential.point2.Y < minY {
			minY, maxY = maxY, minY
		}

		for i := minY; i <= maxY; i++ {
			for j := minX; j <= maxX; j++ {
				if fullGrid.Map[i][j] != 'O' {
					// timer()
					continue OuterLoop
				}
			}
		}

		// timer()
		return int(potential.result)
	}

	return 0
}

func getLargestArea(redTiles []grid.Point) int {
	result := 0.0

	for _, point1 := range redTiles {
		for _, point2 := range redTiles {
			width := math.Abs(float64(point1.X)-float64(point2.X)) + 1
			height := math.Abs(float64(point1.Y)-float64(point2.Y)) + 1
			area := width * height
			if area > result {
				result = area
			}
		}
	}
	return int(result)
}

func getRedTiles(filename string) []grid.Point {
	file := utilities.ReadFile(filename)

	var redTiles []grid.Point
	for _, row := range file {
		coordinates := strings.Split(row, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		redTiles = append(redTiles, grid.Point{X: x, Y: y})
	}

	return redTiles
}
