package config

import "os"

// Config represents the configuration for the Discord bot.
type Config struct {
	BotToken      string
	ChannelID     string
	ApplicationID string
	GuildID       string
}

// LoadConfig loads the configuration from a file or environment variables.
func LoadConfig() (*Config, error) {
	botToken := os.Getenv("BOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	applicationID := os.Getenv("APPLICATION_ID")
	guildID := os.Getenv("GUILD_ID")
	return &Config{
		BotToken:      botToken,
		ChannelID:     channelID,
		ApplicationID: applicationID,
		GuildID:       guildID,
	}, nil
}
