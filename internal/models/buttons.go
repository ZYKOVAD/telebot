package models

import "telegram_bot/internal/common"

var CityButtonsMap = map[string][][]common.InlineKeyboardButton{
	"northwest": {
		{
			{Text: "Санкт-Петербург", CallbackData: "ct+sankt-peterburg"},
			{Text: "Калининград", CallbackData: "ct+kaliningrad"},
		},
	},
	"central": {
		{
			{Text: "Москва", CallbackData: "ct+moscow"},
			{Text: "Воронеж", CallbackData: "ct+voronezh"},
		},
	},
	"south": {
		{
			{Text: "Краснодар", CallbackData: "ct+krasnodar"},
			{Text: "Астрахань", CallbackData: "ct+astrakhan"},
		},
	},
	"northCaucasus": {
		{
			{Text: "Махачкала", CallbackData: "ct+makhachkala"},
			{Text: "Ставрополь", CallbackData: "ct+stavropol"},
		},
	},
}
