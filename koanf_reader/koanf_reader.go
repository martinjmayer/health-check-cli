package koanf_reader

import (
	"fmt"
	"health-check-tui/api_calls"
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type KoanfConfigAndSecretReader struct {
	ConfigFilePath string
}

// Global koanf instance with "." as the key path delimiter.
var k = koanf.New(".")

func (koanfReader KoanfConfigAndSecretReader) ReadStringConfig(key string) (string, error) {
	if err := k.Load(file.Provider(koanfReader.ConfigFilePath), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	return k.String(key), nil
}

func (koanfReader KoanfConfigAndSecretReader) ReadEndpointsConfig() (map[int]api_calls.EndpointConfig, error) {
	if err := k.Load(file.Provider(koanfReader.ConfigFilePath), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	var endpointConfigs []api_calls.EndpointConfig
	if err := k.Unmarshal("Endpoints", &endpointConfigs); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}
	endpointConfigsAsMap := map[int]api_calls.EndpointConfig{}
	for index, endpointConfig := range endpointConfigs {
		endpointConfigsAsMap[index+1] = endpointConfig
	}
	return endpointConfigsAsMap, nil
}

func (koanfReader KoanfConfigAndSecretReader) ReadStringSecret(key string) (string, error) {
	return fmt.Sprintf("secret value for '%s'", key), nil
}
