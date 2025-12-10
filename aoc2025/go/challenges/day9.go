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

	timer = log.Timer("Part 2 Original Timer")
	largestAreaPart2 := getLargestAreaWithIntervals(redTiles, intervals)
	timer()

	timer = log.Timer("Part 2 Optimized Timer")
	largestAreaPart2Optimized := getLargestAreaWithIntervalsOptimized(redTiles, intervals)
	timer()

	log.Info(
		"Results",
		"Part 1",
		largestAreaPart1,
		"Part 2",
		largestAreaPart2,
		"Part 2 Optimized",
		largestAreaPart2Optimized,
	)
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

	timer := log.Timer("precompute max widths")
	maxWidthAtY := make(map[int]int)
	for y, ivs := range intervals {
		maxW := 0
		for _, iv := range ivs {
			w := iv.End - iv.Start + 1
			if w > maxW {
				maxW = w
			}
		}
		maxWidthAtY[y] = maxW
	}
	timer()

	timer = log.Timer("get potential tiles")
	type candidate struct {
		minX, maxX, minY, maxY int
		area                   int
	}

	candidates := make([]candidate, 0, len(redTiles)*(len(redTiles)-1)/2)

	for i := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			p1, p2 := redTiles[i], redTiles[j]
			minX, maxX := p1.X, p2.X
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			minY, maxY := p1.Y, p2.Y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			width := maxX - minX + 1
			height := maxY - minY + 1
			candidates = append(candidates, candidate{
				minX: minX, maxX: maxX, minY: minY, maxY: maxY,
				area: width * height,
			})
		}
	}
	timer()

	timer = log.Timer("sort potential tiles")
	slices.SortFunc(candidates, func(a, b candidate) int {
		return b.area - a.area
	})
	timer()

	timer = log.Timer("check rectangles")
	for _, c := range candidates {
		width := c.maxX - c.minX + 1

		canFit := true
		for y := c.minY; y <= c.maxY; y++ {
			if maxWidthAtY[y] < width {
				canFit = false
				break
			}
		}
		if !canFit {
			continue
		}

		if isRectangleInside(c.minX, c.maxX, c.minY, c.maxY, intervals) {
			timer()
			return c.area
		}
	}
	timer()

	return 0
}

// Gemini algorithm
// Segment Tree for Range Minimum Query
type SegmentTree struct {
	tree    []int
	yCoords []int
	yIndex  map[int]int
	size    int
}

// Gemini algorithm
func buildSegmentTree(maxWidthAtY map[int]int) *SegmentTree {
	yCoords := make([]int, 0, len(maxWidthAtY))
	for y := range maxWidthAtY {
		yCoords = append(yCoords, y)
	}
	slices.Sort(yCoords)

	yIndex := make(map[int]int, len(yCoords))
	for i, y := range yCoords {
		yIndex[y] = i
	}

	size := len(yCoords)
	tree := make([]int, 4*size)

	var build func(arrIndex, start, end int)
	build = func(arrIndex, start, end int) {
		if start == end {
			tree[arrIndex] = maxWidthAtY[yCoords[start]]
			return
		}
		mid := (start + end) / 2
		build(2*arrIndex+1, start, mid)
		build(2*arrIndex+2, mid+1, end)
		tree[arrIndex] = min(tree[2*arrIndex+1], tree[2*arrIndex+2])
	}

	if size > 0 {
		build(0, 0, size-1)
	}

	return &SegmentTree{tree: tree, yCoords: yCoords, yIndex: yIndex, size: size}
}

