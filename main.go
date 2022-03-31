package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
	"strings"
)

type model struct {
	dockerImages []string
	cursor       int
}

func initalModel() model {
	return model{}
}

type dockerImagesOutput []string

func getDockerImages() tea.Msg {
	command := exec.Command("bash", "-c", "docker image ls | awk '(NR>1) {print $1}'")
	out, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("docker image ls failed with error: %v", err)
	}
	output := strings.TrimSpace(string(out))
	return dockerImagesOutput(strings.Split(output, "\n"))
}

func (m model) Init() tea.Cmd {
	return getDockerImages
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.dockerImages)-1 {
				m.cursor++
			}
		}

	case dockerImagesOutput:
		m.dockerImages = msg
	}
	return m, nil
}

func (m model) View() string {
	s := "Welcome to the Observation Deck! v0.1\n\n"

	for index, image := range m.dockerImages {
		cursor := " "
		if m.cursor == index {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, image)
	}

	return s
}

func main() {
	p := tea.NewProgram(initalModel())
	if err := p.Start(); err != nil {
		fmt.Printf("The deck is burning, error: %v", err)
		os.Exit(1)
	}
}
