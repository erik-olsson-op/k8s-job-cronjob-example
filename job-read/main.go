package main

import (
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/utils"
)

func main() {
	logger.Logger.Info("Reading - START ###")
	persons := utils.Read()

	for _, person := range persons {
		logger.Logger.Info(person)
	}
	logger.Logger.Info("Reading - END ###")
}
