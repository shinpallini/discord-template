package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// InteractionResponse構造体の初期化のためのOptionと関数
type InteractionResponseOption func(r *discordgo.InteractionResponse)

func SetType(t discordgo.InteractionResponseType) InteractionResponseOption {
	return func(r *discordgo.InteractionResponse) {
		r.Type = t
	}
}

func SetData(rd *discordgo.InteractionResponseData) InteractionResponseOption {
	return func(r *discordgo.InteractionResponse) {
		r.Data = rd
	}
}

func NewInteractionResponse(options ...InteractionResponseOption) *discordgo.InteractionResponse {
	ir := &discordgo.InteractionResponse{}

	for _, opt := range options {
		opt(ir)
	}
	return ir
}

// InteractionResponseData構造体の初期化のためのOptionと関数
type InteractionRsponseDataOption func(rd *discordgo.InteractionResponseData)

func SetContent(content string) InteractionRsponseDataOption {
	return func(rd *discordgo.InteractionResponseData) {
		rd.Content = content
	}
}

func SetEmbed(e *discordgo.MessageEmbed) InteractionRsponseDataOption {
	return func(rd *discordgo.InteractionResponseData) {
		rd.Embeds = []*discordgo.MessageEmbed{}
		rd.Embeds = append(rd.Embeds, e)
	}
}

func NewInteractionResponseData(options ...InteractionRsponseDataOption) *discordgo.InteractionResponseData {
	ird := &discordgo.InteractionResponseData{}

	for _, opt := range options {
		opt(ird)
	}
	return ird
}

// MessageEmbed構造体の初期化のためのOptionと関数
type MessageEmbedOption func(e *discordgo.MessageEmbed)

func SetEmbedType(t discordgo.EmbedType) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Type = t
	}
}

func SetTitle(s string) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Title = s
	}
}

func SetDescription(s string) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Description = s
	}
}

// colorは16進数で指定する
func SetColor(i int) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Color = i
	}
}

func NewMessageEmbed(options ...MessageEmbedOption) *discordgo.MessageEmbed {
	e := &discordgo.MessageEmbed{}

	for _, opt := range options {
		opt(e)
	}
	return e
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
