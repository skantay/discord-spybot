package main

import (
	"log/slog"

	"github.com/skantay/discord-spybot/internal/app"
)

func main() {
	configPath := "config/config.yaml"

	if err := app.Run(configPath); err != nil {
		slog.Error(err.Error())
	}
}
