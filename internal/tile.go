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
	log.Debug("render tile", "pos", t.Position, "occ", t.Occupant)

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

	// bot := style.Render("   ")

	final := lipgloss.JoinVertical(lipgloss.Left, top, mid) //, bot)
	return final
}
