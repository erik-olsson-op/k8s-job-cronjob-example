package main

import (
	"github.com/erik-olsson-op/shared/database"
	"github.com/erik-olsson-op/shared/logger"
)

func main() {
	persons := database.Read()
	for _, person := range persons {
		logger.Logger.Info(person)
	}
}
