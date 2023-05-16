package theme

import "github.com/charmbracelet/lipgloss"

const colourNotChecked = "74"
const colourGreen = "35"
const colourAmber = "214"
const colourRed = "196"

const colourTitle = "63"

const boxPadding = 1
const boxMarginTop = 1
const boxMarginBottom = 1
const boxMarginLeft = 2
const boxMarginRight = 2
const boxWidth = 35

const numColumns = 3

var titleBoxWidth = (boxWidth * numColumns) + ((boxMarginLeft + boxMarginRight) * (numColumns))

func GetTitleBoxStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(colourTitle)).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color(colourTitle)).
		Border(lipgloss.RoundedBorder()).
		Padding(boxPadding).
		MarginTop(boxMarginTop).
		MarginBottom(boxMarginBottom).
		MarginLeft(boxMarginLeft).
		MarginRight(boxMarginRight).
		Align(lipgloss.Center).
		Width(titleBoxWidth)
}

func GetEndpointUncheckedBoxStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(colourNotChecked)).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color(colourNotChecked)).
		Border(lipgloss.RoundedBorder()).
		Padding(boxPadding).
		MarginTop(boxMarginTop).
		MarginBottom(boxMarginBottom).
		MarginLeft(boxMarginLeft).
		MarginRight(boxMarginRight).
		Width(boxWidth)
}

func GetEndpointHealthyBoxStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(colourGreen)).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color(colourGreen)).
		Border(lipgloss.RoundedBorder()).
		Padding(boxPadding).
		MarginTop(boxMarginTop).
		MarginBottom(boxMarginBottom).
		MarginLeft(boxMarginLeft).
		MarginRight(boxMarginRight).
		Width(boxWidth)
}

func GetEndpointInconclusiveBoxStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(colourAmber)).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color(colourAmber)).
		Border(lipgloss.RoundedBorder()).
		Padding(boxPadding).
		MarginTop(boxMarginTop).
		MarginBottom(boxMarginBottom).
		MarginLeft(boxMarginLeft).
		MarginRight(boxMarginRight).
		Width(boxWidth)
}

func GetEndpointUnhealthyBoxStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(colourRed)).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color(colourRed)).
		Border(lipgloss.RoundedBorder()).
		Padding(boxPadding).
		MarginTop(boxMarginTop).
		MarginBottom(boxMarginBottom).
		MarginLeft(boxMarginLeft).
		MarginRight(boxMarginRight).
		Width(boxWidth)
}
