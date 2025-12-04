package grid

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
}

func New(width, height int) *Grid {
	return &Grid{
		Width:  width,
		Height: height,
	}
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
