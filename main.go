package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	const token string = "test"

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creacting Discord session,", err)
	}

	// todo sendMessageされたとき、そのチャンネル名がtimes_で始まれば
	// timesチャンネルで発生したsendMessageを全てtimelineチャンネルに送信
	dg.AddHandler(sendTimeline)

}

func sendTimeline(s *discordgo.Session, m *discordgo.MessageCreate){
	reciveChannelID := m.ChannelID
	// todo チャンネルidからチャンネル名を取得したい
	reciveChannelName := reciveChannelID.
	if 
}