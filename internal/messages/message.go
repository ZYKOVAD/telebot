package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"telegram_bot/internal/common"
)

func SendMessage(chatID int64, text string, replyMarkup *common.InlineKeyboardMark) error {
	reqBody := &common.SendMessageRequest{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(common.TelegramAPI, common.Token, "sendMessage")
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func DeleteMessage(chatID int64, messageID int64) error {
	method := fmt.Sprintf("deleteMessage?chat_id=%d&message_id=%d", chatID, messageID)
	url := fmt.Sprintf(common.TelegramAPI, common.Token, method)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
