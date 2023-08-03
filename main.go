package main

import (
	"scrim-go-bot/pkg/bot"
)

func main() {
	// loading all variables from the config.toml file
	config := bot.LoadConfig()

	// assign config variables
	bot.BotToken = config.Bot_Token
	bot.GUILD_ID = config.Guild_ID
	bot.Forbidden_Roles = config.Forbidden_Roles
	bot.Run()
}
