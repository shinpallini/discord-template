package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	// Nameで定義された文字列がKeyになるので同時に書ける
	embed := NewMessageEmbed(
		SetEmbedType(discordgo.EmbedTypeRich),
		SetTitle("Embed!"),
		SetDescription("Description!"),
		SetColor(0x15e81c),
	)
	responseData := NewInteractionResponseData(
		SetContent("This is a basic-command with ResponseData Option!"),
		SetEmbed(embed),
	)
	response := NewInteractionResponse(
		SetType(discordgo.InteractionResponseChannelMessageWithSource),
		SetData(responseData),
	)
	// addCommand(
	// 	&discordgo.ApplicationCommand{
	// 		Name:        "basic-command",
	// 		Description: "Basic-command",
	// 	},
	// 	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
	// 			Type: discordgo.InteractionResponseChannelMessageWithSource,
	// 			Data: &discordgo.InteractionResponseData{
	// 				Content: "This is a basic command!",
	// 			},
	// 		})
	// 	},
	// )
	addCommand(
		&discordgo.ApplicationCommand{
			Name:        "basic-command",
			Description: "Basic-command",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, response)
		},
	)
}
