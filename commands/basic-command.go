package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	customID := "single-select"
	// Nameで定義された文字列がKeyになるので同時に書ける
	embeds := NewList(
		NewMessageEmbed(
			SetEmbedType(discordgo.EmbedTypeRich),
			SetTitle("Embed!"),
			SetDescription("Description!"),
			SetColor(0x15e81c),
			SetEmbedField(
				NewList(
					NewMessageEmbedField(
						SetEmbedFieldName("Embed Field Name"),
						SetEmbedFieldValue("Embed Field Value"),
						SetEmbedFieldInline(true),
					),
				),
			),
		),
		NewMessageEmbed(
			SetEmbedType(discordgo.EmbedTypeRich),
			SetTitle("Embed2!"),
			SetDescription("Description2!"),
			SetColor(0x3a6b8d),
		),
	)
	component := NewList[discordgo.MessageComponent](
		*NewActionsRow(
			AddLinkButton("Linked Button", "https://discord.com/developers/docs/interactions/message-components"),
			AddCustomButton(discordgo.PrimaryButton, "Custom Button", "test"),
		),
		*NewActionsRow(
			AddSingleSelectMenu(
				customID,
				NewList(
					*NewSelectMenuOption(
						"Select a",
						"select_a",
						AddSelectDescription("Selection A"),
						AddSelectDefaultEmoji("😀"),
					),
					*NewSelectMenuOption(
						"Select b",
						"select_b",
						AddSelectDescription("Selection B"),
						AddSelectDefaultEmoji("🐶"),
					),
					*NewSelectMenuOption(
						"Select c",
						"select_c",
						AddSelectDescription("Selection C"),
						AddSelectDefaultEmoji("🐱"),
					),
				),
			),
		),
	)
	responseData := NewInteractionResponseData(
		SetContent("This is a basic-command with ResponseData Option!"),
		SetEmbed(embeds),
		SetComponent(component),
	)
	response := NewInteractionResponse(
		SetType(discordgo.InteractionResponseChannelMessageWithSource),
		SetData(responseData),
	)
	log.Println(responseData)
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
	AddCommandWithComponent(
		&discordgo.ApplicationCommand{
			Name:        "basic-command",
			Description: "Basic-command",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, response)
		},
		customID,
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.MessageComponentData().Values[0]
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Your selection is %s!.", data),
				},
			})
		},
	)
}
