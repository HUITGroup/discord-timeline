package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var db *sqlx.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("info: no .env file")
	}

	// MYSQLへの接続
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDRESS")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPass, mysqlAddr, mysqlPort, mysqlDBName)
	log.Println("info: DSN ->", dsn)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		errMessage := fmt.Sprintf("\x1b[31m%s\x1b[0m\n%s", "Error: Start SQL first.", err)
		log.Fatal(errMessage)

	}

	// MYSQLのスキーマ定義
	schema1 := `
	CREATE TABLE IF NOT EXISTS timeline_channel (
		guild_id    BIGINT UNSIGNED NOT NULL PRIMARY KEY,
		timeline_id BIGINT UNSIGNED NOT NULL
	);`

	schema2 := `
	CREATE TABLE IF NOT EXISTS timeline_message (
		timeline_message_id BIGINT UNSIGNED NOT NULL PRIMARY KEY,
		link_message_id 	BIGINT UNSIGNED NOT NULL,
		original_message_id BIGINT UNSIGNED NOT NULL
	);`
	db.MustExec(schema1)
	db.MustExec(schema2)
}

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("Error creacting Discord session,", err)
	}

	// timelineチャンネルの登録(!timeline)
	dg.AddHandler(registTimelineChannel)
	// timelineに送る
	dg.AddHandler(sendTimeline)
	// 元のメッセージが編集されたら、タイムライン側も編集
	dg.AddHandler(editTimeline)
	// 元のチャットが削除されたら、タイムライン側も削除
	dg.AddHandler(delTimeline)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection,", err)
		return
	}

	log.Println("Bot is now running. Press Ctrl-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
