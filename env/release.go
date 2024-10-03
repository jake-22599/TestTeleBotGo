//go:build !DEBUG_BUILD

package env

import (
	"os"
)

const KEY_TEL_BOT_TOKEN string = "TELEGRAM_BOT_TOKEN"

func GetTelegramBotToken() string {
	return os.Getenv(KEY_TEL_BOT_TOKEN)
}

func GetTelegramBotDebugFlag() bool {
	return false
}
