package grid

import (
	"fmt"

	"aoc/utilities"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

const (
	EMPTY       = '.'
	PAPER_ROLL  = '@'
	MOVED_UP    = '^'
	MOVED_DOWN  = 'v'
	MOVED_LEFT  = '<'
	MOVED_RIGHT = '>'
	START       = 'S'
)

type Point struct {
	X int
	Y int
}

var (
	RIGHT      = Point{X: 1, Y: 0}
	LEFT       = Point{X: -1, Y: 0}
	DOWN       = Point{X: 0, Y: 1}
	UP         = Point{X: 0, Y: -1}
	RIGHT_DOWN = Point{X: 1, Y: 1}
	RIGHT_UP   = Point{X: 1, Y: -1}
	LEFT_DOWN  = Point{X: -1, Y: 1}
	LEFT_UP    = Point{X: -1, Y: -1}
)

type Visit struct {
	Point     Point
	Direction Point
}

type Grid struct {
	Width   int
	Height  int
	Start   Point
	Current Point
	Visited []Visit
	Map     [][]rune
}

func New(filename string) *Grid {
	grid := makeGrid(filename)
	return grid
}

func (g *Grid) SetStart(point Point) {
	g.Start = point
	g.Current = point
	g.Visited = []Visit{{
		Point:     point,
		Direction: Point{X: 0, Y: 0},
	}}
}

func (g *Grid) Print() {
	for _, row := range g.Map {
		for _, char := range row {
			print(string(char))
		}
		print("\n")
	}
}

func directionToArrow(direction Point) rune {
	// Check for starting position (zero direction)
	if direction.X == 0 && direction.Y == 0 {
		return START
	}

	switch direction {
	case UP:
		return MOVED_UP
	case DOWN:
		return MOVED_DOWN
	case LEFT:
		return MOVED_LEFT
	case RIGHT:
		return MOVED_RIGHT
	default:
		return 'X'
	}
}

func (g *Grid) PrintVisited() {
	gridCopy := make([][]rune, len(g.Map))
	visitMap := make(map[Point]bool)

	for i := range g.Map {
		gridCopy[i] = make([]rune, len(g.Map[i]))
		copy(gridCopy[i], g.Map[i])
	}

	for _, visit := range g.Visited {
		gridCopy[visit.Point.Y][visit.Point.X] = directionToArrow(visit.Direction)
		visitMap[visit.Point] = true
	}

	for y, row := range gridCopy {
		for x, char := range row {
			if visitMap[Point{X: x, Y: y}] {
				if char == START {
					fmt.Print(ColorGreen + ColorBold + string(char) + ColorReset)
				} else {
					fmt.Print(ColorCyan + ColorBold + string(char) + ColorReset)
				}
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func (g *Grid) PrintVisitedSteps() {
	for step := 0; step <= len(g.Visited); step++ {
		fmt.Printf("\n%s=== Step %d ===%s\n", ColorYellow+ColorBold, step, ColorReset)

		gridCopy := make([][]rune, len(g.Map))
		visitMap := make(map[Point]bool)

		for i := range g.Map {
			gridCopy[i] = make([]rune, len(g.Map[i]))
			copy(gridCopy[i], g.Map[i])
		}

		for i := 0; i < step && i < len(g.Visited); i++ {
			visit := g.Visited[i]
			gridCopy[visit.Point.Y][visit.Point.X] = directionToArrow(visit.Direction)
			visitMap[visit.Point] = true
		}

		for y, row := range gridCopy {
			for x, char := range row {
				if visitMap[Point{X: x, Y: y}] {
					if char == START {
						fmt.Print(ColorGreen + ColorBold + string(char) + ColorReset)
					} else {
						fmt.Print(ColorCyan + ColorBold + string(char) + ColorReset)
					}
				} else {
					fmt.Print(string(char))
				}
			}
			fmt.Println()
		}
	}
}

func (g *Grid) PositionContainsObject(point Point, object rune) bool {
	if g.PointWithinBounds(point) == false {
		return false
	}
	return g.Map[point.Y][point.X] == object
}

func (g *Grid) PointWithinBounds(point Point) bool {
	return point.X < 0 || point.X >= g.Width || point.Y < 0 || point.Y >= g.Height
}

func (g *Grid) CanMoveRight() bool {
	return g.Current.X < g.Width-1
}

func (g *Grid) CanMoveLeft() bool {
	return g.Current.X > 0
}

func (g *Grid) CanMoveDown() bool {
	return g.Current.Y < g.Height-1
}

func (g *Grid) CanMoveUp() bool {
	return g.Current.Y > 0
}

func (g *Grid) CanMove(direction Point) bool {
	switch direction {
	case RIGHT:
		return g.CanMoveRight()
	case LEFT:
		return g.CanMoveLeft()
	case DOWN:
		return g.CanMoveDown()
	case UP:
		return g.CanMoveUp()
	default:
		canMoveX := true
		canMoveY := true

		if direction.X > 0 {
			canMoveX = g.CanMoveRight()
		} else if direction.X < 0 {
			canMoveX = g.CanMoveLeft()
		}

		if direction.Y > 0 {
			canMoveY = g.CanMoveDown()
		} else if direction.Y < 0 {
			canMoveY = g.CanMoveUp()
		}

		return canMoveX && canMoveY
	}
}

func (g *Grid) Move(direction Point) bool {
	canMove := g.CanMove(direction)

	if !canMove {
		return false
	}

	point := Point{
		X: g.Current.X + direction.X,
		Y: g.Current.Y + direction.Y,
	}

	visit := Visit{
		Point:     point,
		Direction: direction,
	}

	g.Visited = append(g.Visited, visit)
	g.Current = point

	return canMove
}

func makeGrid(filename string) *Grid {
	file := utilities.ReadFile(filename)
	gridMap := make([][]rune, len(file))

	for i, row := range file {
		gridMap[i] = make([]rune, len(row))
		for j, char := range row {
			gridMap[i][j] = char
		}
	}

	return &Grid{
		Width:  len(file[0]),
		Height: len(file),
		Map:    gridMap,
	}
}
