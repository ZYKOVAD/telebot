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

func GetOpenWeatherToken() string {
	token := os.Getenv("OPEN_WEATHER_TOKEN")
	if token == "" {
		log.Fatal("no open weather token in env")
	}
	return token
}
