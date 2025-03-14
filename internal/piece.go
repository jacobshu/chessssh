package game

import (
	"fmt"
	"slices"
)

type PieceType int

const (
	WhitePawn PieceType = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing

	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

var pieceName = map[PieceType]string{
	WhitePawn:   "",
	WhiteKnight: "N",
	WhiteBishop: "B",
	WhiteRook:   "R",
	WhiteQueen:  "Q",
	WhiteKing:   "K",
	BlackPawn:   "",
	BlackKnight: "N",
	BlackBishop: "B",
	BlackRook:   "R",
	BlackQueen:  "Q",
	BlackKing:   "K",
}

var pieceGlyphs = map[PieceType]string{
	WhitePawn:   "♟",
	WhiteRook:   "♜",
	WhiteKnight: "♞",
	WhiteBishop: "♝",
	WhiteKing:   "♚",
	WhiteQueen:  "♛",

	BlackPawn:   "♙",
	BlackRook:   "♖",
	BlackKnight: "♘",
	BlackBishop: "♗",
	BlackKing:   "♔",
	BlackQueen:  "♕",
}

func (p PieceType) Notation() string {
	return pieceName[p]
}

func (p PieceType) String() string {
	return pieceGlyphs[p]
}

type Piece struct {
	IsWhite    bool
	Glyph      string
	Position   Position
	Type       PieceType
	Notation   string
	IsCaptured bool
	HasMoved   bool
}

type PieceMove interface {
	GetMoves() []Position
	IsValidMove(Position, string) bool
	GetPosition() Position
}

func (p Piece) String() string {
	return fmt.Sprintf("%s@%s", p.Type.Notation(), p.Position.String())
}

func NewPiece(pos Position, pt PieceType) *Piece {
	p := Piece{
		Position:   pos,
		IsCaptured: false,
		HasMoved:   false,
	}

	whitePieces := []PieceType{WhitePawn, WhiteRook, WhiteKnight, WhiteBishop, WhiteKing, WhiteQueen}
	blackPieces := []PieceType{BlackPawn, BlackRook, BlackKnight, BlackBishop, BlackKing, BlackQueen}

	if slices.Contains(whitePieces, pt) {
		p.IsWhite = true
		p.Glyph = pt.String()
		p.Type = pt
		p.Notation = pt.Notation()
	} else if slices.Contains(blackPieces, pt) {
		p.IsWhite = true
		p.Glyph = pt.String()
		p.Type = pt
		p.Notation = pt.Notation()
	} else {
		panic("what kind of piece is it?")
	}

	return &p
}
