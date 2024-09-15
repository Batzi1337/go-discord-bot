package main

import (
	"fmt"
	"go-discord-bot/internal/api/hub"
	"go-discord-bot/internal/config"
	"go-discord-bot/internal/embed"

	"github.com/bwmarrin/discordgo"
)

var (
	cfg            *config.Config
	dg             *discordgo.Session
	nightwaveCache = make(map[string]bool)
	messageCache   = make(map[string]interface{})
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	dg, err = initializeBot(cfg.BotToken)
	if err != nil {
		logError("Error creating discord bot:", err)
		return
	}
	defer dg.Close()

	sendNightwaveInfo()
	sendFissuresInfo()
}

func sendNightwaveInfo() {
	nightwaveData, e := hub.GetNightwave(hub.PC)
	if e != nil {
		panic(e)
	}
	fmt.Println(nightwaveData)

	if !nightwaveCache[nightwaveData.ID] {
		embed := embed.CreateNightwaveEmbed(*nightwaveData)

		if _, err := dg.ChannelMessageSendEmbed(cfg.ChannelID, embed); err != nil {
			logError("Error sending message:", err)
		}

		nightwaveCache[nightwaveData.ID] = true
	}
}

func sendFissuresInfo() {
	fissuresData, e := hub.GetFissures(hub.PC)
	if e != nil {
		panic(e)
	}
	fmt.Println(fissuresData)

	for _, fissure := range fissuresData {
		embed := embed.CreateFissureEmbed(fissure)
		if _, err := dg.ChannelMessageSendEmbed(cfg.ChannelID, embed); err != nil {
			logError("Error sending message:", err)
		}
	}
}

func initializeBot(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	if err := dg.Open(); err != nil {
		return nil, err
	}

	return dg, nil
}

func logError(message string, err error) {
	fmt.Println(message, err)
}
