package main

import "github.com/charmbracelet/lipgloss"

// Styles
const (
	primaryColor   = "#B8BB26" // A nice lime green
	secondaryColor = "#D3869B" // A soft purple
	subtleColor    = "#504945" // Dark grey
	// dangerColor    = "#cc4949ff" //Danger color
)

var (
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(secondaryColor)).
			Bold(true).
			MarginBottom(1)

	counterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(primaryColor)).
			Background(lipgloss.Color(subtleColor)).
			Padding(1, 3).
			Align(lipgloss.Center).
			Bold(true)

	playerBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(subtleColor)).
			Padding(0, 4).
			Width(30).
			Align(lipgloss.Center)

	footerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(secondaryColor)).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Italic(true).
			MarginTop(1)
)
