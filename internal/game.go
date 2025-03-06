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
	Pieces  [32]*Piece
	Players []Player
}

func NewGame() Game {
	pcs := [32]*Piece{
		NewPiece(Position{Rank: Rank1, File: FileH}, WhiteRook),
		NewPiece(Position{Rank: Rank1, File: FileG}, WhiteKnight),
		NewPiece(Position{Rank: Rank1, File: FileF}, WhiteBishop),
		NewPiece(Position{Rank: Rank1, File: FileE}, WhiteKing),
		NewPiece(Position{Rank: Rank1, File: FileD}, WhiteQueen),
		NewPiece(Position{Rank: Rank1, File: FileC}, WhiteBishop),
		NewPiece(Position{Rank: Rank1, File: FileB}, WhiteKnight),
		NewPiece(Position{Rank: Rank1, File: FileA}, WhiteRook),
		NewPiece(Position{Rank: Rank2, File: FileA}, WhitePawn),
		NewPiece(Position{Rank: Rank2, File: FileB}, WhitePawn),
		NewPiece(Position{Rank: Rank2, File: FileC}, WhitePawn),
		NewPiece(Position{Rank: Rank2, File: FileD}, WhitePawn),
		NewPiece(Position{Rank: Rank2, File: FileE}, WhitePawn),
		NewPiece(Position{Rank: Rank2, File: FileF}, WhitePawn),
		NewPiece(Position{Rank: Rank2, File: FileG}, WhitePawn),
		NewPiece(Position{Rank: Rank2, File: FileH}, WhitePawn),
		NewPiece(Position{Rank: Rank8, File: FileH}, BlackRook),
		NewPiece(Position{Rank: Rank8, File: FileG}, BlackKnight),
		NewPiece(Position{Rank: Rank8, File: FileF}, BlackBishop),
		NewPiece(Position{Rank: Rank8, File: FileE}, BlackKing),
		NewPiece(Position{Rank: Rank8, File: FileD}, BlackQueen),
		NewPiece(Position{Rank: Rank8, File: FileC}, BlackBishop),
		NewPiece(Position{Rank: Rank8, File: FileB}, BlackKnight),
		NewPiece(Position{Rank: Rank8, File: FileA}, BlackRook),
		NewPiece(Position{Rank: Rank7, File: FileA}, BlackPawn),
		NewPiece(Position{Rank: Rank7, File: FileB}, BlackPawn),
		NewPiece(Position{Rank: Rank7, File: FileC}, BlackPawn),
		NewPiece(Position{Rank: Rank7, File: FileD}, BlackPawn),
		NewPiece(Position{Rank: Rank7, File: FileE}, BlackPawn),
		NewPiece(Position{Rank: Rank7, File: FileF}, BlackPawn),
		NewPiece(Position{Rank: Rank7, File: FileG}, BlackPawn),
		NewPiece(Position{Rank: Rank7, File: FileH}, BlackPawn),
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
					nt := NewTile(pos, p)
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
