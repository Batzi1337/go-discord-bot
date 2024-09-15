package embed

// This module creates an embed message for the discord bot

import (
	"fmt"

	"go-discord-bot/internal/model"

	"github.com/bwmarrin/discordgo"
)

func createEmbed(title, description, imageURL string, color int, fields []*discordgo.MessageEmbedField) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: title,
		Type:  discordgo.EmbedTypeRich,
		Color: color,
		Provider: &discordgo.MessageEmbedProvider{
			Name: "WarframeStatUs",
			URL:  "https://api.warframestat.us/",
		},
		Description: description,
		Image: &discordgo.MessageEmbedImage{
			URL: imageURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Powered by Batzi1337",
		},
		Fields: fields,
	}
}

func CreateNightwaveEmbed(nwi model.Nightwave) *discordgo.MessageEmbed {
	description := `**Nightwave** is a pirate radio station hosted by the enigmatic [Nora Night](https://warframe.fandom.com/wiki/Nora_Night). In-game, Nightwave serves as a unique [Syndicate](https://warframe.fandom.com/wiki/Syndicate) system that offers various player rewards such as [Mods](https://warframe.fandom.com/wiki/Mod), [Resources](https://warframe.fandom.com/wiki/Resources), [Cosmetic](https://warframe.fandom.com/wiki/Warframe_Cosmetics) items and more through the completion of various mission challenges.

Apart from the reward system, Nightwave will also feature various stories that further expand upon [WARFRAME](https://warframe.fandom.com/wiki/WARFRAME)'s lore through their broadcasts.`
	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Activation",
			Value:  nwi.Activation.String(),
			Inline: true,
		},
		{
			Name:   "Expiry",
			Value:  nwi.Expiry.String(),
			Inline: true,
		},
		{
			Name:   "Active",
			Value:  fmt.Sprintf("%v", nwi.Active),
			Inline: false,
		},
		{
			Name:   "Tag",
			Value:  fmt.Sprintf("%v", nwi.Tag),
			Inline: false,
		},
	}
	return createEmbed("Nightwave Information", description, "https://img.freepik.com/fotos-kostenlos/alien-planet-gebaeude_456031-58.jpg?t=st=1725554536~exp=1725558136~hmac=b829f9d40aefcce898034024f7811832dcaf47af44073d3238039f78ce35b716&w=1800", 0x00ff00, fields)
}

func CreateFissureEmbed(fissure model.Fissure) *discordgo.MessageEmbed {
	description := `**Void Fissures**, also known as **Void Storms** in [Empyrean](https://warframe.fandom.com/wiki/Empyrean) missions, are a type of [mission](https://warframe.fandom.com/wiki/Mission) in which players can open [Void Relics](https://warframe.fandom.com/wiki/Void_Relic) and obtain the treasure within. Relics were introduced with [Update: Specters of the Rail 0.0](https://warframe.fandom.com/wiki/Update_19#Update:_Specters_of_the_Rail) (2016-07-08).`

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Activation",
			Value:  fissure.Activation.String(),
			Inline: true,
		},
		{
			Name:   "Expiry",
			Value:  fissure.Expiry.String(),
			Inline: true,
		},
		{
			Name:   "Mission Type",
			Value:  fmt.Sprintf("%v", fissure.MissionType),
			Inline: false,
		},
		{
			Name:   "Enemy",
			Value:  fmt.Sprintf("%v", fissure.Enemy),
			Inline: false,
		},
		{
			Name:   "Tier",
			Value:  fmt.Sprintf("%v", fissure.Tier),
			Inline: false,
		},
		{
			Name:   "Eta",
			Value:  fmt.Sprintf("%v", fissure.Eta),
			Inline: false,
		},
	}
	return createEmbed("Fissure Information", description, "https://img.freepik.com/fotos-kostenlos/alien-planet-gebaeude_456031-58.jpg?t=st=1725554536~exp=1725558136~hmac=b829f9d40aefcce898034024f7811832dcaf47af44073d3238039f78ce35b716&w=1800", 0x9933ff, fields)
}
