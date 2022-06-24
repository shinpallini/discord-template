package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Option func(r *discordgo.InteractionResponse)

func SetType(t discordgo.InteractionResponseType) Option {
	return func(r *discordgo.InteractionResponse) {
		r.Type = t
	}
}

func SetData(content string) Option {
	return func(r *discordgo.InteractionResponse) {
		r.Data = &discordgo.InteractionResponseData{
			Content: content,
		}
	}
}

func NewInteractionResponse(options ...Option) *discordgo.InteractionResponse {
	ir := &discordgo.InteractionResponse{}

	for _, opt := range options {
		opt(ir)
	}
	return ir
}

var (
	Commands        = make([]*discordgo.ApplicationCommand, 0)
	CommandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
)

func addCommand(command *discordgo.ApplicationCommand, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	_, exist := CommandHandlers[command.Name]
	if exist {
		log.Fatal(fmt.Sprintf("[%s] ← このコマンド名が重複しています！", command.Name))
	}
	// コマンド部分のNameをそのままmapのKeyとして設定しておく
	CommandHandlers[command.Name] = fn
	Commands = append(Commands, command)
}
