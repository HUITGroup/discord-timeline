package main

import (
	"github.com/bwmarrin/discordgo"
)

func searchTimelineChannelID(guildID string) (timelineChannelID string, err error) {
	// 発言元ギルドIDがSQLに登録されていればそのチャンネルのtimelineIDを送る
	err = db.QueryRowx("SELECT timeline_id FROM timeline_channel WHERE guild_id = ?", guildID).Scan(&timelineChannelID)
	return
}

// オリジナルのメッセージIDから、タイムライン側のメッセージIDを取る(Update時)
func getTimelineMessageID(s *discordgo.Session, originalMessageID string) (timelineMessageID string, err error) {
	err = db.QueryRowx("SELECT timeline_message_id FROM timeline_message WHERE original_message_id = ?", originalMessageID).Scan(
		&timelineMessageID)
	return
}

func getLinkAndTimelineMessageID(originalMessageID string) (linkMessageID, timelineMessageID string, err error) {
	err = db.QueryRowx("SELECT link_message_id, timeline_message_id FROM timeline_message WHERE original_message_id = ?", originalMessageID).Scan(
		&linkMessageID, &timelineMessageID)
	return
}

func delTimelineFromDB(originalMessageID string) (err error) {
	del, err := db.Preparex("DELETE FROM timeline_message WHERE original_message_id = ?")
	del.Exec(originalMessageID)
	return
}
