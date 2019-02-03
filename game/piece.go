package game

func black(p Piece) *Piece {
	p.IsWhite = false
	return &p
}

func white(p Piece) *Piece {
	p.IsWhite = true
	return &p
}

func square(loc Coord) Piece {
	return Piece{CurrentLocation: loc, IsSquare: true}
}

func circle(loc Coord) Piece {
	return Piece{CurrentLocation: loc, IsSquare: false}
}

type Piece struct {
	CurrentLocation Coord
	IsSquare        bool
	IsWhite         bool
}
