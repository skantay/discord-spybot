package app

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bwmarrin/discordgo"
	"github.com/skantay/discord-spybot/config"
)

func Run(configPath string) error {
	// Load config struct
	cfg, err := config.Get(configPath)
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	// Create discord session with provided token
	sessDiscord, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		return fmt.Errorf("failed to create a discord session: %w", err)
	}

	// Create websocket connection to discord
	if err := sessDiscord.Open(); err != nil {
		return fmt.Errorf("failed to create a websocket connection to discord: %w", err)
	}
	defer sessDiscord.Close()

	// Create AWS sessions
	sessAWS, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.S3.Region),
		Credentials: credentials.NewStaticCredentials(cfg.S3.AccessKey, cfg.S3.SecretKey, ""),
	},
	)
	if err != nil {
		return fmt.Errorf("failed to create an AWS session: %w", err)
	}

	// Create S3 client
	s3 := s3.New(sessAWS)

	// TODO
	repo := newRepo(s3)

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// TODO
	// Handlers are registered and daemon is started
	registerController(nil, repo, sessDiscord, log)
	fmt.Println("bot started")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("shuting down")
	return nil
}
