package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loding .env file")
	}
	token := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creacting Discord session,", err)
	}

	dg.AddHandler(sendTimesline)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func sendTimesline(s *discordgo.Session, m *discordgo.MessageCreate) {
	// bot自身の発言は処理しない
	if m.Author.ID == s.State.User.ID {
		return
	}

	// メッセージが送られたチャンネルを取得
	reciveMessageChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error could not fetch Channel info,", err)
		return
	}

	messageURL := "https://discord.com/channels/" + m.GuildID + "/" + m.ChannelID + "/" + m.ID
	contents := m.Author.Username + "\n" + m.Content + "\n" + messageURL

	// メッセージが送られたチャンネルの名前にtimes_を含んでいれば、処理を続ける
	if strings.Contains(reciveMessageChannel.Name, "times_") {
		// 発言されたギルドのGuild構造体を取得
		guildChannels, err := s.GuildChannels(m.GuildID)
		if err != nil {
			log.Println("Error cloud not fetch Guild info,", err)
			return
		}
		// 発言されたギルド配下のチャンネルを全て探索
		// timelineという名前のチャンネルがあれば、そこに発言を送る
		for _, channelInGuild := range guildChannels {
			fmt.Println(channelInGuild)
			// timelineチャンネルにメッセージを送信
			if strings.Contains(channelInGuild.Name, "timeline") {
				fmt.Println(channelInGuild.Name)
				s.ChannelMessageSend(channelInGuild.ID, contents)
				return
			}
		}
	}
	return
}
