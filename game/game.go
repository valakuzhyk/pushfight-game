package game

import "fmt"

type Game struct {
	pieces []Piece
	Anchor *Coord
	Board
}

// CreateGame creates a game of PushFight
func CreateGame() Game {
	return Game{
		pieces: []Piece{
			/*                       */ black(square(Coord{4, 3})), white(square(Coord{5, 3})),
			/*                       */ black(circle(Coord{4, 2})), white(circle(Coord{5, 2})), white(square(Coord{6, 2})),
			black(square(Coord{3, 1})), black(circle(Coord{4, 1})), white(circle(Coord{5, 1})),
			/*                       */ black(square(Coord{4, 0})), white(square(Coord{5, 0})),
		},
		Anchor: nil,
		Board:  CreateBoard(),
	}

}

func (g Game) String() string {
	output := ""
	for j := g.Board.Height - 1; j >= 0; j-- {
		for i := 0; i < g.Board.Width; i++ {
			output += fmt.Sprint(g.GetCoordType(Coord{i, j}).CoordSymbol())
		}
		output += fmt.Sprint("\n")
	}
	return output
}

// GetCoordType returns what's going on at a certain space.
func (g Game) GetCoordType(location Coord) CoordType {
	if !g.IsValidSpace(location) {
		return InvalidCoord
	} else if !g.IsOnBoard(location) {
		return OffTheBoard
	}

	for _, piece := range g.pieces {
		if piece.CurrentLocation == location {
			switch true {
			case piece.IsSquare && piece.IsWhite:
				return WhiteSquare
			case piece.IsSquare && !piece.IsWhite:
				return BlackSquare
			case !piece.IsSquare && piece.IsWhite:
				return WhiteCircle
			case !piece.IsSquare && !piece.IsWhite:
				return BlackCircle
			}
		}
	}

	return Open
}
