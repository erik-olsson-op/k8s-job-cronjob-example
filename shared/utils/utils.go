package utils

import (
	"context"
	"github.com/erik-olsson-op/shared/database"
	"github.com/erik-olsson-op/shared/models"
	"os"
)

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(key)
	}
	return value
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
