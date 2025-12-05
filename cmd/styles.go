package main

import "github.com/charmbracelet/lipgloss"

// Styles
const (
	primaryColor   = "#B8BB26"   // A nice lime green
	secondaryColor = "#D3869B"   // A soft purple
	subtleColor    = "#504945"   // Dark grey
	dangerColor    = "#cc4949ff" //Danger color
)

var (
	// Style for the "Current Count" label
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(secondaryColor)).
			Bold(true).
			MarginBottom(1) // Add space below the label

	// Style for the actual number
	counterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(primaryColor)).
			Background(lipgloss.Color(subtleColor)).
			Padding(1, 3). // Top/Bottom: 1, Left/Right: 3
			Align(lipgloss.Center).
			Bold(true)

	// Style for the help text at the bottom
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")). // ANSI 256 color code for grey
			Italic(true).
			MarginTop(2) // Add space above the help text

	// Style for the help text at the bottom
	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dangerColor)). // ANSI 256 color code for grey
			Italic(true).
			MarginTop(2) // Add space above the help text

	selectedContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(primaryColor))
	containerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))
)
