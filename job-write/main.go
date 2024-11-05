package main

import (
	"context"
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

	err = Write(persons)
	if err != nil {
		panic(err)
	}
	logger.Logger.Info("Writing - END ###")
}

func Write(persons []models.Person) error {
	// Start a transaction
	tx, err := database.Connection.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	q := "INSERT INTO person (name, email, phone) VALUES (?, ?, ?)"
	// Prepare the statement once
	stmt, err := tx.PrepareContext(context.Background(), q)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	// Execute the statement for each person
	for _, person := range persons {
		_, err := stmt.Exec(person.Name, person.Email, person.Phone)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit()
}
