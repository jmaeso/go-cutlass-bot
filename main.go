package main

import (
	"log"

	"github.com/jmaeso/go-cutlass-bot/app"
	"github.com/jmaeso/go-cutlass-bot/tools/yaml"
	"gopkg.in/telegram-bot-api.v4"
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

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
