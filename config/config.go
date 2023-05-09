package config

import (
	"log"
	"os"
)

func GetToken() string {
	token := os.Getenv("API_TOKEN")
	if token == "" {
		log.Fatal("no api token in env")
	}
	return token
}
