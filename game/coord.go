package game

import "fmt"

type Coord struct {
	X int
	Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type Direction int

const (
	Invalid Direction = iota
	Up
	Down
	Left
	Right
)

// Move the coordinate in the given direction
func (c Coord) Move(dir Direction) Coord {
	switch dir {
	case Up:
		return Coord{c.X, c.Y + 1}
	case Down:
		return Coord{c.X, c.Y - 1}
	case Right:
		return Coord{c.X + 1, c.Y}
	case Left:
		return Coord{c.X - 1, c.Y}
	default:
		return c
	}
}
