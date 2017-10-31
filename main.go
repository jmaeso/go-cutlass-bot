package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	url := "https://api.telegram.org/bot" + token + "/getUpdates"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	var updateReceived app.Update

	if err := json.NewDecoder(resp.Body).Decode(&updateReceived); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message received to bot: %+v\n", updateReceived)
}
