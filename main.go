package main

import (
	"log"
	"os"
	"scrim-go-bot/pkg/bot"

	"github.com/joho/godotenv"
)

func getEnv(key string) (string, bool) {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		return "", false
	}

	currentKey := os.Getenv(key)

	// Check if the Environment is not empty
	if currentKey != "" {
		return os.Getenv(key), true
	} else {
		return os.Getenv(key), false
	}
}

func main() {
	botToken, ok := getEnv("BOT_TOKEN")

	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	bot.BotToken = botToken
	bot.Run()
}
