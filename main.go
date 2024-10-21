package main

import (
	"log"
	"strings"

	// third-party libraries
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	// internal
	"tester/telegram/env"
	"tester/telegram/tools"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	telegramBotToken := env.GetTelegramBotToken()
	log.Printf("TELEGRAM_BOT_TOKEN %s", telegramBotToken)

	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	// in order to get more information about the requests being sent to Telegram
	bot.Debug = env.GetTelegramBotDebugFlag()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	u := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	u.Timeout = 60

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(u)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {

		// split the message text message based on white space delimiter
		texts := strings.Fields(update.Message.Text)

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."

		case "open":
			msg.Text = "Opening a keyboard input"
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.Text = "Closing a keyboard input"
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "test_upload_doc":
			filepath := texts[1]
			_, err := tools.SendDocument(bot, update.Message.Chat.ID, filepath)
			if err != nil {
				log.Panic(err)
				continue
			}

		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	log.Printf("Application shutdown gracefully")
} // func main
