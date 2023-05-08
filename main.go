package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"telegram_bot/config"
)

const telegramAPI = "https://api.telegram.org/bot%s/%s"

var token = config.GetToken()

type updateResponse struct {
	Ok     bool     `json:"ok"`
	Result []update `json:"result"`
}

type update struct {
	UpdateID      int            `json:"update_id"`
	Message       *message       `json:"message"`
	CallbackQuery *callbackQuery `json:"callback_query"`
}

type message struct {
	MessageID int64  `json:"message_id"`
	From      *user  `json:"from"`
	Chat      *chat  `json:"chat"`
	Text      string `json:"text"`
}

type user struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

type callbackQuery struct {
	ID              string   `json:"id"`
	From            *user    `json:"from"`
	Message         *message `json:"message"`
	Data            string   `json:"data"`
	ChatInstance    string   `json:"chat_instance"`
	InlineMessageID string   `json:"inline_message_id"`
}

type sendMessageRequest struct {
	ChatID      int64               `json:"chat_id"`
	Text        string              `json:"text"`
	ReplyMarkup *inlineKeyboardMark `json:"reply_markup,omitempty"`
}

type inlineKeyboardMark struct {
	InlineKeyboard [][]inlineKeyboardButton `json:"inline_keyboard"`
}

type inlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

func main() {
	lastUpdateID := 0
	for {
		updates, err := getUpdates(lastUpdateID)
		if err != nil {
			log.Println("Error getting updates:", err)
			continue
		}

		for _, update := range updates {
			lastUpdateID = update.UpdateID + 1
			if update.Message != nil && update.Message.Text == "/keyboard" {
				keyboard := [][]inlineKeyboardButton{
					{
						{Text: "кнопка 1", CallbackData: "button1"},
						{Text: "кнопка 2", CallbackData: "button2"},
					},
					{
						{Text: "кнопка 3", CallbackData: "button3"},
						{Text: "кнопка 4", CallbackData: "button4"},
					},
					{
						{Text: "кнопка 5", CallbackData: "button5"},
						{Text: "кнопка 6", CallbackData: "button6"},
					},
					{
						{Text: "кнопка 7", CallbackData: "button7"},
						{Text: "кнопка 8", CallbackData: "button8"},
					},
				}
				replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
				sendMessage(update.Message.Chat.ID, "Выберите опцию:", &replyMarkup)
			} else if update.CallbackQuery != nil {
				// Обрабатываем нажатие кнопки
				switch update.CallbackQuery.Data {
				case "button1":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "Привет 1", CallbackData: "hello1"},
							{Text: "Привет 2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 1. Выберите сообщение:", &replyMarkup)
				case "button2":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "1", CallbackData: "hello1"},
							{Text: "2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 2. Выберите сообщение:", &replyMarkup)
				case "button3":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "1", CallbackData: "hello1"},
							{Text: "2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 3. Выберите сообщение:", &replyMarkup)
				case "button4":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "1", CallbackData: "hello1"},
							{Text: "2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 4. Выберите сообщение:", &replyMarkup)
				case "button5":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "1", CallbackData: "hello1"},
							{Text: "2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 5. Выберите сообщение:", &replyMarkup)
				case "button6":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "1", CallbackData: "hello1"},
							{Text: "2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 6. Выберите сообщение:", &replyMarkup)
				case "button7":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "1", CallbackData: "hello1"},
							{Text: "2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 7. Выберите сообщение:", &replyMarkup)
				case "button8":
					keyboard := [][]inlineKeyboardButton{
						{
							{Text: "1", CallbackData: "hello1"},
							{Text: "2", CallbackData: "hello2"},
						},
					}
					replyMarkup := inlineKeyboardMark{InlineKeyboard: keyboard}
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Вы выбрали кнопку 8. Выберите сообщение:", &replyMarkup)
				case "hello1":
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Привет ты выбрал 1", nil)
				case "hello2":
					deleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
					sendMessage(update.CallbackQuery.Message.Chat.ID, "Привет ты выбрал 2", nil)
				}
			}
		}
	}
}

func getUpdates(offset int) ([]update, error) {
	url := fmt.Sprintf(telegramAPI, token, "getUpdates")
	if offset > 0 {
		url += fmt.Sprintf("?offset=%d", offset)
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var updateResp updateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		return nil, err
	}
	return updateResp.Result, nil
}

func sendMessage(chatID int64, text string, replyMarkup *inlineKeyboardMark) error {
	reqBody := &sendMessageRequest{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	url := fmt.Sprintf(telegramAPI, token, "sendMessage")
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func deleteMessage(chatID int64, messageID int64) error {
	method := fmt.Sprintf("deleteMessage?chat_id=%d&message_id=%d", chatID, messageID)
	url := fmt.Sprintf(telegramAPI, token, method)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
