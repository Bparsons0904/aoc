package challenges

import (
	"log/slog"

	. "aoc/grid"
)

func Day4() {
	slog.Info("day4")
	grid := New("day4.part0")

	grid.Print()

	grid.SetStart(Point{X: 1, Y: 1})
	grid.Move(RIGHT)
	grid.Move(DOWN)
	grid.Move(DOWN)
	grid.Move(DOWN)
	grid.PrintVisitedSteps()
}
