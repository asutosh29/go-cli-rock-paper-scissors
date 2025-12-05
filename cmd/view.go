package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	// GAME UI
	header := labelStyle.Render("=== ROCK PAPER SCISSORS ARENA ===")
	var footer string
	var help string
	help = helpStyle.Render("• Press 'q' to exit \n• Press 'e' to rename \n• Press 'tab' to switch")

	// PLAYERS INFO
	playerViews := []string{}
	for i, player := range m.players {
		var moveName move
		var controlHints string

		var moveDisplay string
		if i == 0 {
			moveName = m.players[0].move
			controlHints = "• A(Rock) \n• S(Paper) \n• D(Scissors)"
		} else {
			moveName = m.players[1].move
			controlHints = "• J(Rock) \n• K(Paper) \n• L(Scissors)"
		}

		switch m.state {
		case stateLobby:
			moveDisplay = "Waiting..."
			// RENAME BAR
			if m.isEditing {
				footer = footerStyle.Render(m.textInput.View())
				help = helpStyle.Render("• Press 'esc' to exit \n• Press 'enter' to accept ")
			} else {
				help = helpStyle.Render("• Press 'q' to exit \n• Press 'e' to rename \n• Press 'tab' to switch")
			}
		case stateCountdown:
			if moveName != moveNone {
				// If they picked something, show they are ready
				moveDisplay = lipgloss.NewStyle().
					Foreground(lipgloss.Color(primaryColor)).
					Bold(true).
					Render("Selected")
			} else {
				moveDisplay = "Thinking..."
			}
			help = helpStyle.Render("• Press 'q' to exit")
		case stateResult:
			// Show moves
			moveDisplay = lipgloss.NewStyle().
				Foreground(lipgloss.Color(secondaryColor)).
				Bold(true).
				Render(moveName.String())
			help = helpStyle.Render("• Press 'q' to exit\n• Press 'space' or 'enter' to go the Lobby!")
		}
		name := labelStyle.Render(player.name)
		score := counterStyle.Render(fmt.Sprintf("Score: %d", player.wins))

		innerBox := lipgloss.JoinVertical(lipgloss.Center, name, score, "\n", moveDisplay)
		playerBox := ""
		if i == m.selectedPlayer {
			playerBox = playerBoxStyle.BorderForeground(lipgloss.Color(primaryColor)).Render(innerBox)
		} else {
			playerBox = playerBoxStyle.Render(innerBox)

		}

		hint := helpStyle.Render(controlHints)
		fullColumn := lipgloss.JoinVertical(lipgloss.Center, playerBox, hint)
		playerViews = append(playerViews, fullColumn)
	}

	gameMessage := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#D3869B")).
		Padding(1, 6).
		Bold(true).
		Margin(1).
		Render(m.message)

	playersDisplay := lipgloss.JoinHorizontal(lipgloss.Center, playerViews...)

	logView := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(subtleColor)).
		MarginTop(1).
		Padding(0, 1).
		Render(m.viewPort.View())

	var ui string
	// FINAL UI RENDER
	if footer != "" {
		ui = lipgloss.JoinVertical(lipgloss.Center,
			header,
			gameMessage,
			playersDisplay,
			logView,
			footer,
			help,
		)
	} else {
		ui = lipgloss.JoinVertical(lipgloss.Center,
			header,
			gameMessage,
			playersDisplay,
			logView,
			help,
		)
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(primaryColor)).
		Padding(0, 2).
		Render(ui)
}
