package main

import (
	"github.com/bwmarrin/discordgo"
)

func searchTimelineChannelID(s *discordgo.Session, guildID string) (timelineChannelID string) {
	// 発言元ギルドIDがSQLに登録されていればそのチャンネルのtimelineIDを送る
	db.QueryRowx("SELECT timeline_id FROM timeline_channel WHERE guild_id = ?", guildID).Scan(&timelineChannelID)
	return timelineChannelID
}

func alreadyTimeline(s *discordgo.Session, messageID string) (timelineMessageID, originalMessageID string) {
	err := db.QueryRowx("SELECT * FROM timeline_message WHERE original_message_id = ?", messageID).Scan(&timelineMessageID, originalMessageID)
	if err != nil {
		return "", ""
	}
	return timelineMessageID, originalMessageID
}
