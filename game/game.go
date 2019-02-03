package game

import "fmt"

type Game struct {
	pieces []*Piece
	Anchor *Coord
	Board
}

// Push attempts to push a piece one square to a side, moving all the pieces in the way.
func (g Game) Push(p *Piece, moveDir Direction) error {
	if moveDir == Invalid {
		return fmt.Errorf("Invalid direction")
	}

	// Check the place in that direction, if it is full, add that to the list and check the next one
	// If the place is empty, add it to the list and check the next one.
	pushList := []*Piece{p}
	for true {
		adjacentSpot := p.CurrentLocation.Move(moveDir)
		coordType := g.GetCoordType(adjacentSpot)
		switch coordType {
		case InvalidCoord:
			return fmt.Errorf("Cannot move here due to wall")
		case Open:
			for _, p := range pushList {
				p.CurrentLocation = p.CurrentLocation.Move(moveDir)
			}
			return nil
		case OffTheBoard:
			for _, p := range pushList {
				p.CurrentLocation = p.CurrentLocation.Move(moveDir)
			}
			if pushList[len(pushList)-1].IsWhite {
				return fmt.Errorf("Game is finished: Black wins")
			} else {
				return fmt.Errorf("Game is finished: White wins")
			}
			// complete the movement and end the game
		case WhiteCircle, WhiteSquare, BlackCircle, BlackSquare:
			nextPiece := g.GetPieceAt(adjacentSpot)
			pushList = append(pushList, nextPiece)
		}

	}

	if err := g.push(p, moveDir); err != nil {
		*g.Anchor = p.CurrentLocation
	}

	return nil
}

func (g Game) push(p *Piece, moveDir Direction) error {

	return nil
}

// Move attempts to move a piece to that location.
func (g Game) Move(p *Piece, loc Coord) bool {
	possibleMoves := g.ValidMoves(*p)
	for _, move := range possibleMoves {
		if move == loc {
			p.CurrentLocation = loc
			return true
		}
	}
	return false
}

// ValidMoves returns the spaces that a piece can move to.
func (g Game) ValidMoves(p Piece) []Coord {
	canAccess := map[Coord]bool{p.CurrentLocation: true}
	toProcess := []Coord{p.CurrentLocation}
	validMoves := []Coord{}

	for len(toProcess) > 0 {
		nextSpot := toProcess[0]
		toProcess = toProcess[1:]
		neighboringSpots := g.getOpenNeighborSpots(nextSpot)
		for _, spot := range neighboringSpots {
			if _, alreadyChecked := canAccess[spot]; alreadyChecked {
				continue
			}
			canAccess[spot] = true
			validMoves = append(validMoves, spot)
			toProcess = append(toProcess, spot)
		}
	}
	return validMoves
}

func (g Game) getOpenNeighborSpots(loc Coord) []Coord {
	directions := []Direction{Up, Down, Left, Right}

	openNeighbors := []Coord{}
	for _, direction := range directions {
		neighbor := loc.Move(direction)
		if g.GetCoordType(neighbor) == Open {
			openNeighbors = append(openNeighbors, neighbor)
		}
	}
	return openNeighbors
}

// CreateGame creates a game of PushFight
func CreateGame() Game {
	return Game{
		pieces: []*Piece{
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

func (g Game) GetPieceAt(location Coord) *Piece {
	for _, p := range g.pieces {
		if p.CurrentLocation == location {
			return p
		}
	}
	return nil
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
