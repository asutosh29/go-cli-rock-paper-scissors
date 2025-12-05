package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type player struct {
	counter int
	name    string
}

type model struct {
	players       []player
	message       string
	currentPlayer int // will 0 or 1

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

	vp := viewport.New(30, 3)

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

		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.message = "Quitting thanks for playing!"
			return m, tea.Quit
		case "e":
			m.isEditing = true
			m.textInput.Focus()
			m.textInput.SetValue(m.players[m.currentPlayer].name)

			m.message = "Editting name... press esc to exit!"
			return m, textinput.Blink

		case "up", "j":
			if m.players[m.currentPlayer].counter >= 10 {
				m.message = "Can't go above 10..."
				return m, nil
			}
			m.message = ""
			m.players[m.currentPlayer].counter++
		case "down", "k":
			if m.players[m.currentPlayer].counter <= 0 {
				m.message = "Can't go below 0..."
				return m, nil
			}
			m.message = ""
			m.players[m.currentPlayer].counter--
		case "tab":
			m.currentPlayer = (m.currentPlayer + 1) % len(m.players)
		}
	}
	return m, nil
}
func (m model) View() string {

	label := labelStyle.Render("=== Counter Game ===")

	playersRow := []string{}
	for i, player := range m.players {
		name := labelStyle.Render(player.name)
		counter := counterStyle.Render(fmt.Sprintf("Score: %d", player.counter))

		container := lipgloss.JoinVertical(lipgloss.Center, name, counter)
		if m.currentPlayer == i {
			container = selectedContainerStyle.Render(container)
		} else {
			container = containerStyle.Render(container)
		}
		playersRow = append(playersRow, container)
	}
	playersRowString := lipgloss.JoinHorizontal(lipgloss.Center, playersRow...)

	var footer string
	if m.isEditing {
		footer = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(secondaryColor)).
			Padding(0, 1).
			Render(m.textInput.View())
	} else {
		footer = helpStyle.Render("• Press up/down or j/k to increase/decrease the counter \n• Press 'e' to rename \n• Tab to switch")
	}

	m.viewPort.SetContent(messageStyle.Render(m.message))
	viewPortText := m.viewPort.View()
	ui := lipgloss.JoinVertical(lipgloss.Center, label, playersRowString, footer, viewPortText)

	container := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()). // Nice rounded corners
		BorderForeground(lipgloss.Color(primaryColor)).
		Padding(1, 2). // Add breathing room inside the border
		Render(ui)
	return container
}

func main() {
	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
