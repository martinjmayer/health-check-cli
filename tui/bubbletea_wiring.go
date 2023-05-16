package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"health-check-tui/api_calls"
	"health-check-tui/debug_helpers"
	"health-check-tui/theme"
	"net/http"
	"os"
	"time"
)

type model struct {
	endpointConfigs map[int]api_calls.EndpointConfig
	healthStates    map[int]api_calls.HealthState
	uptimePercent   map[int]float64
	selected        int
}

type tickMsg time.Time

func InitialiseBubbleTea(configFile string) {
	m := model{
		// replace this with content from the config file
		endpointConfigs: map[int]api_calls.EndpointConfig{
			1: {
				Url:                "http://example1.com",
				ApiKeySecretId:     "api1ApiKey",
				ApiKeyRequestParam: "api-key"},
			2: {
				Url:                "http://example2.com",
				ApiKeySecretId:     "api2ApiKey",
				ApiKeyRequestParam: "api-key"},
			3: {
				Url:                "http://example3.com",
				ApiKeySecretId:     "api3ApiKey",
				ApiKeyRequestParam: "api-key"},
			4: {
				Url:                "http://example4.com",
				ApiKeySecretId:     "api4ApiKey",
				ApiKeyRequestParam: "api-key"},
			5: {
				Url:                "http://example5.com",
				ApiKeySecretId:     "api5ApiKey",
				ApiKeyRequestParam: "api-key"},
			6: {
				Url:                "http://example6.com",
				ApiKeySecretId:     "api6ApiKey",
				ApiKeyRequestParam: "api-key"},
		},
		healthStates:  map[int]api_calls.HealthState{},
		uptimePercent: map[int]float64{},
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to start the program: %v\n", err)
		os.Exit(1)
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
			return m, checkEndpoint(m.endpointConfigs[m.selected])
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
		return m, tea.Batch(checkEndpoint(m.endpointConfigs[m.selected]), GetNewTick())

	case string:
		m.healthStates[m.selected] = api_calls.GetHealthStateFromString(msg)
		m.uptimePercent[m.selected] = calculateUptimePercentage()
	}
	return m, nil
}

func (m model) View() string {
	var lgPanels [6]string
	const titleText = "API Health Check"
	statuses := debug_helpers.GetDebugResponses()   // m.healthStates
	uptimePercent := debug_helpers.GetDebugUptime() //m.uptimePercent

	for i := 1; i <= 6; i++ {
		healthState, ok := statuses[i]
		if !ok {
			healthState = api_calls.Unchecked
		}

		uptime, ok := uptimePercent[i]
		if !ok {
			uptime = 0
		}

		boxStyle := getEndpointBoxStyle(healthState)
		if m.selected == i {
			boxStyle = boxStyle.BorderForeground(lipgloss.Color("205"))
		}

		boxContent := fmt.Sprintf("%s\nStatus: %s\nUptime: %.2f%%", m.endpointConfigs[i].Url, api_calls.GetHealthStateText(healthState), uptime)

		switch healthState {
		case api_calls.Unchecked:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		case api_calls.Healthy:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		case api_calls.Inconclusive:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		case api_calls.Unhealthy:
			lgPanels[i-1] = boxStyle.Render(boxContent)
		}
	}

	titleStyle := theme.GetTitleBoxStyle()
	s := titleStyle.Render(titleText)                                                    // Title
	s += lipgloss.JoinHorizontal(lipgloss.Bottom, lgPanels[0], lgPanels[1], lgPanels[2]) // Services - first row
	s += lipgloss.JoinHorizontal(lipgloss.Bottom, lgPanels[3], lgPanels[4], lgPanels[5]) // Services - second row
	return s
}

func checkEndpoint(endpoint api_calls.EndpointConfig) tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get(endpoint.Url)
		if err != nil {
			return api_calls.Unhealthy
		}
		defer closeBody(resp)
		if resp.StatusCode == http.StatusOK {
			return api_calls.Healthy
		}
		return api_calls.Unhealthy
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
