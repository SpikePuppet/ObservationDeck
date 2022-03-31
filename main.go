package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	choices []string
}

func initalModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Welcome to the Observation Deck! v0.1"
}

func main() {
	p := tea.NewProgram(initalModel())
	if err := p.Start(); err != nil {
		fmt.Printf("The deck is burning, error: %v", err)
		os.Exit(1)
	}
}
