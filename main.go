package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var numericKeyboard = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("1"),
        tgbotapi.NewKeyboardButton("2"),
        tgbotapi.NewKeyboardButton("3"),
    ),
)
func test() string {
	return "Upload a picture to extract the text"
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the bot token from the environment variable
	token := os.Getenv("TELEGRAM_APITOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_APITOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
        if update.Message == nil { // ignore any non-Message updates
            continue
        }
msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

        // Extract the command from the Message.
switch update.Message.Text {
        case "Open":
            msg.ReplyMarkup = numericKeyboard
		case "1", "2" :
            msg.Text = "You selected " + update.Message.Text

		case "3":
			msg.Text = test()

					case "Cancel":
            msg.Text = "Cancelled"
            msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)	
        case "Close":
            msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
        }

        if _, err := bot.Send(msg); err != nil {
            log.Panic(err)
        }
    }

	// Add your bot logic here
}