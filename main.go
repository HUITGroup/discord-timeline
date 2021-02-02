package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
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

	// todo sendMessageされたとき、そのチャンネル名がtimes_で始まれば
	// timesチャンネルで発生したsendMessageを全てtimelineチャンネルに送信
	dg.AddHandler(sendTimeline)

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

func sendTimeline(s *discordgo.Session, m *discordgo.MessageCreate) {
	// bot自身の発言は処理する必要なし
	if m.Author.ID == s.State.User.ID {
		return
	}

	discordChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error could not get Channel info,", err)
		return
	}
	//
	fmt.Println(discordChannel.Name)

	// todo チャンネルidからチャンネル名を取得したい
}
