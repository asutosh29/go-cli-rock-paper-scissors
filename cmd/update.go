package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.isEditing {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				return SubmitRename(&m, msg)
			case "esc":
				return QuitRename(&m, msg)
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
			return QuitGame(&m, msg)
		}
	}

	switch m.state {
	case stateLobby:
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case " ":
				m.state = stateCountdown
				m.timeLeft = 4
				m.players[0].move = moveNone
				m.players[1].move = moveNone
				m.message = "LET's GO!!"
				return m, tick()
			case "tab":
				m.selectedPlayer = (m.selectedPlayer + 1) % (len(m.players))
			case "e":
				return StartRename(&m, msg)
			}
		}
		m.message = "Press SPACE to start round"

	// ROCK PAPER SCISSOR Countdown
	// Updates Timer and player press keys to select their moves
	case stateCountdown:
		switch msg := msg.(type) {
		case tickMsg:
			return UpdateTimer(&m, msg)
		case tea.KeyMsg:
			switch msg.String() {
			// Player 1
			case "a":
				m.players[0].move = moveRock
			case "s":
				m.players[0].move = movePaper
			case "d":
				m.players[0].move = moveScissors

			// Player 2
			case "j":
				m.players[1].move = moveRock
			case "k":
				m.players[1].move = movePaper
			case "l":
				m.players[1].move = moveScissors
			}
		}
	// Result screen to show users the Results
	// Players can press another key to start a new round
	case stateResult:
		if msg, ok := msg.(tea.KeyMsg); ok && (msg.String() == " " || msg.String() == "enter") {
			m.state = stateLobby
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
