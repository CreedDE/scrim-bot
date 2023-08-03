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
	// loading all variables from the config.toml file
	config := bot.LoadConfig()

	// assign config variables
	bot.BotToken = config.Bot_Token
	bot.GUILD_ID = config.Guild_ID
	bot.Forbidden_Roles = config.Forbidden_Roles
	bot.Run()
}
