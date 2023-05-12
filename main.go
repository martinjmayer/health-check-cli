package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	endpoints     map[int]string
	healthStates  map[int]HealthState
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
		healthStates:  map[int]HealthState{},
		uptimePercent: map[int]float64{},
	}

	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Printf("Failed to start the program: %s\n", err)
	}
}

func (m model) Init() tea.Cmd {
	return GetNewTick()
}

func GetNewTick() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
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
		m.uptimePercent[m.selected] = calculateUptimePercentage()
		return m, tea.Batch(checkEndpoint(m.endpoints[m.selected]), GetNewTick())

	case string:
		m.healthStates[m.selected] = getHealthStateFromString(msg)
		m.uptimePercent[m.selected] = calculateUptimePercentage()
	}
	return m, nil
}

func getHealthStateFromString(healthStateString string) HealthState {
	return getHealthStateMapKeyedOnString()[healthStateString]
}

func getHealthStateText(healthState HealthState) string {
	return getHealthStateMap()[healthState]
}

func getHealthStateMapKeyedOnString() map[string]HealthState {
	hsMap := map[string]HealthState{}
	for key, element := range getHealthStateMap() {
		hsMap[element] = key
	}
	return hsMap
}

func getHealthStateMap() map[HealthState]string {
	return map[HealthState]string{
		Unchecked:    "Unchecked",
		Healthy:      "Healthy",
		Unhealthy:    "Unhealthy",
		Inconclusive: "Inconclusive",
	}
}

func (m model) View() string {
	var lgPanels [6]string

	const titleText = "API Health Check"

	statuses := m.healthStates
	uptimePercent := m.uptimePercent

	statuses = getDebugResponses()
	uptimePercent = getDebugUptime()

	for i := 1; i <= 6; i++ {
		healthState, ok := statuses[i]
		if !ok {
			healthState = Unchecked
		}
		uptime, ok := uptimePercent[i]
		if !ok {
			uptime = 0
		}

		boxStyle := getEndpointBoxStyle(healthState)

		if m.selected == i {
			boxStyle = boxStyle.BorderForeground(lipgloss.Color("205"))
		}

		boxContent := fmt.Sprintf("%s\nStatus: %s\nUptime: %.2f%%", m.endpoints[i], getHealthStateText(healthState), uptime)

		switch healthState {
		case Unchecked:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		case Healthy:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		case Inconclusive:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		case Unhealthy:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		}
	}

	titleStyle := getTitleBoxStyle()

	// Title
	s := titleStyle.Render(titleText)

	// Services - first row
	s += lipgloss.JoinHorizontal(lipgloss.Bottom, lgPanels[0], lgPanels[1], lgPanels[2])

	// Services - second row
	s += lipgloss.JoinHorizontal(lipgloss.Bottom, lgPanels[3], lgPanels[4], lgPanels[5])

	return s
}

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

func getTitleBoxStyle() lipgloss.Style {
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

func getEndpointBoxStyle(healthState HealthState) lipgloss.Style {

	boxStyleHealthStateUnchecked :=
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

	boxStyleHealthStateHealthy :=
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

	boxStyleHealthStateInconclusive :=
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

	boxStyleHealthStateUnhealthy :=
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

	switch healthState {
	case Unchecked:
		return boxStyleHealthStateUnchecked
	case Healthy:
		return boxStyleHealthStateHealthy
	case Inconclusive:
		return boxStyleHealthStateInconclusive
	case Unhealthy:
		return boxStyleHealthStateUnhealthy
	}

	panic(errors.New(fmt.Sprintf("Unknown Health State '%d'", healthState)))
}

func getDebugUptime() map[int]float64 {

	rand.NewSource(time.Now().UnixNano())
	debugUptimes := map[int]float64{
		1: randFloatPresetMinMax(), 2: randFloatPresetMinMax(), 3: randFloatPresetMinMax(),
		4: randFloatPresetMinMax(), 5: randFloatPresetMinMax(), 6: randFloatPresetMinMax(),
	}

	return debugUptimes
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randFloatPresetMinMax() float64 {
	return randFloat(97, 100)
}

func getDebugResponses() map[int]HealthState {
	debugResponses := map[int]HealthState{1: Unchecked, 2: Healthy, 3: Healthy, 4: Unhealthy, 5: Unchecked, 6: Inconclusive}

	return debugResponses
}

type HealthState int

const (
	Unchecked HealthState = iota + 1
	Healthy
	Unhealthy
	Inconclusive
)

func checkEndpoint(endpoint string) tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get(endpoint)

		if err != nil {
			return Unhealthy
		}

		defer closeBody(resp)

		if resp.StatusCode == http.StatusOK {
			return Healthy
		}

		return Unhealthy
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
