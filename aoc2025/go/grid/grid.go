package grid

import (
	"aoc/utilities"
)

const (
	EMPTY      = '.'
	PAPER_ROLL = '@'
)

type Point struct {
	X int
	Y int
}

var (
	RIGHT      = Point{X: 0, Y: 1}
	LEFT       = Point{X: 0, Y: -1}
	DOWN       = Point{X: 1, Y: 0}
	UP         = Point{X: -1, Y: 0}
	RIGHT_DOWN = Point{X: 1, Y: 1}
	RIGHT_UP   = Point{X: -1, Y: 1}
	LEFT_DOWN  = Point{X: 1, Y: -1}
	LEFT_UP    = Point{X: -1, Y: -1}
)

type Grid struct {
	Width   int
	Height  int
	Current Point
	Visited []Point
	Map     [][]rune
}

func New(filename string) *Grid {
	grid := makeGrid(filename)
	return grid
}

func (g *Grid) Print() {
	for _, row := range g.Map {
		for _, char := range row {
			print(string(char))
		}
		print("\n")
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
	return g.Current.X < g.Width
}

func (g *Grid) CanMoveLeft() bool {
	return g.Current.X > 0
}

func (g *Grid) CanMoveDown() bool {
	return g.Current.Y < g.Height
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
			canMoveX = g.CanMoveDown()
		} else if direction.X < 0 {
			canMoveX = g.CanMoveUp()
		}

		if direction.Y > 0 {
			canMoveY = g.CanMoveRight()
		} else if direction.Y < 0 {
			canMoveY = g.CanMoveLeft()
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

	g.Visited = append(g.Visited, point)
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
