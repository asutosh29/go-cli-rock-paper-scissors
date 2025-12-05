package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type player struct {
	counter int
	name    string
}

type model struct {
	players       []player
	message       string
	currentPlayer int // will 0 or 1

	logMessages []string

	textInput textinput.Model
	isEditing bool

	viewPort viewport.Model
}

func (m model) Init() tea.Cmd { return nil }

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter new name"
	ti.CharLimit = 20
	ti.Width = 30

	vp := viewport.New(30, 10)

	return model{
		players: []player{
			{
				name: "Player 1",
			},
			{
				name: "Player 2",
			},
		},
		textInput: ti,
		isEditing: false,

		viewPort: vp,
	}
}
