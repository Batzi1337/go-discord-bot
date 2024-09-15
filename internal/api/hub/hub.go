package hub

// This module gets all the information about the current items from the warframestat.us api
// and returns the information as a struct

import (
	"encoding/json"
	"go-discord-bot/internal/model"
	"log"
	"net/http"
)

type Platform string

const (
	PC   Platform = "pc"
	PS4  Platform = "ps4"
	Xbox Platform = "xb1"
	SW   Platform = "swi"
)

// GetNightwave makes a rest request to the warframestat.us api and returns the current nightwave information
func GetNightwave(p Platform) (*model.Nightwave, error) {
	log.Println("Starting GetNightwave for platform:", p)
	defer log.Println("Finished GetNightwave for platform:", p)

	url := "https://api.warframestat.us/" + string(p) + "/nightwave"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error in GetNightwave http.Get:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var nightwave model.Nightwave
	if err := json.NewDecoder(resp.Body).Decode(&nightwave); err != nil {
		log.Println("Error in GetNightwave json.Decode:", err)
		return nil, err
	}

	return &nightwave, nil
}

// GetAlerts makes a rest request to the warframestat.us api and returns the current alerts
func GetAlerts(p Platform) ([]model.Alert, error) {
	log.Println("Starting GetAlerts for platform:", p)
	defer log.Println("Finished GetAlerts for platform:", p)

	url := "https://api.warframestat.us/" + string(p) + "/alerts"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error in GetAlerts http.Get:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var alerts []model.Alert
	if err := json.NewDecoder(resp.Body).Decode(&alerts); err != nil {
		log.Println("Error in GetAlerts json.Decode:", err)
		return nil, err
	}

	return alerts, nil
}

// GetFissures makes a rest request to the warframestat.us api and returns the current fissures
func GetFissures(p Platform) ([]model.Fissure, error) {
	log.Println("Starting GetFissures for platform:", p)
	defer log.Println("Finished GetFissures for platform:", p)

	url := "https://api.warframestat.us/" + string(p) + "/fissures"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error in GetFissures http.Get:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var fissures []model.Fissure
	if err := json.NewDecoder(resp.Body).Decode(&fissures); err != nil {
		log.Println("Error in GetFissures json.Decode:", err)
		return nil, err
	}

	return fissures, nil
}
