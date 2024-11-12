package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/erik-olsson-op/shared/database"
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/models"
)

func main() {
	logger.Logger.Info("Writing - START ###")
	err := gofakeit.Seed(0)
	if err != nil {
		panic(err)
	}
	var persons = make([]models.Person, 10)
	gofakeit.Slice(&persons)

	err = database.Write(persons)
	if err != nil {
		panic(err)
	}
	logger.Logger.Info("Writing - END ###")
}
