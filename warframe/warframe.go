package warframe

import (
	"encoding/json"
	"log"

	"../comm"
	"../models"
)

// API : Simple API for retrieving some statuses of the game
type API struct {
	platform string
	language string
	client   comm.HTTPClient
}

// CreateAPI : Sets the api parameters and creates an instance
func CreateAPI() API {
	e := API{"pc", "en", comm.CreateHTTPClient("https://api.warframestat.us/")}
	return e
}

// GetCetusCycle : Retrieves Cetus cycle data
func (api API) GetCetusCycle() models.Cetus {
	var cetus models.Cetus
	data, err := api.client.Get("/cetusCycle", api.platform, api.language)
	if err != nil {
		log.Print("Request Error", err)
	}
	json.Unmarshal([]byte(data), &cetus)
	log.Print("JSON Unmarshalled", cetus)
	return cetus
}

// GetVallisCycle : Retrieves Vallis cycle data
func (api API) GetVallisCycle() models.Vallis {
	var vallis models.Vallis
	data, err := api.client.Get("/vallisCycle", api.platform, api.language)
	if err != nil {
		log.Print("Request Error", err)
	}
	json.Unmarshal([]byte(data), &vallis)
	log.Print("JSON Unmarshalled", vallis)
	return vallis
}

// GetNews : Retrieve the list of News
func (api API) GetNews() []models.News {
	var news []models.News
	data, err := api.client.Get("/news", api.platform, api.language)
	if err != nil {
		log.Print("Request Error", err)
	}
	json.Unmarshal([]byte(data), &news)
	log.Print("JSON Unmarshalled", news)
	return news
}
