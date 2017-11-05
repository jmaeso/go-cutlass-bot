package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	token := os.Getenv("CUTLASS_BOT_TOKEN")
	if token == "" {
		log.Fatal("CUTLASS_BOT_TOKEN env var is required.")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	lastID := -1

	for {
		u := tgbotapi.NewUpdate(lastID + 1)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)
		if err != nil {
			log.Panic(err)
		}

		for update := range updates {
			if update.Message == nil {
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
			lastID++
		}
	}
}
