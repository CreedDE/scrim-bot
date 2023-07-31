package main

import (
	"log"
	"net/http"
	"os"
	"scrim-go-bot/pkg/bot"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	botToken, ok := getEnv("BOT_TOKEN")

	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	guildId, ok := getEnv("DISCORD_GUILD_ID")

	if !ok {
		log.Fatal("Must set Guild ID as env variable: DISCORD_GUILD_ID")
	}

	bot.BotToken = botToken
	bot.GUILD_ID = guildId
	bot.Run()

	http.ListenAndServe(":8080", r)
}
