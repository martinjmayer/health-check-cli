package config_and_secrets

import "health-check-tui/api_calls"

type ConfReader interface {
	ReadStringConfig(key string) (string, error)
	ReadEndpointsConfig() (map[int]api_calls.EndpointConfig, error)
}

type SecretReader interface {
	ReadStringSecret(key string) (string, error)
}
