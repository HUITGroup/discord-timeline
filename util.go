package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func searchTimelineChannelID(s *discordgo.Session, guildID string) (timelineChannelID string) {
	// 発言元ギルドIDがSQLに登録されていればそのチャンネルのtimelineIDを送る
	db.QueryRowx("SELECT timeline_id FROM timeline_channel WHERE guild_id = ?", guildID).Scan(&timelineChannelID)
	return timelineChannelID
}

// オリジナルのメッセージIDから、タイムライン側のメッセージIDを取る(Update時)
func getTimelineMessageID(s *discordgo.Session, originalMessageID string) (timelineMessageID string) {
	err := db.QueryRowx("SELECT timeline_message_id FROM timeline_message WHERE original_message_id = ?", originalMessageID).Scan(
		&timelineMessageID)
	if err != nil {
		log.Println(err)
		return ""
	}
	return timelineMessageID
}