// Gemini algorithm
func (st *SegmentTree) Query(minY, maxY int) int {
	if st.size == 0 {
		return 0
	}

	l, okL := st.yIndex[minY]
	r, okR := st.yIndex[maxY]

	// If the exact Y is not a key, find the next available one.
	if !okL {
		// Find where minY would be inserted
		idx, _ := slices.BinarySearch(st.yCoords, minY)
		if idx >= len(st.yCoords) {
			return 0
		}
		l = idx
	}
	if !okR {
		idx, _ := slices.BinarySearch(st.yCoords, maxY)
		if idx == 0 {
			return 0
		}
		r = idx - 1
	}

	if l > r {
		// This can happen if the range is empty or falls between our yCoords.
		// A single row check might need this.
		if _, ok := st.yIndex[minY]; ok && minY == maxY {
			r = l
		} else {
			return 0
		}
	}

	var query func(arrIndex, start, end, qStart, qEnd int) int
	query = func(arrIndex, start, end, qStart, qEnd int) int {
		if qStart > end || qEnd < start {
			return math.MaxInt
		}
		if qStart <= start && qEnd >= end {
			return st.tree[arrIndex]
		}
		mid := (start + end) / 2
		leftQuery := query(2*arrIndex+1, start, mid, qStart, qEnd)
		rightQuery := query(2*arrIndex+2, mid+1, end, qStart, qEnd)
		return min(leftQuery, rightQuery)
	}

	return query(0, 0, st.size-1, l, r)
}

// Gemini algorithm
func getLargestAreaWithIntervalsOptimized(redTiles []grid.Point, intervals map[int][]Interval) int {
	log := logger.New("Day9").With("Function", "getLargestAreaWithIntervalsOptimized")

	timer := log.Timer("precompute max widths")
	maxWidthAtY := make(map[int]int)
	for y, ivs := range intervals {
		maxW := 0
		for _, iv := range ivs {
			w := iv.End - iv.Start + 1
			if w > maxW {
				maxW = w
			}
		}
		maxWidthAtY[y] = maxW
	}
	timer()

	timer = log.Timer("build segment tree")
	segTree := buildSegmentTree(maxWidthAtY)
	timer()

	timer = log.Timer("get potential tiles")
	type candidate struct {
		minX, maxX, minY, maxY int
		area                   int
	}

	candidates := make([]candidate, 0, len(redTiles)*(len(redTiles)-1)/2)

	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			p1, p2 := redTiles[i], redTiles[j]
			minX, maxX := p1.X, p2.X
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			minY, maxY := p1.Y, p2.Y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			width := maxX - minX + 1
			height := maxY - minY + 1
			// Optimization: if a single row can't fit width, no point generating candidate
			if _, ok := maxWidthAtY[p1.Y]; ok && maxWidthAtY[p1.Y] < width {
				continue
			}
			if _, ok := maxWidthAtY[p2.Y]; ok && maxWidthAtY[p2.Y] < width {
				continue
			}

			candidates = append(candidates, candidate{
				minX: minX, maxX: maxX, minY: minY, maxY: maxY,
				area: width * height,
			})
		}
	}
	timer()

	timer = log.Timer("sort potential tiles")
	slices.SortFunc(candidates, func(a, b candidate) int {
		return b.area - a.area
	})
	timer()

	timer = log.Timer("check rectangles")
	for _, c := range candidates {
		width := c.maxX - c.minX + 1

		minWidthInRange := segTree.Query(c.minY, c.maxY)

		if minWidthInRange < width {
			continue
		}

		if isRectangleInside(c.minX, c.maxX, c.minY, c.maxY, intervals) {
			timer()
			return c.area
		}
	}
	timer()

	return 0
}

func isRectangleInside(minX, maxX, minY, maxY int, intervals map[int][]Interval) bool {
	for y := minY; y <= maxY; y++ {
		rowIntervals, exists := intervals[y]
		if !exists {
			return false
		}

		// Binary search: find the rightmost interval with Start <= minX
		lo, hi := 0, len(rowIntervals)
		for lo < hi {
			mid := (lo + hi) / 2
			if rowIntervals[mid].Start <= minX {
				lo = mid + 1
			} else {
				hi = mid
			}
		}

		// Check if the interval before 'lo' contains our range
		if lo == 0 || rowIntervals[lo-1].End < maxX {
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
