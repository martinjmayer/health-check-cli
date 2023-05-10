package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	endpoints     map[int]string
	statuses      map[int]string
	uptimePercent map[int]float64
	selected      int
}

type tickMsg time.Time

func main() {
	m := model{
		endpoints: map[int]string{
			1: "http://example1.com",
			2: "http://example2.com",
			3: "http://example3.com",
			4: "http://example4.com",
			5: "http://example5.com",
			6: "http://example6.com",
		},
		statuses:      make(map[int]string),
		uptimePercent: make(map[int]float64),
	}

	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Printf("Failed to start the program: %s\n", err)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(time.Second*15, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			return m, checkEndpoint(m.endpoints[m.selected])
		}
		if msg.Type == tea.KeyUp || msg.String() == "k" {
			if m.selected > 1 {
				m.selected--
			}
		}
		if msg.Type == tea.KeyDown || msg.String() == "j" {
			if m.selected < 6 {
				m.selected++
			}
		}

	case tickMsg:
		return m, checkEndpoint(m.endpoints[m.selected])

	case string:
		m.statuses[m.selected] = msg
		m.uptimePercent[m.selected] = calculateUptimePercentage() // Replace with your actual calculation
	}
	return m, nil
}

func (m model) View() string {
	var lgPanels [6]string

	const titleText = "API Health Check"

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

	statuses := m.statuses
	uptimePercent := m.uptimePercent

	statuses = getDebugResponses()
	uptimePercent = getDebugUptime()

	for i := 1; i <= 6; i++ {
		status, ok := statuses[i]
		if !ok {
			status = "not checked"
		}
		uptime, ok := uptimePercent[i]
		if !ok {
			uptime = 0
		}

		boxStyleNotChecked :=
			lipgloss.NewStyle().
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

		boxStyleGreen :=
			lipgloss.NewStyle().
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

		boxStyleAmber :=
			lipgloss.NewStyle().
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

		boxStyleRed :=
			lipgloss.NewStyle().
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

		if m.selected == i {
			boxStyleNotChecked = boxStyleNotChecked.BorderForeground(lipgloss.Color("205"))
		}

		boxContent := fmt.Sprintf("%s\nStatus: %s\nUptime: %.2f%%", m.endpoints[i], status, uptime)
		switch status {
		case unchecked:
			lgPanels[i-1] = boxStyleNotChecked.Render(boxContent)
		case healthy:
			lgPanels[i-1] = boxStyleGreen.Render(boxContent)
		case inconclusive:
			lgPanels[i-1] = boxStyleAmber.Render(boxContent)
		case unhealthy:
			lgPanels[i-1] = boxStyleRed.Render(boxContent)
		}
	}

	titleBoxWidth := (boxWidth * numColumns) + ((boxMarginLeft + boxMarginRight) * (numColumns))
	boxStyleTitle :=
		lipgloss.NewStyle().
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

	title := boxStyleTitle

	// Title
	s := title.Render(titleText)

	// Services - first row
	s += lipgloss.JoinHorizontal(lipgloss.Bottom, lgPanels[0], lgPanels[1], lgPanels[2])

	// Services - second row
	s += lipgloss.JoinHorizontal(lipgloss.Bottom, lgPanels[3], lgPanels[4], lgPanels[5])

	return s
}

func getDebugUptime() map[int]float64 {

	rand.NewSource(time.Now().UnixNano())
	debugUptimes := map[int]float64{1: 97, 2: 99.9, 3: 100, 4: 94.2, 5: 98, 6: 99.25236227}

	return debugUptimes
}

/*
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64() * (max - min)
	}
	return res
}
*/

func getDebugResponses() map[int]string {
	debugResponses := map[int]string{1: unchecked, 2: healthy, 3: healthy, 4: unhealthy, 5: unchecked, 6: inconclusive}

	return debugResponses
}

const unchecked = "unchecked"
const healthy = "healthy"
const unhealthy = "unhealthy"
const inconclusive = "inconclusive"

func checkEndpoint(endpoint string) tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get(endpoint)

		if err != nil {
			return unhealthy
		}

		defer closeBody(resp)

		if resp.StatusCode == http.StatusOK {
			return healthy
		}

		return unhealthy
	}
}

func closeBody(resp *http.Response) {
	if resp == nil {
		return
	}

	if resp.Body == nil {
		return
	}

	resp.Body.Close()
}

func calculateUptimePercentage() float64 {
	// Replace this with the actual uptime calculation
	return 100.0
}
