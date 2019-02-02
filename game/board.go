package game

// Board represents the physical board.
type Board struct {
	Height, Width  int
	OffBoardSpaces []Coord
}

// CreateBoard returns the board for PushFight!
func CreateBoard() Board {
	return Board{
		Height: 4,
		Width:  10,
		OffBoardSpaces: []Coord{
			{0, 3}, {1, 3}, {2, 3} /*   */, {8, 3}, {9, 3},
			{0, 2} /*                           */, {9, 2},
			{0, 1} /*                           */, {9, 1},
			{0, 0}, {1, 0} /*   */, {7, 0}, {8, 0}, {9, 0},
		},
	}
}

// IsValidSpace returns whether a piece can be moved to this space.
func (b Board) IsValidSpace(loc Coord) bool {
	if loc.Y < 0 || b.Height <= loc.Y ||
		loc.X < 0 || b.Width <= loc.X {
		return false
	}
	return true
}

// IsOnBoard returns true if a piece can stand on this spot without losing the game.
func (b Board) IsOnBoard(loc Coord) bool {
	if !b.IsValidSpace(loc) {
		return false
	}
	for _, offBoardSpace := range b.OffBoardSpaces {
		if offBoardSpace == loc {
			return false
		}
	}
	return true
}
