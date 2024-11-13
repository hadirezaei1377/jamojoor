package telebut

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true
	// log.Printf("Logged in as %s", bot.Self.UserName)

	// recieve new messages from bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil && update.Message.Text != "" {
			userMessage := update.Message.Text

			// create a response
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, userMessage)

			if _, err := bot.Send(msg); err != nil {
				log.Println("Failed to send message:", err)
			}
		}
	}
}
