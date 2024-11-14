package main

import (
	"fmt"
	"github.com/erik-olsson-op/shared/database"
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/models"
	"github.com/erik-olsson-op/shared/utils"
	"github.com/go-resty/resty/v2"
	"math/rand"
)

func main() {
	// Create a Resty client
	client := resty.New()
	// Set the base URL
	host := utils.GetEnv("PRODUCER_HOST")
	addr := fmt.Sprintf("http://%v", host)
	client.SetBaseURL(addr)

	randomInt := rand.Intn(100) // Generates a random integer between 0 and 99
	randomInt++                 // add, can be zero
	endpoint := fmt.Sprintf("/produce/%d", randomInt)
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
	logger.Logger.Info(fmt.Sprintf("writing %d persons to db", len(persons)))
	err = database.Write(persons)
	if err != nil {
		panic(err)
	}
}
