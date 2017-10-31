package main

import (
	"fmt"
	"log"

	"github.com/jmaeso/go-cutlass-bot/app"
	"github.com/jmaeso/go-cutlass-bot/tools/yaml"
)

func main() {
	var settings app.Settings
	if err := yaml.Load("settings.yml", &settings); err != nil {
		log.Fatalf("could not load config. err: %s\n", err.Error())
	}

	token := settings.Token
	if token == "" {
		log.Fatalf("Token required in settings.yml")
	}

	fmt.Printf("Token: %s\n", token)
}
