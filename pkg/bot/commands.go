package bot

import (
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "register-team",
			Description: "Register your Valorant Team",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "team-name",
					Description: "Name of your Valorant Team",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "team-color",
					Description: "HEX Color Code for the role. Example: #123456",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "premier-division",
					Description: "Type in the Division your team get placed in the Premier Mode (only numbers allowed)",
					Required:    false,
				},
			},
		},
		{
			Name:        "add-players",
			Description: "Add the players to your Team",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "team-role",
					Description: "Select the Team where the players should be assigned to",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionMentionable,
					Name:        "player-1",
					Description: "Add the first player",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionMentionable,
					Name:        "player-2",
					Description: "Add the second player",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionMentionable,
					Name:        "player-3",
					Description: "Add the third player",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionMentionable,
					Name:        "player-4",
					Description: "Add the fourth player",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionMentionable,
					Name:        "player-5",
					Description: "Add the fith player",
					Required:    false,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"register-team": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// This example stores the provided arguments in an []interface{}
			// which will be used to format the bot's response
			margs := make([]interface{}, 0, len(options))
			msgformat := "You learned how to use command options! " +
				"Take a look at the value(s) you entered:\n"

			// Get the value from the option map.
			// When the option exists, ok = true
			if option, ok := optionMap["team-name"]; ok {
				// Option values must be type asserted from interface{}.
				// Discordgo provides utility functions to make this simple.
				margs = append(margs, option.StringValue())
				msgformat += "> team-name: %s\n"
			}

			if opt, ok := optionMap["team-color"]; ok {
				margs = append(margs, opt.StringValue())
				msgformat += "> team-color: %s\n"
			}

			if opt, ok := optionMap["premier-division"]; ok {
				margs = append(margs, opt.IntValue())
				msgformat += "> premier-division: %d\n"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
		"add-players": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// This example stores the provided arguments in an []interface{}
			// which will be used to format the bot's response
			margs := make([]interface{}, 0, len(options))
			msgformat := "You learned how to use command options! " +
				"Take a look at the value(s) you entered:\n"

			// Get the value from the option map.
			// When the option exists, ok = true
			if option, ok := optionMap["team-role"]; ok {
				// Option values must be type asserted from interface{}.
				// Discordgo provides utility functions to make this simple.
				margs = append(margs, option.RoleValue(nil, "").ID)
				msgformat += "> team-role: %s\n"
			}

			if opt, ok := optionMap["player-1"]; ok {
				margs = append(margs, opt.UserValue(nil).ID)
				msgformat += "> player-1: %s\n"
			}

			if opt, ok := optionMap["player-2"]; ok {
				margs = append(margs, opt.UserValue(nil).ID)
				msgformat += "> player-2: %s\n"
			}

			if opt, ok := optionMap["player-3"]; ok {
				margs = append(margs, opt.UserValue(nil).ID)
				msgformat += "> player-3: %s\n"
			}

			if opt, ok := optionMap["player-4"]; ok {
				margs = append(margs, opt.UserValue(nil).ID)
				msgformat += "> player-4: %s\n"
			}

			if opt, ok := optionMap["player-5"]; ok {
				margs = append(margs, opt.UserValue(nil).ID)
				msgformat += "> player-5: %s\n"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
	}

	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
)

func initCommands(session *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(session, i)
	}
}

func createCommands(session *discordgo.Session) {
	log.Println("Adding commands...")
	err := session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for i, v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, "", v)
		log.Println("Adding Command", v.Name)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer session.Close()

}

func removeCommands(session *discordgo.Session) {
	log.Println("Removing Commands...")
	err := session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	for _, v := range registeredCommands {
		err := session.ApplicationCommandDelete(session.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	defer session.Close()
}
