package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/errors"
	"health-check-tui/api_calls"
	"health-check-tui/theme"
)

func getEndpointBoxStyle(healthState api_calls.HealthState) lipgloss.Style {

	switch healthState {
	case api_calls.Unchecked:
		return theme.GetEndpointUncheckedBoxStyle()
	case api_calls.Healthy:
		return theme.GetEndpointHealthyBoxStyle()
	case api_calls.Unhealthy:
		return theme.GetEndpointUnhealthyBoxStyle()
	case api_calls.Inconclusive:
		return theme.GetEndpointInconclusiveBoxStyle()
	}

	panic(
		errors.New(
			fmt.Sprintf(
				"Unknown Health State '%d'",
				healthState)))
}
