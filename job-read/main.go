package main

import (
	"fmt"
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/models"
	"github.com/erik-olsson-op/shared/utils"
	"github.com/go-resty/resty/v2"
)

func main() {
	// Create a Resty client
	client := resty.New()
	// Set the base URL
	host := utils.GetEnv("CONSUMER_HOST")
	addr := fmt.Sprintf("http://%v", host)
	client.SetBaseURL(addr)
	endpoint := fmt.Sprintf("/consume")
	// Example of making a GET request
	var persons []models.Person
	_, err := client.R().
		SetHeader(utils.HeaderAccept, "application/json").
		SetResult(&persons).
		Get(endpoint)
	if err != nil {
		logger.Logger.Error(err)
		panic(err)
	}
	logger.Logger.Info(fmt.Sprintf("job read - fetched %d persons", len(persons)))
	for _, person := range persons {
		logger.Logger.Info(person)
	}
}
