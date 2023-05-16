package api_calls

type HealthState int

const (
	Unchecked HealthState = iota + 1
	Healthy
	Unhealthy
	Inconclusive
)

func GetHealthStateFromString(healthStateString string) HealthState {
	return getHealthStateMapKeyedOnString()[healthStateString]
}

func GetHealthStateText(healthState HealthState) string {
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
