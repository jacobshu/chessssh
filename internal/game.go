package game

import (
	"errors"
	"fmt"
)

var (
	ErrNoPiece     = errors.New("no piece at position\n")
	ErrInvalidMove = errors.New("invalid move for piece\n")
)

type Game struct {
	Board   [8][8]*Tile
	Pieces  [32]Piece
	Players []Player
}

func NewGame() Game {
	pcs := [32]Piece{
		{
			IsWhite:    true,
			Glyph:      WhiteRook.String(),
			Position:   Position{Rank: Rank1, File: FileH},
			Type:       WhiteRook,
			Notation:   WhiteRook.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhiteKnight.String(),
			Position:   Position{Rank: Rank1, File: FileG},
			Type:       WhiteKnight,
			Notation:   WhiteKnight.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhiteBishop.String(),
			Position:   Position{Rank: Rank1, File: FileF},
			Type:       WhiteBishop,
			Notation:   WhiteBishop.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhiteKing.String(),
			Position:   Position{Rank: Rank1, File: FileE},
			Type:       WhiteKing,
			Notation:   WhiteKing.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhiteQueen.String(),
			Position:   Position{Rank: Rank1, File: FileD},
			Type:       WhiteQueen,
			Notation:   WhiteQueen.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhiteBishop.String(),
			Position:   Position{Rank: Rank1, File: FileC},
			Type:       WhiteBishop,
			Notation:   WhiteBishop.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhiteKnight.String(),
			Position:   Position{Rank: Rank1, File: FileB},
			Type:       WhiteKnight,
			Notation:   WhiteKnight.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhiteRook.String(),
			Position:   Position{Rank: Rank1, File: FileA},
			Type:       WhiteRook,
			Notation:   WhiteRook.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileA},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileB},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileC},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileD},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileE},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileF},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileG},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileH},
			Type:       WhitePawn,
			Notation:   WhitePawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackRook.String(),
			Position:   Position{Rank: Rank8, File: FileH},
			Type:       BlackRook,
			Notation:   BlackRook.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackKnight.String(),
			Position:   Position{Rank: Rank8, File: FileG},
			Type:       BlackKnight,
			Notation:   BlackKnight.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackBishop.String(),
			Position:   Position{Rank: Rank8, File: FileF},
			Type:       BlackBishop,
			Notation:   BlackBishop.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackKing.String(),
			Position:   Position{Rank: Rank8, File: FileE},
			Type:       BlackKing,
			Notation:   BlackKing.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackQueen.String(),
			Position:   Position{Rank: Rank8, File: FileD},
			Type:       BlackQueen,
			Notation:   BlackQueen.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackBishop.String(),
			Position:   Position{Rank: Rank8, File: FileC},
			Type:       BlackBishop,
			Notation:   BlackBishop.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackKnight.String(),
			Position:   Position{Rank: Rank8, File: FileB},
			Type:       BlackKnight,
			Notation:   BlackKnight.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackRook.String(),
			Position:   Position{Rank: Rank8, File: FileA},
			Type:       BlackRook,
			Notation:   BlackRook.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileA},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileB},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileC},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileD},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileE},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileF},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileG},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileH},
			Type:       BlackPawn,
			Notation:   BlackPawn.Notation(),
			IsCaptured: false,
			HasMoved:   false,
		},
	}

	var ranks = [8]Rank{
		Rank1,
		Rank2,
		Rank3,
		Rank4,
		Rank5,
		Rank6,
		Rank7,
		Rank8,
	}

	var files = [8]File{
		FileA,
		FileB,
		FileC,
		FileD,
		FileE,
		FileF,
		FileG,
		FileH,
	}

	var b [8][8]*Tile
	for i, r := range ranks {
		row := [8]*Tile{}
		for j, f := range files {

			pos := Position{Rank: r, File: f}
			found := false

			for _, p := range pcs {
				if p.Position == pos {
					found = true
					nt := NewTile(pos, &p)
					row[j] = nt
				}

				if found {
					break
				}
			}

			if !found {
				row[j] = NewTile(pos, nil)
			}

		}
		b[i] = row
	}

	g := Game{
		Board:  b,
		Pieces: pcs,
	}

	return g
}

func (g *Game) PieceAt(p Position) Piece {
	return *g.Board[p.Rank][p.File].Occupant
}

//	A      B      C      D      E      F      G      H
//
// 8 (0,0)  (1,0)  (2,0)  (3,0)  (4,0)  (5,0)  (6,0)  (7,0)
// 7 (0,1)  (1,1)  (2,1)  (3,1)  (4,1)  (5,1)  (6,1)  (7,1)
// 6 (0,2)  (1,2)  (2,2)  (3,2)  (4,2)  (5,2)  (6,2)  (7,2)
// 5 (0,3)  (1,3)  (2,3)  (3,3)  (4,3)  (5,3)  (6,3)  (7,3)
// 4 (0,4)  (1,4)  (2,4)  (3,4)  (4,4)  (5,4)  (6,4)  (7,4)
// 3 (0,5)  (1,5)  (2,5)  (3,5)  (4,5)  (5,5)  (6,5)  (7,5)
// 2 (0,6)  (1,6)  (2,6)  (3,6)  (4,6)  (5,6)  (6,6)  (7,6)
// 1 (0,7)  (1,7)  (2,7)  (3,7)  (4,7)  (5,7)  (6,7)  (7,7)
func GetPawnMoves(p Piece) []Position {
	var moves = []Position{}

	var attackToA, attackToH, forward, forwardTwo Position
	if p.IsWhite {
		attackToA = p.Position.ToBlack(1).ToA(1)
		attackToH = p.Position.ToBlack(1).ToH(1)
		forward = p.Position.ToBlack(1)
		forwardTwo = p.Position.ToBlack(2)
	} else {
		attackToA = p.Position.ToWhite(1).ToA(1)
		attackToH = p.Position.ToWhite(1).ToH(1)
		forward = p.Position.ToWhite(1)
		forwardTwo = p.Position.ToWhite(2)
	}

	fmt.Printf("%v, %v, %v, %v", attackToA, attackToH, forward, forwardTwo)
	//default
	moves = append(moves, forward)
	if !p.HasMoved {
		moves = append(moves, forward, forwardTwo)
	}
	return moves
}
