package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// timelineにメッセージを送る
func sendTimeline(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println("create,", m)
	// bot自身の発言は処理しない
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.Content == "!timeline" {
		return
	}

	// メッセージが送られたチャンネルを取得
	reciveMessageChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error: could not fetch Channel info,", err)
		return
	}

	// メッセージが送られたチャンネルの名前にtimes_を含んでいれば、処理を続ける
	if strings.Contains(reciveMessageChannel.Name, "times_") {
		timelineChannelID := searchTimelineChannelID(s, m.GuildID)
		// timeline チャンネルがない==メッセージの送り先がないため、終了
		if timelineChannelID == "" {
			return
		}

		// timelineチャンネルがある場合、そこに送る
		// messageUpdateならば編集処理、messageCreateならば送信＆DB登録処理
		messageURL := "https://discord.com/channels/" + m.GuildID + "/" + m.ChannelID + "/" + m.ID
		messageEmbedAuthor := &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		}

		messageEmbed := &discordgo.MessageEmbed{
			Description: m.Message.Content,
			Color:       0x111,
			Author:      messageEmbedAuthor,
		}
		// 当該メッセージURL+タイムラインのコンテンツと2つ送信する
		s.ChannelMessageSend(timelineChannelID, messageURL)
		timelineMessage, err := s.ChannelMessageSendEmbed(timelineChannelID, messageEmbed)
		if err != nil {
			log.Println(err)
		}
		// timelineにbotが投稿したメッセージのID, timesに投稿されたメッセージのIDをSQLに登録
		ins, err := db.Prepare("INSERT INTO timeline_message(timeline_message_id, original_message_id) VALUES(?,?)")
		ins.Exec(timelineMessage.ID, m.Message.ID)
		defer ins.Close()
	}
	return
}

// timelineのメッセージを編集する
func editTimeline(s *discordgo.Session, mup *discordgo.MessageUpdate) {
	log.Println("update,", mup.Author)

	// embedMessege をsend するとなぜかupdateイベントが起こってnilポインタ参照して落ちる
	// それの回避のため、nilポインタかどうかを確かめている
	if mup.Author == nil || mup.Author.ID == s.State.User.ID {
		return
	}

	// 編集されたメッセージが、既にtimeline_messageテーブルに登録されていれば、
	// 編集された内容をtimelineにも反映する
	timelineMessageID := getTimelineMessageID(s, mup.Message.ID)
	if timelineMessageID != "" {
		// 更新点の反映
		messageEmbedAuthor := &discordgo.MessageEmbedAuthor{
			Name:    mup.Author.Username,
			IconURL: mup.Author.AvatarURL(""),
		}
		messageEmbed := &discordgo.MessageEmbed{
			Description: mup.Message.Content,
			Color:       0x111,
			Author:      messageEmbedAuthor,
		}

		timelineChannelID := searchTimelineChannelID(s, mup.GuildID)
		log.Println(timelineChannelID, timelineMessageID, messageEmbed)
		_, err := s.ChannelMessageEditEmbed(timelineChannelID, timelineMessageID, messageEmbed)
		if err != nil {
			log.Println(err)
		}
	}
	return
}

func registTimelineChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Message.Content == "!timeline" {
		// timelineチャンネルとしてそのチャンネルをSQLに登録
		// 既に登録している場合、アップデート
		insert, err := db.Prepare("INSERT INTO timeline_channel(guild_id, timeline_id) VALUES(?,?) ON DUPLICATE KEY UPDATE timeline_id = ?")
		if err != nil {
			log.Println(err)
			return
		}
		insert.Exec(m.GuildID, m.ChannelID, m.ChannelID)
		defer insert.Close()
		s.ChannelMessageSend(m.ChannelID, "Set timeline Channel")
	}
	return
}
