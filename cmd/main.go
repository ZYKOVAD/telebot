package main

import (
	"log"
	"telegram_bot/internal/commands"
	"telegram_bot/internal/updates"
)

func main() {
	lastUpdateID := 0
	for {
		thisUpdates, err := updates.GetUpdates(lastUpdateID)
		if err != nil {
			log.Println("Error getting updates:", err)
			continue
		}
		for _, update := range thisUpdates {
			lastUpdateID = update.UpdateID + 1
			commands.Command(update)
		}
	}
}
