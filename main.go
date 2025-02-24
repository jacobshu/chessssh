package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	game "github.com/jacobshu/chessssh/internal"
)

func main() {
	//logger := log.New(os.Stdout)
	logfilePath := "debug.log" // os.Getenv("BUBBLETEA_LOG")
	if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(model{Game: game.NewGame()})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type model struct {
	Game game.Game
	Time time.Duration
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m model) Init() tea.Cmd {
	return tick
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}
	}
	return m, nil
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m model) View() string {
	var s strings.Builder
	for i := 0; i < len(m.Game.Board); i++ {
		for j := 0; j < len(m.Game.Board[i]); j++ {
			s.WriteString(m.Game.Board[i][j].View())
		}
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
