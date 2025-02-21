package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	ErrNoPiece     = errors.New("no piece at position\n")
	ErrInvalidMove = errors.New("invalid move for piece\n")
)

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

type Game struct {
	Board   [8][8]Tile
	Pieces  [32]Piece
	Players []Player
}

func (g Game) String(asWhite bool) string {
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

	tilesAcross := 8
	tileWidth := 3
	var b strings.Builder

	b.WriteString(strings.Repeat("-", tilesAcross*(tileWidth+1)))
	if asWhite {
		for i, j := 0, len(ranks)-1; i < j; i, j = i+1, j-1 {
			ranks[i], ranks[j] = ranks[j], ranks[i]
		}
	} else {
		for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
			files[i], files[j] = files[j], files[i]
		}
	}

	for r := 0; r < len(ranks); r++ {
		b.WriteString("\n|")
		b.WriteString(strings.Repeat("   |", tilesAcross))
		b.WriteString("\n|")
		b.WriteString(strings.Repeat(fmt.Sprintf(" %s |", whitePawn), tilesAcross))
		b.WriteString("\n|")
		for f := 0; f < len(files); f++ {
			b.WriteString(color.HiBlackString("%s%s ", files[f], ranks[r]))
			b.WriteString("|")
		}
		b.WriteString("\n")
		b.WriteString(strings.Repeat("-", tilesAcross*(tileWidth+1)))
	}

	return b.String()
}

type Player struct {
	Name     string
	IsActive bool
}

type Tile struct {
	Position   Position
	IsOccupied bool
	Occupant   *Piece
}

type Piece struct {
	IsWhite    bool
	Glyphs     []string
	Position   Position
	Value      int
	Notation   string
	IsCaptured bool
}

type PieceMove interface {
	GetMoves(*Game) []Position
	IsValidMove(Position, string) bool
	GetPosition() Position
}

func (g *Game) PieceAt(p Position) Piece {
	return *g.Board[p.Rank][p.File].Occupant
}

type Rank int

const (
	Rank1 Rank = iota + 1
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

func (r Rank) Index() int {
	return int(r)
}

func (r Rank) String() string {
	return fmt.Sprintf("%d", int(r))
}

type File int

const (
	FileA File = iota + 1
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) Index() int {
	return int(f)
}

func (f File) String() string {
	return [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}[f-1]
}

type Position struct {
	Rank Rank
	File File
}

func (p Position) String() string {
	return fmt.Sprintf("%s%s", p.Rank.String(), p.File.String())
}

func (p Position) ToBlack(step int) Position {
	return Position{
		Rank: Rank(p.Rank.Index() + step),
		File: p.File,
	}
}

func (p Position) ToWhite(step int) Position {
	return Position{
		Rank: Rank(p.Rank.Index() - step),
		File: p.File,
	}
}

func (p Position) ToA(step int) Position {
	return Position{
		Rank: p.Rank,
		File: File(p.File.Index() - step),
	}
}

func (p Position) ToH(step int) Position {
	return Position{
		Rank: p.Rank,
		File: File(p.File.Index() + step),
	}
}

type Pawn struct {
	Piece
	HasMoved bool
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
func (p *Pawn) GetMoves(g *Game) []Position {
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

func NewGame() Game {
	g := Game{
		Board: [8][8]Tile{},
	}

	return g
}
