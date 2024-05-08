package spybot

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type recorder interface {
	start()
	stop()
}

type repository interface {
	uploadFile(bucket, key string, file []byte) error
	getDownloadURL(bucket, key string, expiration time.Duration) (string, error)
	listFiles(bucket string) ([]string, error)
}

type controller struct {
	record recorder
	repo   repository
	log    *slog.Logger
}

func registerController(record recorder, repo repository, sessDicord *discordgo.Session, log *slog.Logger) {

	ctrl := controller{
		record: record,
		repo:   repo,
		log:    log,
	}

	go ctrl.spyDaemon(sessDicord)

	sessDicord.AddHandler(ctrl.commands)
}

func (c controller) commands(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch m.Content {
	case "!list":
		c.list(s, m)
	case "!download":
		c.list(s, m)
	}

}

func (c controller) list(s *discordgo.Session, m *discordgo.MessageCreate) {

	if s.Id == m.State.User.ID {
		return
	}

	records, err := c.repo.listFiles("spybot-records")
	if err != nil {
		c.log.Error("failed to list records", "error", err.Error())
		s.ChannelMessageSend(m.ChannelID, "[Internal Server Error] I cannot list records, sorry...")

		return
	}

	var msg strings.Builder

	for i := range records {
		file := records[i]
		file = strings.TrimRight(file, ".")
		_, err := msg.Write([]byte(file))
		if err != nil {
			c.log.Error("failed to create a message", "error", err.Error())
			s.ChannelMessageSend(m.ChannelID, "[Internal Server Error] I cannot list records, sorry...")

			return
		}
	}

	fmt.Println("!list")

	s.ChannelMessageSend(m.ChannelID, "!list")
}

func (c controller) spyDaemon(sessDiscord *discordgo.Session) {
	return

}
