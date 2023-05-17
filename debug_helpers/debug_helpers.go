package debug_helpers

import (
	"health-check-tui/api_calls"
	"log"
	"math/rand"
	"time"
)

func GetDebugResponses() (map[int]api_calls.HealthState, error) {
	debugResponses := map[int]api_calls.HealthState{
		1: api_calls.Unchecked,
		2: api_calls.Healthy,
		3: api_calls.Healthy,
		4: api_calls.Unhealthy,
		5: api_calls.Unchecked,
		6: api_calls.Inconclusive}

	return debugResponses, nil
}

func GetDebugUptime() (map[int]float64, error) {
	rand.NewSource(time.Now().UnixNano())
	debugUptimes := map[int]float64{}
	for i := 1; i <= 6; i++ {
		result, err := randFloatPresetMinMax()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		debugUptimes[i] = result
	}

	return debugUptimes, nil
}

func randFloatPresetMinMax() (float64, error) {
	return randFloat(97, 100)
}

func randFloat(min, max float64) (float64, error) {
	return min + rand.Float64()*(max-min), nil
}
