package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.isEditing {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.textInput.Value() != "" {
					m.players[m.currentPlayer].name = m.textInput.Value()
				}

				m.isEditing = false
				m.textInput.Blur()
				return m, nil
			case "esc":
				m.isEditing = false
				m.textInput.Blur()
				return m, nil
			}
		}
		var cmds []tea.Cmd
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)
		m.viewPort, cmd = m.viewPort.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.message = "Quitting thanks for playing!"
			m.logMessages = append(m.logMessages, m.message)
			return m, tea.Quit
		case "e":
			m.isEditing = true
			m.textInput.Focus()
			m.textInput.SetValue(m.players[m.currentPlayer].name)

			m.message = "Editting name... press esc to exit!"
			m.logMessages = append(m.logMessages, "Editting name")
			return m, textinput.Blink

		case "up", "j":
			if m.players[m.currentPlayer].counter < 10 {
				m.players[m.currentPlayer].counter++
				m.message = ""

				logEntry := fmt.Sprintf("%s scored! Total: %d", m.players[m.currentPlayer].name, m.players[m.currentPlayer].counter)
				m.logMessages = append(m.logMessages, logEntry)

				m.viewPort.SetContent(strings.Join(m.logMessages, "\n"))
				m.viewPort.GotoBottom()
			} else {
				m.message = "Can't go above 10..."
			}
		case "down", "k":
			if m.players[m.currentPlayer].counter > 0 {
				m.players[m.currentPlayer].counter--
				m.message = ""

				logEntry := fmt.Sprintf("%s lost a point. Total: %d", m.players[m.currentPlayer].name, m.players[m.currentPlayer].counter)
				m.logMessages = append(m.logMessages, logEntry)

				m.viewPort.SetContent(strings.Join(m.logMessages, "\n"))
				m.viewPort.GotoBottom()
			} else {
				m.message = "Can't go below 0..."
			}
		case "tab":
			m.currentPlayer = (m.currentPlayer + 1) % (len(m.players))
		}
	}
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.viewPort, cmd = m.viewPort.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
