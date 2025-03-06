package game

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type Tile struct {
	Position        Position
	IsOccupied      bool
	IsDark          bool
	IsPotentialMove bool
	IsSelected      bool
	IsHovered       bool
	Occupant        *Piece
	OffsetX         int
	OffsetY         int
	Moves           []Position
}

func NewTile(p Position, o *Piece) *Tile {
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
	return &t
}

func (t Tile) Info() string {
	var i string
	if t.Position.Rank == Rank1 && t.Position.File == FileA {
		i = t.Position.String()
	} else if t.Position.Rank == Rank1 {
		i = t.Position.File.String() + " "
	} else if t.Position.File == FileA {
		i = t.Position.Rank.String() + " "
	} else {
		i = "  "
	}

	return i + " "
}

func (t Tile) SetSelected() {
	t.IsSelected = true
	if t.IsOccupied {
		// t.Occupant.GetMoves()
	}
}

func (t Tile) Render() string {
	log.Debug("render tile", "pos", t.Position, "occ", t.Occupant)

	var c lipgloss.ANSIColor
	if t.IsDark {
		c = lipgloss.ANSIColor(235)
	} else {
		c = lipgloss.ANSIColor(247)
	}

	if t.IsPotentialMove {
		c = lipgloss.ANSIColor(42)
	} else if t.IsSelected {
		c = lipgloss.ANSIColor(6)
	} else if t.IsHovered {
		c = lipgloss.ANSIColor(2)
	}

	style := lipgloss.NewStyle().Background(c).Foreground(lipgloss.ANSIColor(240))
	bStyle := style.Foreground(lipgloss.ANSIColor(243))
	wStyle := style.Foreground(lipgloss.ANSIColor(15))

	// top := style.Render("   ")
	top := style.Render(t.Info())
	var mid string
	if t.IsOccupied {
		s := style.Render(" ")
		var p string
		if t.Occupant.IsWhite {
			p = wStyle.Render(t.Occupant.Glyph)
		} else {
			p = bStyle.Render(t.Occupant.Glyph)
		}
		mid = s + p + s
	} else {
		mid = style.Render("   ")
	}

	// bot := style.Render("   ")

	final := lipgloss.JoinVertical(lipgloss.Left, top, mid) //, bot)
	return final
}
