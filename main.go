package main

import (
	"health-check-tui/koanf_reader"
	"health-check-tui/tui"
	"log"
)

func main() {

	koanfReader := koanf_reader.KoanfConfigAndSecretReader{
		ConfigFilePath: "app_config.yml",
	}

	err := tui.InitialiseBubbleTea(koanfReader, koanfReader)
	if err != nil {
		log.Fatal(err)
	}
}
