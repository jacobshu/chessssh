package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	game "github.com/jacobshu/chessssh/internal"
)

func main() {
	logfilePath := "debug.log"                                                     // os.Getenv("BUBBLETEA_LOG")
	f, err := os.OpenFile(logfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600) //nolint:gomnd
	if err != nil {
		fmt.Printf("error opening file for logging: %v", err)
	}
	log.SetOutput(f)
	log.SetLevel(log.ErrorLevel)
	// if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
	// 	log.Fatal(err)
	// }

	g := game.NewGame()

	p := tea.NewProgram(model{
		Game:       &g,
		timer:      timer.New(time.Minute * 10),
		boardSize:  8,
		padding:    1,
		tileWidth:  3,
		tileHeight: 2,
	}, tea.WithMouseAllMotion(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type model struct {
	Game         *game.Game
	timer        timer.Model
	mouseEvent   tea.MouseEvent
	info         string
	termWidth    int
	termHeight   int
	padding      int
	tileWidth    int
	tileHeight   int
	boardSize    int
	boardOffsetX int
	boardOffsetY int
}

func (m model) Init() tea.Cmd {
	log.Info("starting...")
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		m.boardOffsetX = (msg.Width - (m.boardSize * m.tileWidth)) / 2
		m.boardOffsetY = (msg.Height - (m.boardSize * m.tileHeight)) / 2
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}
	case timer.TickMsg:
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)
	case tea.MouseMsg:
		m.info = fmt.Sprintf("(%d, %d) %s\n", msg.X, msg.Y, tea.MouseEvent(msg))
		m.info = fmt.Sprintf("%s(%d, %d) board offset\n", m.info, m.boardOffsetX, m.boardOffsetY)
		if msg.Action == tea.MouseActionRelease {
			m.SetSelectedTile(msg.X, msg.Y)
		}
		m.SetHoveredTile(msg.X, msg.Y)
		return m, nil
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	boardStyle := lipgloss.NewStyle().
		Align(lipgloss.Center).
		BorderForeground(lipgloss.ANSIColor(245))

	var s strings.Builder
	for i := range m.Game.Board {
		row := []string{}
		for j := range m.Game.Board[i] {
			row = append(row, m.Game.Board[i][j].Render())
		}
		s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, row...))
		s.WriteString("\n")
	}
	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		boardStyle.Render(s.String())+"\n"+m.info)
}

func (m model) SetHoveredTile(x, y int) {
	log.Debug("SetHoveredTile", "x", x, "y", y)
	for i, r := range m.Game.Board {
		for j, f := range r {
			txl := m.boardOffsetX + (j * 3)
			txh := txl + 2
			tyl := m.boardOffsetY + (i * 2) - 2 // -2 for info lines
			tyh := tyl + 1
			if txl <= x && txh >= x && tyl <= y && tyh >= y {
				f.IsHovered = true
			} else {
				f.IsHovered = false
			}
		}
	}
}

func (m model) SetSelectedTile(x, y int) {
	log.Debug("SetHoveredTile", "x", x, "y", y)
	for i, r := range m.Game.Board {
		for j, f := range r {
			txl := m.boardOffsetX + (j * 3)
			txh := txl + 2
			tyl := m.boardOffsetY + (i * 2) - 2 // -2 for info lines
			tyh := tyl + 1
			if txl <= x && txh >= x && tyl <= y && tyh >= y {
				f.IsSelected = true
			} else {
				f.IsSelected = false
			}
		}
	}
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}
