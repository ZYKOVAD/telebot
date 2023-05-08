package menu

import (
	"telegram_bot/internal/common"
	"telegram_bot/internal/messages"
)

func GetMenu(update common.Update) {
	if update.Message != nil && update.Message.Text == "/keyboard" {
		keyboard := [][]common.InlineKeyboardButton{
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
		replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
		messages.SendMessage(update.Message.Chat.ID, "Выберите опцию:", &replyMarkup)
	} else if update.CallbackQuery != nil {
		chatID := update.CallbackQuery.Message.Chat.ID
		messageID := update.CallbackQuery.Message.MessageID
		// Обрабатываем нажатие кнопки
		switch update.CallbackQuery.Data {
		case "button1":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "Привет 1", CallbackData: "hello1"},
					{Text: "Привет 2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 1. Выберите сообщение:", &replyMarkup)
		case "button2":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "1", CallbackData: "hello1"},
					{Text: "2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 2. Выберите сообщение:", &replyMarkup)
		case "button3":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "1", CallbackData: "hello1"},
					{Text: "2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 3. Выберите сообщение:", &replyMarkup)
		case "button4":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "1", CallbackData: "hello1"},
					{Text: "2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 4. Выберите сообщение:", &replyMarkup)
		case "button5":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "1", CallbackData: "hello1"},
					{Text: "2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 5. Выберите сообщение:", &replyMarkup)
		case "button6":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "1", CallbackData: "hello1"},
					{Text: "2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 6. Выберите сообщение:", &replyMarkup)
		case "button7":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "1", CallbackData: "hello1"},
					{Text: "2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 7. Выберите сообщение:", &replyMarkup)
		case "button8":
			keyboard := [][]common.InlineKeyboardButton{
				{
					{Text: "1", CallbackData: "hello1"},
					{Text: "2", CallbackData: "hello2"},
				},
			}
			replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Вы выбрали кнопку 8. Выберите сообщение:", &replyMarkup)
		case "hello1":
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Привет ты выбрал 1", nil)
		case "hello2":
			messages.DeleteMessage(chatID, messageID)
			messages.SendMessage(chatID, "Привет ты выбрал 2", nil)
		}
	}
}
