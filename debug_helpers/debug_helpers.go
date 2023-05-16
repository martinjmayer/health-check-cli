package debug_helpers

import (
	"health-check-tui/api_calls"
	"math/rand"
	"time"
)

func GetDebugResponses() map[int]api_calls.HealthState {
	debugResponses := map[int]api_calls.HealthState{
		1: api_calls.Unchecked,
		2: api_calls.Healthy,
		3: api_calls.Healthy,
		4: api_calls.Unhealthy,
		5: api_calls.Unchecked,
		6: api_calls.Inconclusive}

	return debugResponses
}

func GetDebugUptime() map[int]float64 {

	rand.NewSource(time.Now().UnixNano())
	debugUptimes := map[int]float64{
		1: randFloatPresetMinMax(), 2: randFloatPresetMinMax(), 3: randFloatPresetMinMax(),
		4: randFloatPresetMinMax(), 5: randFloatPresetMinMax(), 6: randFloatPresetMinMax(),
	}

	return debugUptimes
}

func randFloatPresetMinMax() float64 {
	return randFloat(97, 100)
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
