package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
