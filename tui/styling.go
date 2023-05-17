package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/errors"
	"health-check-tui/api_calls"
	"health-check-tui/theme"
	"log"
)

func getEndpointBoxStyle(healthState api_calls.HealthState) (lipgloss.Style, error) {

	switch healthState {
	case api_calls.Unchecked:
		return theme.GetEndpointUncheckedBoxStyle(), nil
	case api_calls.Healthy:
		return theme.GetEndpointHealthyBoxStyle(), nil
	case api_calls.Unhealthy:
		return theme.GetEndpointUnhealthyBoxStyle(), nil
	case api_calls.Inconclusive:
		return theme.GetEndpointInconclusiveBoxStyle(), nil
	}
	err := errors.New(fmt.Sprintf("Unknown Health State '%d'", healthState))
	log.Fatal(err)
	return theme.GetEndpointUncheckedBoxStyle(), err
}
