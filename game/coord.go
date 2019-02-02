package game

import "fmt"

type Coord struct {
	X int
	Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}
