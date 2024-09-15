package market

import (
	"encoding/json"
	"go-discord-bot/internal/model"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Payload struct {
		Items []*model.Item `json:"items"`
	} `json:"payload"`
}

func FetchItems() ([]*model.Item, error) {
	url := "https://api.warframe.market/v1/items"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers to the request
	req.Header.Set("accept", "application/json")
	req.Header.Set("Language", "en")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Error in FetchItems json.Unmarshal:", err)
		return nil, err
	}

	return response.Payload.Items, nil
}
