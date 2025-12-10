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

type Interval struct {
	Start, End int
}

func Day9() {
	log := logger.New("Day9")

	filename := "day9.part1"
	redTiles := getRedTiles(filename)

	timer := log.Timer("Part 1 Timer")
	largestAreaPart1 := getLargestArea(redTiles)
	timer()

	timer = log.Timer("Part 2 Timer getInsideIntervals")
	intervals := getInsideIntervals(redTiles)
	timer()

	timer = log.Timer("Part 2 Timer")
	largestAreaPart2 := getLargestAreaWithIntervals(redTiles, intervals)
	timer()

	log.Info("Results", "Part 1", largestAreaPart1, "Part 2", largestAreaPart2)
}

func getInsideIntervals(redTiles []grid.Point) map[int][]Interval {
	log := logger.New("Day9").With("Function", "getInsideIntervals")

	type VerticalEdge struct {
		X          int
		MinY, MaxY int
	}
	type HorizontalEdge struct {
		Y          int
		MinX, MaxX int
	}

	var verticalEdges []VerticalEdge
	var horizontalEdges []HorizontalEdge

	timer := log.Timer("build edges")
	for i := range redTiles {
		current := redTiles[i]
		next := redTiles[(i+1)%len(redTiles)]

		if current.X == next.X {
			minY, maxY := current.Y, next.Y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			verticalEdges = append(
				verticalEdges,
				VerticalEdge{X: current.X, MinY: minY, MaxY: maxY},
			)
		} else {
			minX, maxX := current.X, next.X
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			horizontalEdges = append(horizontalEdges, HorizontalEdge{Y: current.Y, MinX: minX, MaxX: maxX})
		}
	}
	timer()

	minY, maxY := redTiles[0].Y, redTiles[0].Y
	for _, p := range redTiles {
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	timer = log.Timer("build intervals")
	intervals := make(map[int][]Interval)

	for y := minY; y <= maxY; y++ {
		var crossings []int

		for _, edge := range verticalEdges {
			if y >= edge.MinY && y < edge.MaxY {
				crossings = append(crossings, edge.X)
			}
		}

		slices.Sort(crossings)

		var rowIntervals []Interval
		for i := 0; i+1 < len(crossings); i += 2 {
			rowIntervals = append(rowIntervals, Interval{Start: crossings[i], End: crossings[i+1]})
		}

		for _, edge := range horizontalEdges {
			if edge.Y == y {
				rowIntervals = append(rowIntervals, Interval{Start: edge.MinX, End: edge.MaxX})
			}
		}

		if len(rowIntervals) > 0 {
			intervals[y] = mergeIntervals(rowIntervals)
		}
	}
	timer()

	return intervals
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	slices.SortFunc(intervals, func(a, b Interval) int {
		return a.Start - b.Start
	})

	merged := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		curr := intervals[i]

		if curr.Start <= last.End+1 {
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}

func getLargestAreaWithIntervals(redTiles []grid.Point, intervals map[int][]Interval) int {
	log := logger.New("Day9").With("Function", "getLargestAreaWithIntervals")

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	timer := log.Timer("get potential tiles")
	type potentialTile struct {
		point1, point2 grid.Point
		area           int
	}
	potentialTiles := make([]potentialTile, 0, len(redTiles)*len(redTiles))

	for _, p1 := range redTiles {
		for _, p2 := range redTiles {
			if p1 == p2 {
				continue
			}
			width := abs(p1.X-p2.X) + 1
			height := abs(p1.Y-p2.Y) + 1
			potentialTiles = append(potentialTiles, potentialTile{
				point1: p1,
				point2: p2,
				area:   width * height,
			})
		}
	}
	timer()

	timer = log.Timer("sort potential tiles")
	slices.SortFunc(potentialTiles, func(a, b potentialTile) int {
		return b.area - a.area
	})
	timer()

	timer = log.Timer("check rectangles")
	for _, pt := range potentialTiles {
		minX, maxX := pt.point1.X, pt.point2.X
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		minY, maxY := pt.point1.Y, pt.point2.Y
		if minY > maxY {
			minY, maxY = maxY, minY
		}

		if isRectangleInside(minX, maxX, minY, maxY, intervals) {
			timer()
			return pt.area
		}
	}

	return 0
}

func isRectangleInside(minX, maxX, minY, maxY int, intervals map[int][]Interval) bool {
	for y := minY; y <= maxY; y++ {
		rowIntervals, exists := intervals[y]
		if !exists {
			return false
		}

		contained := false
		for _, iv := range rowIntervals {
			if iv.Start <= minX && maxX <= iv.End {
				contained = true
				break
			}
		}
		if !contained {
			return false
		}
	}
	return true
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
