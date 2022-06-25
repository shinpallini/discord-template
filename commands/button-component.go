package commands

import "github.com/bwmarrin/discordgo"

func init() {
	addComponent("test", func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Button is clicked!",
			},
		})
	})
}
