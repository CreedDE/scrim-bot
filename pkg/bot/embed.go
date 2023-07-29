package bot

import "github.com/bwmarrin/discordgo"

func createEmbed(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == s.State.User.ID {
		return
	}

	newEmbed := &discordgo.MessageEmbed{
		Title:       "How to register your Team?",
		Description: "With our Bot u can register your Team and get your own role.\nThe Role helps everyone to know which team u are part of and they can use it to contact your Team.\n\n**Create your Team**\nTo create your Team, u have to enter the command **'/register-team'** and Discord will show you which parameters you need to give the Bot. The Team name is obvious, after that u have to enter your team color. This color will be present for the role and it will be your color in the discord. To help you choose an color, u can use an simple Color Picker from the Internet like this one: https://htmlcolorcodes.com/color-picker/ and after u got your color, u need to copy the 6 characters of the HEX Colorcode on the site. But watchout, the Bot doesnt accept the # character.\n\nHere is an Example how it should look like:\nTeam Name: Sussy Kittens\nColor: 123456\n\n**Adding players to your Team**\nIf are inviting your team members over to this discord, they also should have your team role. And u dont need to wait for an Admin to do it.\nU just type **'/add-players'** first you have to mention the team role where they should get added to and after that u can mention up to 5 players that you wanna add to your team.",
		Color:       16728831,
	}

	if msg.Content == "scregister" {
		s.ChannelMessageSendEmbed(msg.ChannelID, newEmbed)
	}
}
