package game

// CoordType classifies the coordinate based on what's going on.
type CoordType int

func (c CoordType) String() string {
	switch c {
	case InvalidCoord:
		return "InvalidCoord"
	case OffTheBoard:
		return "OffTheBoard"
	case WhiteSquare:
		return "WhiteSquare"
	case BlackSquare:
		return "BlackSquare"
	case WhiteCircle:
		return "WhiteCircle"
	case BlackCircle:
		return "BlackCircle"
	case Open:
		return "Open"
	}
	return "UnknownCoordType"
}

const (
	InvalidCoord CoordType = iota
	OffTheBoard
	WhiteSquare
	BlackSquare
	WhiteCircle
	BlackCircle
	Open
)

func (c CoordType) CoordSymbol() string {
	switch c {
	case OffTheBoard:
		return "X"
	case WhiteSquare:
		return "O"
	case BlackSquare:
		return "O"
	case WhiteCircle:
		return "o"
	case BlackCircle:
		return "o"
	case Open:
		return "_"
	}
	return "?"
}
