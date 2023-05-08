package commands

import (
	"telegram_bot/internal/common"
	"telegram_bot/internal/messages"
)

func Command(update common.Update) {
	if update.Message != nil && update.Message.Text == "/keyboard" {
		firstMenu(update)
	} else if update.Message.Text == "/start" || update.Message.Text == "/help" {
		startedMessage(update)
	} else if update.CallbackQuery != nil {
		downKeyboard(update)
	} else {
		badMessage(update)
	}
}

func firstMenu(update common.Update) {
	keyboard := [][]common.InlineKeyboardButton{
		{
			{Text: "кнопка 1", CallbackData: "button1"},
			{Text: "кнопка 2", CallbackData: "button2"},
		},
		{
			{Text: "кнопка 3", CallbackData: "button3"},
			{Text: "кнопка 4", CallbackData: "button4"},
		},
	}
	replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
	messages.SendMessage(update.Message.Chat.ID, "Выберите опцию:", &replyMarkup)
}

func downKeyboard(update common.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	messageID := update.CallbackQuery.Message.MessageID
	// Обрабатываем нажатие кнопки
	switch update.CallbackQuery.Data {
	case "button1":
		secondMenu(chatID, messageID, "Вы выбрали кнопку 1. Выберите сообщение:")
	case "button2":
		secondMenu(chatID, messageID, "Вы выбрали кнопку 2. Выберите сообщение:")
	case "button3":
		secondMenu(chatID, messageID, "Вы выбрали кнопку 3. Выберите сообщение:")
	case "button4":
		secondMenu(chatID, messageID, "Вы выбрали кнопку 4. Выберите сообщение:")
	case "hello1":
		messages.DeleteMessage(chatID, messageID)
		messages.SendMessage(chatID, "Привет ты выбрал 1", nil)
	case "hello2":
		messages.DeleteMessage(chatID, messageID)
		messages.SendMessage(chatID, "Привет ты выбрал 2", nil)
	}
}

func secondMenu(chatID int64, messageID int64, sendMessage string) {
	keyboard := [][]common.InlineKeyboardButton{
		{
			{Text: "1", CallbackData: "hello1"},
			{Text: "2", CallbackData: "hello2"},
		},
	}
	replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
	messages.DeleteMessage(chatID, messageID)
	messages.SendMessage(chatID, sendMessage, &replyMarkup)
}

func startedMessage(update common.Update) {
	message := "Привет\nЧтобы получить клавиатуру введите команду /keyboard"
	messages.SendMessage(update.Message.Chat.ID, message, nil)
}

func badMessage(update common.Update) {
	message := "я тебя не понимаю введи существующую команду /help или /keyboard"
	messages.SendMessage(update.Message.Chat.ID, message, nil)
}
