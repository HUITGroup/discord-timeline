package main

import (
	"github.com/bwmarrin/discordgo"
)

// Service example discord, slack, twitter, etc...
type Service struct {
	Discord Discord
}

// Doer method for controll text service (CRUD)
type Doer interface {
	Send()
	SendEmbed()
	Edit()
	EditEmbed()
	Delete()
}

// Discord info
type Discord struct {
	*discordgo.Session
}

// NewDiscord is constructor
func NewDiscord(discordToken string) (*Discord, error) {
	var d interface{}
	d, err := discordgo.New("Bot " + discordToken)
	discord := d.(*Discord)
	return discord, err
}

// Send -> send message content for Discord
func (discord *Discord) Send(ChannelID, content string) (*discordgo.Message, error) {
	Message, err := discord.ChannelMessageSend(ChannelID, content)
	return Message, err
}

// SendEmbed -> send embed message content for Discord
func (discord *Discord) SendEmbed(ChannelID, content string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	Message, err := discord.ChannelMessageSendEmbed(ChannelID, embed)
	return Message, err
}

// Edit -> edit message content for Discord
func (discord *Discord) Edit(ChannelID, messageID, content string) (*discordgo.Message, error) {
	Message, err := discord.ChannelMessageEdit(ChannelID, messageID, content)
	return Message, err
}

// EditEmbed -> edit embed message content for Discord
func (discord *Discord) EditEmbed(ChannelID, messageID string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	Message, err := discord.ChannelMessageEditEmbed(ChannelID, messageID, embed)
	return Message, err
}

// Delete -> delete message for Discord
func (discord *Discord) Delete(ChannelID, messageID string) error {
	err := discord.ChannelMessageDelete(ChannelID, messageID)
	return err
}
