package game

type Game struct {
	State   []Piece
	Players []Player
}

type Player struct {
	Name     string
	IsActive bool
}

type Position struct {
	x int
	y int
}

type Piece struct {
	IsWhite  bool
	Glyphs   []string
	Position Position
	Notation string
}

type PieceMovement interface {
	GetMoves(Position) []Position
	IsValidMove(Position, string) bool
}

func (p *Piece) positionToNotation() string {
	var rank, file string

	switch p.Position.x {
	case 0:
		file = "a"
	case 1:
		file = "b"
	case 2:
		file = "c"
	case 3:
		file = "d"
	case 4:
		file = "e"
	case 5:
		file = "f"
	case 6:
		file = "g"
	case 7:
		file = "h"
	}

	switch p.Position.y {
	case 0:
		rank = "8"
	case 1:
		rank = "7"
	case 2:
		rank = "6"
	case 3:
		rank = "5"
	case 4:
		rank = "4"
	case 5:
		rank = "3"
	case 6:
		rank = "2"
	case 7:
		rank = "1"
	}
	return file + rank
}

type Pawn struct {
	Glyphs []string
}

const (
	whitePawn   = "♟"
	whiteRook   = "♜"
	whiteKnight = "♞"
	whiteBishop = "♝"
	whiteKing   = "♚"
	whiteQueen  = "♛"

	blackPawn   = "♙"
	blackRook   = "♖"
	blackKnight = "♘"
	blackBishop = "♗"
	blackKing   = "♔"
	blackQueen  = "♕"
)
