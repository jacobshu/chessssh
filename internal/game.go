package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	ErrNoPiece     = errors.New("no piece at position\n")
	ErrInvalidMove = errors.New("invalid move for piece\n")
)

type BoxPart int

const (
	TopLeft BoxPart = iota
	TopIntersection
	TopRight
	MidLeft
	MidIntersection
	MidRight
	BottomLeft
	BottomIntersection
	BottomRight
	Horizontal
	Vertical
)

var boxGlyphs = map[BoxPart]string{
	TopLeft:            "┌",
	TopIntersection:    "┬",
	TopRight:           "┐",
	MidLeft:            "├",
	MidIntersection:    "┼",
	MidRight:           "┤",
	BottomLeft:         "└",
	BottomIntersection: "┴",
	BottomRight:        "┘",
	Horizontal:         "─",
	Vertical:           "│",
}

func (b BoxPart) String() string {
	return boxGlyphs[b]
}

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

type Game struct {
	Board   [8][8]Tile
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
			Position:   Position{Rank: Rank2, File: FileH},
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
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileH},
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
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileH},
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
			IsWhite:    true,
			Glyph:      WhitePawn.String(),
			Position:   Position{Rank: Rank2, File: FileH},
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
			Position:   Position{Rank: Rank7, File: FileH},
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
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileH},
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
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileH},
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
		{
			IsWhite:    false,
			Glyph:      BlackPawn.String(),
			Position:   Position{Rank: Rank7, File: FileH},
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
	var b [8][8]Tile
	for i, r := range ranks {
		row := [8]Tile{}
		for j, f := range files {
			for _, p := range pcs {
				pos := Position{Rank: r, File: f}
				if p.Position == pos {
					row[j] = NewTile(pos, &p)
				} else {
					row[j] = NewTile(pos, nil)
				}
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

type Player struct {
	Name     string
	IsActive bool
}

type Tile struct {
	Position        Position
	IsOccupied      bool
	IsDark          bool
	IsPotentialMove bool
	IsSelected      bool
	Occupant        *Piece
}

func NewTile(p Position, o *Piece) Tile {
	t := Tile{
		Position: p,
	}

	if o != nil {
		t.IsOccupied = true
		t.Occupant = o
	} else {
		t.IsOccupied = false
	}

	isDark := true
	if (p.File%2 == 0 && p.Rank%2 == 1) || (p.File%2 == 1 && p.Rank%2 == 0) {
		isDark = false
	}

	t.IsDark = isDark
	return t
}

func (t Tile) Render() string {
	log.Debug("render tile", "pos", t.Position)

	var c lipgloss.ANSIColor
	if t.IsDark {
		c = lipgloss.ANSIColor(235)
	} else if t.IsPotentialMove {
		c = lipgloss.ANSIColor(42)
	} else if t.IsSelected {
		c = lipgloss.ANSIColor(102)
	}
	style := lipgloss.NewStyle().Background(c).Foreground(lipgloss.ANSIColor(243))
	pieceStyle := style.Foreground(lipgloss.ANSIColor(7))

	top := style.Render("   ")
	var mid string
	if t.IsOccupied {
		s := style.Render(" ")
		p := pieceStyle.Render(t.Occupant.Glyph)
		mid = s + p + s
	} else {
		mid = style.Render("   ")
	}

	bot := style.Render("   ")

	final := lipgloss.JoinVertical(lipgloss.Left, top, mid, bot)
	return final
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
