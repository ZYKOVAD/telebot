package commands

import (
	"fmt"
	"strings"
	"telegram_bot/config"
	"telegram_bot/internal/common"
	"telegram_bot/internal/messages"
	"telegram_bot/internal/models"
	"telegram_bot/internal/weather"
)

var token = config.GetOpenWeatherToken()

type request struct {
	ChatID           int64
	MessageID        int64
	OpenWeatherToken string
	Text             string
}

func Command(update common.Update) {
	if update.Message != nil && update.Message.Text == "/keyboard" {
		regionMenu(update)
	} else if update.Message != nil && (update.Message.Text == "/start" || update.Message.Text == "/help") {
		startedMessage(update)
	} else if update.CallbackQuery != nil {
		downKeyboard(update)
	} else {
		badMessage(update)
	}
}

func regionMenu(update common.Update) {
	keyboard := [][]common.InlineKeyboardButton{
		{
			{Text: "Центральный", CallbackData: "fm+central"},
			{Text: "Северно-Западный", CallbackData: "fm+northwest"},
		},
		{
			{Text: "Южный", CallbackData: "fm+south"},
			{Text: "Северо-Кавказский", CallbackData: "fm+northCaucasus"},
		},
	}
	replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
	messages.SendMessage(update.Message.Chat.ID, "Выберите округ:", &replyMarkup)
}

func downKeyboard(update common.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	messageID := update.CallbackQuery.Message.MessageID
	req := request{}
	req.ChatID = chatID
	req.MessageID = messageID
	req.OpenWeatherToken = token
	req.Text = update.CallbackQuery.Message.Text

	args := strings.Split(update.CallbackQuery.Data, "+")
	if len(args) != 2 {
		return
	}
	state, data := args[0], args[1]

	switch state {
	case "fm":
		cityMenu(req, data)
	case "ct":
		defaultHandler(req, data)
	}
}

func cityMenu(req request, data string) {
	keyboard := models.CityButtonsMap[data]
	replyMarkup := common.InlineKeyboardMark{InlineKeyboard: keyboard}
	messages.DeleteMessage(req.ChatID, req.MessageID)
	msg := fmt.Sprintf("Вы выбрали округ. Выберите город:")
	messages.SendMessage(req.ChatID, msg, &replyMarkup)
}

func defaultHandler(request request, city string) {
	messages.DeleteMessage(request.ChatID, request.MessageID)
	data, err := weather.GetWeather(city, request.OpenWeatherToken)
	if err != nil {
		messages.SendMessage(request.ChatID, err.Error(), nil)
	}
	messages.SendMessage(request.ChatID, data.String(), nil)
}

func startedMessage(update common.Update) {
	message := "Привет\nЧтобы получить клавиатуру введите команду /keyboard"
	messages.SendMessage(update.Message.Chat.ID, message, nil)
}

func badMessage(update common.Update) {
	message := "я тебя не понимаю введи существующую команду /help или /keyboard"
	messages.SendMessage(update.Message.Chat.ID, message, nil)
}
