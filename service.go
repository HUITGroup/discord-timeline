package main

import "github.com/bwmarrin/discordgo"

// Service method for controll text service (CRUD)
type Service interface {
	Send()
	SendEmbed()
	Edit()
	EditEmbed()
	Delete()
}

// ServiceImpl example discord, slack, twitter, etc...
type ServiceImpl struct {
	Discord Discord
}

// NewService これ意味ある？何か間違ってそう
func NewService(Discord Discord) *ServiceImpl {
	return &ServiceImpl{
		Discord: Discord,
	}
}

// Discord service
type Discord struct {
	*discordgo.Session
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
