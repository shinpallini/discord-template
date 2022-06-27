package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	minValues := 1
	customID := "multi-select"
	addCommandWithComponent(
		&discordgo.ApplicationCommand{
			Name:        "select-command",
			Description: "Select-command with Multi select",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Select value",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.SelectMenu{
									CustomID:    customID,
									Placeholder: "Select values",
									MinValues:   &minValues,
									MaxValues:   2,
									Options: []discordgo.SelectMenuOption{
										{
											Label:       "Mad mate",
											Description: "imposter jinei",
											Value:       "madmate",
											Default:     false,
										},
										{
											Label:       "Majo",
											Description: "Kill later",
											Value:       "Majo",
										},
									},
								},
							},
						},
					},
				},
			})
		},
		customID,
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.MessageComponentData().Values
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Selected: %s", strings.Join(data, ", ")),
				},
			})
		},
	)
}
