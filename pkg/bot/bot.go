package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken        string
	GUILD_ID        string
	Forbidden_Roles []string
)

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)

	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { log.Println("Bot is up!") })

	discord.AddHandler(handlePing)
	discord.AddHandler(initCommands)
	discord.AddHandler(createEmbed)

	createCommands(discord)

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer discord.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-sc
	log.Println("Gracefully shutting down")
	if *RemoveCommands {
		removeCommands(discord)
	}
}

// this function sends an Pong back to if a user types scping
func handlePing(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if message.Content == "scping" {
		discord.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
