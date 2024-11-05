package main

import (
	"context"
	"github.com/erik-olsson-op/shared/database"
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/models"
)

func main() {
	logger.Logger.Info("Reading - START ###")
	persons := Read()

	for _, person := range persons {
		logger.Logger.Info(person)
	}
	logger.Logger.Info("Reading - END ###")
}

func Read() []models.Person {
	query := "SELECT * FROM person"
	rows, err := database.Connection.QueryContext(context.Background(), query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var persons []models.Person
	for rows.Next() {
		var person models.Person
		err := rows.Scan(
			&person.Id,
			&person.Name,
			&person.Email,
			&person.Phone)
		if err != nil {
			panic(err)
		}
		persons = append(persons, person)
	}

	return persons
}
