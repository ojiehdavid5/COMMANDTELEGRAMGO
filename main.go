package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
// var numericKeyboard = tgbotapi.NewReplyKeyboard(
//     tgbotapi.NewKeyboardButtonRow(
//         tgbotapi.NewKeyboardButton("1"),
//         tgbotapi.NewKeyboardButton("2"),
//         tgbotapi.NewKeyboardButton("3"),
//     ),
// )
func test() string {
	return "Upload a picture to extract the text"
}


var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
        tgbotapi.NewInlineKeyboardButtonData("2", "2"),
        tgbotapi.NewInlineKeyboardButtonData("3", "3"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("4", "4"),
        tgbotapi.NewInlineKeyboardButtonData("5", "5"),
        tgbotapi.NewInlineKeyboardButtonData("6", "6"),
    ),
)


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
    if update.Message != nil {
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

        switch update.Message.Text {
        case "Open":
            msg.ReplyMarkup = numericKeyboard
        }

        // Send the message.
        if _, err = bot.Send(msg); err != nil {
            panic(err)
        }
    }

    if update.CallbackQuery != nil {
        // Respond to the callback query, telling Telegram to show the user
        // a message with the data received.
        callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
        if _, err := bot.Request(callback); err != nil {
            panic(err)
        }

        // And finally, send a message containing the data received.
        msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
        if _, err := bot.Send(msg); err != nil {
            panic(err)
        }
    }
}


}