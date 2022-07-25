package telegram

import (
	"bytes"
	"fmt"
	"net/http"
	"news-bot/logger"
)

const (
	chatID       = "Chat ID"
	BOT_TOKEN    = "TOKEN"
	TELEGRAM_URL = "https://api.telegram.org/bot"
)

func SendMessage(text string) {
	data := []byte(fmt.Sprintf(`{"chat_id":%s, "text":"%s", "parse_mode":"HTML", "disable_web_page_preview": true}`, chatID, text))
	tx := bytes.NewReader(data)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TELEGRAM_URL, BOT_TOKEN), "application/json", tx)
	logger.ForError(err)
}
