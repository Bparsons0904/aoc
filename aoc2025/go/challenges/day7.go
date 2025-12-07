package challenges

import (
	"log/slog"

	"aoc/grid"
)

func Day7() {
	slog.Info("day7")

	tachyonGrid := grid.New("day7.part1")
	trimGrid(tachyonGrid)

	part1Count := processTachyonBeamSplitCounter(tachyonGrid)
	part2Count := processTachyonBeamRoutesCounter(tachyonGrid)

	slog.Info("part1", "Part 1", part1Count, "Part 2", part2Count)
}

func processTachyonBeamRoutesCounter(tachyonGrid *grid.Grid) int {
	return processPath(tachyonGrid.Current, tachyonGrid)
}

func processPath(point grid.Point, tachyonGrid *grid.Grid) int {
	if point.Y == len(tachyonGrid.Map)-1 {
		slog.Info("reached end", "point", point)
		return 1
	}

	pathCount := 0
	if tachyonGrid.Map[point.Y+1][point.X] == grid.TACHYON {
		if point.X-1 >= 0 {
			slog.Info("left", "point", point)
			pathCount += processPath(grid.Point{X: point.X - 1, Y: point.Y + 1}, tachyonGrid)
		}

		if point.X+1 < len(tachyonGrid.Map[0]) {
			slog.Info("right", "point", point)
			pathCount += processPath(grid.Point{X: point.X + 1, Y: point.Y + 1}, tachyonGrid)
		}
	} else {
		slog.Info("down", "point", point)
		pathCount += processPath(grid.Point{X: point.X, Y: point.Y + 1}, tachyonGrid)
	}

	return pathCount
}

func processTachyonBeamSplitCounter(tachyonGrid *grid.Grid) int {
	tachyonSplitCounter := 0
	tachyonCurrentLines := make(map[int]bool)
	tachyonCurrentLines[tachyonGrid.Current.X] = true
	for _, row := range tachyonGrid.Map {
		newTachyonCurrentLines := make(map[int]bool)
		for x, space := range row {
			if space == grid.TACHYON && tachyonCurrentLines[x] {
				tachyonSplitCounter++
				newTachyonCurrentLines[x-1] = true
				newTachyonCurrentLines[x+1] = true
				delete(tachyonCurrentLines, x)
			}
		}

		for x := range newTachyonCurrentLines {
			if x < 0 || x >= len(tachyonGrid.Map[0]) {
				continue
			}
			tachyonCurrentLines[x] = true
		}
	}

	return tachyonSplitCounter
}

func trimGrid(tachyonGrid *grid.Grid) {
	for i := 1; i < len(tachyonGrid.Map); i++ {
		hasTachyon := false
		for j := 0; j < len(tachyonGrid.Map[i]); j++ {
			if tachyonGrid.Map[i][j] == grid.TACHYON {
				hasTachyon = true
				break
			}
		}
		if !hasTachyon {
			tachyonGrid.Map = append(tachyonGrid.Map[:i], tachyonGrid.Map[i+1:]...)
			i--
		}
	}
}
