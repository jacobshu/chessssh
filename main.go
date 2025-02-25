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
	log.SetLevel(log.DebugLevel)
	// if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
	// 	log.Fatal(err)
	// }

	g := game.NewGame()

	p := tea.NewProgram(model{
		Game:  g,
		timer: timer.New(time.Minute * 10),
	})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type model struct {
	Game  game.Game
	timer timer.Model
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m model) Init() tea.Cmd {
	log.Info("starting...")
	return tick
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
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
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var s strings.Builder
	for i := 0; i < len(m.Game.Board); i++ {
		row := []string{}
		for j := 0; j < len(m.Game.Board[i]); j++ {
			row = append(row, m.Game.Board[i][j].Render())
		}
		s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, row...))
		s.WriteString("\n")
	}
	return s.String()
}

// Messages are events that we respond to in our Update function. This
// particular one indicates that the timer has ticked.
type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}
