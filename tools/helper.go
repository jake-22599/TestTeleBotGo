package tools

import (
	"fmt"
	"os"

	// third-party libraries
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) (tgbotapi.Message, error) {
	// Create a new MessageConfig.
	msg := tgbotapi.NewMessage(chatID, text)
	msgResp, err := bot.Send(msg)
	return msgResp, err
}

func SendDocument(bot *tgbotapi.BotAPI, chatID int64, filePath string) (tgbotapi.Message, error) {

	fileData, errReadFile := os.ReadFile(filePath)
	if errReadFile != nil {
		fmt.Printf("error read file, %v\n", errReadFile)
		return tgbotapi.Message{}, errReadFile
	}

	fileDataRequest := &tgbotapi.FileBytes{
		Name:  filePath,
		Bytes: fileData,
	}

	msg := tgbotapi.NewDocument(chatID, fileDataRequest)
	msgResp, err := bot.Send(msg)
	return msgResp, err
}
