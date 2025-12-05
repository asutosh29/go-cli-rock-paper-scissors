package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Rename Actions
func StartRename(m *model, msg tea.Msg) (model, tea.Cmd) {
	m.isEditing = true
	m.textInput.Focus()
	m.textInput.SetValue(m.players[m.selectedPlayer].name)

	m.message = "Renaming..."
	return *m, textinput.Blink
}
func SubmitRename(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.textInput.Value() != "" {
		m.players[m.selectedPlayer].name = m.textInput.Value()
	}

	m.isEditing = false
	m.textInput.Blur()
	return *m, nil
}
func QuitRename(m *model, msg tea.Msg) (model, tea.Cmd) {
	m.isEditing = false
	m.textInput.Blur()
	return *m, nil
}

// Game actions
func Increment(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.players[m.selectedPlayer].counter < 10 {
		m.players[m.selectedPlayer].counter++
		m.message = ""

		logEntry := fmt.Sprintf("%s scored! Total: %d", m.players[m.selectedPlayer].name, m.players[m.selectedPlayer].counter)
		m.logMessages = append(m.logMessages, logEntry)

		m.viewPort.SetContent(strings.Join(m.logMessages, "\n"))
		m.viewPort.GotoBottom()
	} else {
		m.message = "Can't go above 10..."
	}
	return *m, nil
}
func Decrement(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.players[m.selectedPlayer].counter > 0 {
		m.players[m.selectedPlayer].counter--
		m.message = ""

		logEntry := fmt.Sprintf("%s lost a point. Total: %d", m.players[m.selectedPlayer].name, m.players[m.selectedPlayer].counter)
		m.logMessages = append(m.logMessages, logEntry)

		m.viewPort.SetContent(strings.Join(m.logMessages, "\n"))
		m.viewPort.GotoBottom()
	} else {
		m.message = "Can't go below 0..."
	}
	return *m, nil
}

func QuitGame(m *model, msg tea.Msg) (model, tea.Cmd) {
	m.message = "Thanks for playing!"
	return *m, tea.Quit
}

// Timer Actions
func UpdateTimer(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.timeLeft > 1 {
		m.timeLeft--
		// Update text based on time left
		if m.timeLeft == 3 {
			m.message = "ROCK!"
		}
		if m.timeLeft == 2 {
			m.message = "PAPER!"
		}
		if m.timeLeft == 1 {
			m.message = "SCISSORS!"
		}
		return *m, tick()
	} else {
		m.state = stateResult
		m.determineWinner()
		return *m, nil
	}
}
