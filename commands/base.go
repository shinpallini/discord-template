package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands           = make([]*discordgo.ApplicationCommand, 0)
	CommandHandlers    = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
	ComponentsHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
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

func addComponent(customID string, fn func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	ComponentsHandlers[customID] = fn
}

// 以下Option型と構造体生成関数の記述
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

func SetEmbed(e []*discordgo.MessageEmbed) InteractionRsponseDataOption {
	return func(rd *discordgo.InteractionResponseData) {
		rd.Embeds = append(rd.Embeds, e...)
	}
}

func SetComponent(c []discordgo.MessageComponent) InteractionRsponseDataOption {
	return func(rd *discordgo.InteractionResponseData) {
		rd.Components = append(rd.Components, c...)
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

func SetEmbedField(ef []*discordgo.MessageEmbedField) MessageEmbedOption {
	return func(e *discordgo.MessageEmbed) {
		e.Fields = append(e.Fields, ef...)
	}
}

func NewMessageEmbed(options ...MessageEmbedOption) *discordgo.MessageEmbed {
	e := &discordgo.MessageEmbed{}

	for _, opt := range options {
		opt(e)
	}
	return e
}

// MessageEmbedField構造体の初期化のためのOptionと関数
type MessageEmbedFieldOption func(ef *discordgo.MessageEmbedField)

func SetEmbedFieldName(s string) MessageEmbedFieldOption {
	return func(ef *discordgo.MessageEmbedField) {
		ef.Name = s
	}
}

func SetEmbedFieldValue(s string) MessageEmbedFieldOption {
	return func(ef *discordgo.MessageEmbedField) {
		ef.Value = s
	}
}

func SetEmbedFieldInline(b bool) MessageEmbedFieldOption {
	return func(ef *discordgo.MessageEmbedField) {
		ef.Inline = b
	}
}

func NewMessageEmbedField(options ...MessageEmbedFieldOption) *discordgo.MessageEmbedField {
	ef := &discordgo.MessageEmbedField{}

	for _, opt := range options {
		opt(ef)
	}
	return ef
}

// MessageComponent構造体初期化のためのOptionと関数
type ActionsRowOption func(*discordgo.ActionsRow)

func AddLinkButton(label string, url string) ActionsRowOption {
	return func(r *discordgo.ActionsRow) {
		r.Components = append(r.Components, discordgo.Button{
			Style: discordgo.LinkButton,
			Label: label,
			URL:   url,
		})
	}
}

func AddCustomButton(style discordgo.ButtonStyle, label string, customID string) ActionsRowOption {
	return func(r *discordgo.ActionsRow) {
		r.Components = append(r.Components, discordgo.Button{
			Style:    style,
			Label:    label,
			CustomID: customID,
		})
	}
}

func NewActionsRow(options ...ActionsRowOption) *discordgo.ActionsRow {
	c := &discordgo.ActionsRow{}

	for _, opt := range options {
		opt(c)
	}
	return c
}
