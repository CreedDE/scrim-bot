package bot

import (
	"flag"
	"fmt"
	"log"
	"strconv"

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
					Description: "HEX Color Code for the role. Example: 123456",
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

			var currentColor string
			var teamName string

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// Get the value from the option map.
			// When the option exists, ok = true
			if name, ok := optionMap["team-name"]; ok {
				teamName = name.StringValue()
			}

			if color, ok := optionMap["team-color"]; ok {
				currentColor = color.StringValue()
			}

			// if division, ok := optionMap["premier-division"]; ok {
			// margs = append(margs, division.IntValue())
			// msgformat += "> premier-division: %d\n"
			// }

			color, err := strconv.ParseInt(currentColor, 16, 0)

			if err != nil {
				// we shouldn't stop the Bot if someone make a mistake with the color.
				log.Println(err)

				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: ":x: **Sorry we couldnt create your Team!**\nYou made an mistake with your team color, if u had an color like this: *#123456*, remove the # and try it again!",
					},
				})
			}

			var decimalColor int = int(color)
			var mentionable = true

			newRole := *&discordgo.RoleParams{
				Name:        teamName,
				Color:       &decimalColor,
				Mentionable: &mentionable,
				Hoist:       &mentionable,
			}

			createdRole, err := s.GuildRoleCreate(GUILD_ID, &newRole)

			if err != nil {
				log.Println(err)
			}

			// here adding the created Role to who fired the command /register-team
			s.GuildMemberRoleAdd(GUILD_ID, i.Member.User.ID, createdRole.ID)

			// adding the Team Manager role to the user who created the team
			// TODO: this need to be changed as soon this bot will be used on the real discord
			s.GuildMemberRoleAdd(GUILD_ID, i.Member.User.ID, "1133429330224099440")

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Congratz we created your Team named " + teamName + " :tada:",
				},
			})
		},
		"add-players": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			var selectedRole string

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			users := make([]interface{}, 0, len(options))
			msgformat := "You made changes to your team <@&%s>\nYou added the following persons: "

			if opt, ok := optionMap["team-role"]; ok {
				selectedRole = opt.RoleValue(nil, "").ID
				users = append(users, opt.RoleValue(nil, "").ID)
			}

			if opt, ok := optionMap["player-1"]; ok {
				users = append(users, opt.UserValue(nil).ID)
				msgformat += "<@%s>"
			}

			if opt, ok := optionMap["player-2"]; ok {
				users = append(users, opt.UserValue(nil).ID)
				msgformat += ", <@%s>"
			}

			if opt, ok := optionMap["player-3"]; ok {
				users = append(users, opt.UserValue(nil).ID)
				msgformat += ", <@%s>"
			}

			if opt, ok := optionMap["player-4"]; ok {
				users = append(users, opt.UserValue(nil).ID)
				msgformat += ", <@%s>"
			}

			if opt, ok := optionMap["player-5"]; ok {
				users = append(users, opt.UserValue(nil).ID)
				msgformat += ", <@%s>"
			}

			for _, user := range users {
				stringUser := fmt.Sprint(user)
				s.GuildMemberRoleAdd(GUILD_ID, stringUser, selectedRole)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(msgformat, users...),
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
