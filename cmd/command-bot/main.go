package main

import (
	"fmt"
	"go-discord-bot/internal/api/market"
	"go-discord-bot/internal/config"
	"go-discord-bot/internal/model"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	commands           []*discordgo.ApplicationCommand
	registeredCommands []*discordgo.ApplicationCommand
)

func main() {
	run()
}

func run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logError("Error loading config:", err)
		return
	}

	dg, err := initializeBot(cfg.BotToken)
	if err != nil {
		logError("Error creating discord bot:", err)
		return
	}
	defer dg.Close()

	if err := registerCommands(dg, cfg.ApplicationID, cfg.GuildID); err != nil {
		logError("Error registering commands:", err)
		return
	}

	items, err := fetchItems()
	if err != nil {
		logError("Error fetching items:", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handleOrderCommand(s, i, items)
	})

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handleAutocomplete(s, i, items)
	})

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "giveaway" {
			handleGiveawayCommand(s, i)
		}
	})

	waitForExit()
}

func waitForExit() {
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	fmt.Println("Shutting down gracefully...")
}

func initializeBot(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	return dg, nil
}

func registerCommands(dg *discordgo.Session, appID, guildID string) error {
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "order",
			Description: "Place an order in the format WTS/WTB",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "type",
					Description: "Type of order (WTS/WTB)",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: "WTS", Value: "WTS"},
						{Name: "WTB", Value: "WTB"},
					},
				},
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "item",
					Description:  "Item to buy or sell",
					Required:     true,
					Autocomplete: true, // Enable autocomplete
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "price",
					Description: "Price of the item",
					Required:    true,
				},
			},
		},
		{
			Name:        "giveaway",
			Description: "Start a giveaway",
		},
	}

	for _, cmd := range commands {
		createdCmd, err := dg.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			return err
		}
		registeredCommands = append(registeredCommands, createdCmd)
	}
	return nil
}

func fetchItems() ([]*model.Item, error) {
	items, err := market.FetchItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func handleGiveawayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	message, err := s.ChannelMessageSend(i.ChannelID, "React with ✅ to enter the giveaway!")
	if err != nil {
		fmt.Println("Error sending message: ", err)
		return
	}

	err = s.MessageReactionAdd(i.ChannelID, message.ID, "✅")
	if err != nil {
		fmt.Println("Error adding reaction: ", err)
		return
	}

	users := make(map[string]bool)
	s.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		if r.MessageID == message.ID && r.Emoji.Name == "✅" {
			users[r.UserID] = true
		}
	})

	time.AfterFunc(1*time.Hour, func() {
		if len(users) == 0 {
			s.ChannelMessageSend(i.ChannelID, "No one entered the giveaway.")
			return
		}

		userIDs := make([]string, 0, len(users))
		for userID := range users {
			userIDs = append(userIDs, userID)
		}

		winnerID := userIDs[rand.Intn(len(userIDs))]
		winner, err := s.User(winnerID)
		if err != nil {
			fmt.Println("Error fetching user: ", err)
			return
		}

		s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("Congratulations %s! You won the giveaway!", winner.Mention()))
	})
}

func handleOrderCommand(s *discordgo.Session, i *discordgo.InteractionCreate, items []*model.Item) {
	if i.ApplicationCommandData().Name != "order" {
		return
	}

	if len(i.ApplicationCommandData().Options) != 3 {
		return
	}

	orderType := i.ApplicationCommandData().Options[0].StringValue()
	itemName := i.ApplicationCommandData().Options[1].StringValue()
	price := i.ApplicationCommandData().Options[2].IntValue()

	item, ok := findItemByName(items, itemName)
	if !ok {
		logInfo("Item not found!")
		return
	}

	response := fmt.Sprintf("%s [%s](https://warframe.market/items/%s) for %d platinum", orderType, itemName, item.URLName, price)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
	if err != nil {
		logError("Error responding to order command:", err)
	}
}

func handleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate, items []*model.Item) {
	if i.Type != discordgo.InteractionApplicationCommandAutocomplete {
		return
	}

	if i.ApplicationCommandData().Name != "order" {
		return
	}

	searchTerm := i.ApplicationCommandData().Options[1].StringValue()
	itemsResult := make([]*model.Item, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, item := range items {
		wg.Add(1)
		go func(item *model.Item) {
			defer wg.Done()
			if strings.HasPrefix(strings.ToLower(item.ItemName), strings.ToLower(searchTerm)) {
				mu.Lock()
				itemsResult = append(itemsResult, item)
				mu.Unlock()
			}
		}(item)
	}

	wg.Wait()

	if len(itemsResult) > 25 {
		itemsResult = itemsResult[:25]
	}

	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0, len(itemsResult))
	for _, item := range itemsResult {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  item.ItemName,
			Value: item.ItemName,
		})
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})

	if err != nil {
		logError("Error responding to autocomplete:", err)
	}
}

func findItemByName(items []*model.Item, itemName string) (*model.Item, bool) {
	for _, it := range items {
		if it.ItemName == itemName {
			return it, true
		}
	}

	return nil, false
}

func logError(message string, err error) {
	log.Println(message, err)
}

func logInfo(message string) {
	log.Println(message)
}
