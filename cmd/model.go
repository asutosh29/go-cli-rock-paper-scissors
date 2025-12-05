package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Game Types

type gameState int

const (
	stateLobby gameState = iota
	stateCountdown
	stateResult
)

type move int

const (
	moveNone move = iota
	moveRock
	movePaper
	moveScissors
)

func (m move) String() string {
	switch m {
	case moveRock:
		return "Rock"
	case movePaper:
		return "Paper"
	case moveScissors:
		return "Scissors"
	default:
		return "Waiting..."
	}
}

type player struct {
	counter int
	name    string
	move    move
	wins    int
}

// Bubble Tea Types
type model struct {
	state gameState

	players        []player
	message        string
	selectedPlayer int // will 0 or 1
	timeLeft       time.Duration

	logMessages []string

	textInput textinput.Model
	isEditing bool

	viewPort viewport.Model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, tick())
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter new name"
	ti.CharLimit = 20
	ti.Width = 30

	vp := viewport.New(48, 6)

	return model{
		state: stateLobby,
		players: []player{
			{
				name: "Player 1",
				move: moveNone,
				wins: 0,
			},
			{
				name: "Player 2",
				move: moveNone,
				wins: 0,
			},
		},
		timeLeft: 4, // 4 for states Ready, Rock, Paper, Scissors

		textInput: ti,
		isEditing: false,

		viewPort: vp,
	}
}

func (m *model) determineWinner() {
	// No moves so draw
	if m.players[0].move == moveNone && m.players[1].move == moveNone {
		m.message = "Draw! (No moves)"
		return
	}

	if m.players[0].move == moveNone {
		m.message = fmt.Sprintf("%s didn't select a move!", m.players[0].name)
		return
	}
	if m.players[1].move == moveNone {
		m.message = fmt.Sprintf("%s didn't select a move!", m.players[1].name)
		return
	}

	// Game Logic
	result := ""

	if m.players[0].move == m.players[1].move {
		result = "It's a Draw!"
	} else if (m.players[0].move == moveRock && m.players[1].move == moveScissors) ||
		(m.players[0].move == movePaper && m.players[1].move == moveRock) ||
		(m.players[0].move == moveScissors && m.players[1].move == movePaper) {
		// Player 1 Wins
		m.players[0].wins++
		result = fmt.Sprintf("%s Wins!", m.players[0].name)
	} else {
		// Player 2 Wins
		m.players[1].wins++
		result = fmt.Sprintf("%s Wins!", m.players[1].name)
	}

	// Log it
	logMsg := fmt.Sprintf("Round: %s vs %s -> %s", m.players[0].move, m.players[1].move, result)
	m.logMessages = append(m.logMessages, logMsg)

	// TODO: Move this to UI
	m.viewPort.SetContent(strings.Join(m.logMessages, "\n"))
	m.viewPort.GotoBottom()

	m.message = result
}
