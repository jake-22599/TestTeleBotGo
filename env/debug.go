//go:build DEBUG_BUILD

package env

func GetTelegramBotToken() string {
	return ""
}

func GetTelegramBotDebugFlag() bool {
	return true
}
