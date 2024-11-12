package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/models"
	"github.com/erik-olsson-op/shared/utils"
	_ "github.com/go-sql-driver/mysql"
)

var connection *sql.DB

func init() {
	user := utils.GetEnv("DATABASE_USER")
	port := utils.GetEnv("DATABASE_PORT")
	password := utils.GetEnv("DATABASE_PASSWORD")
	host := utils.GetEnv("DATABASE_HOST")
	databaseName := utils.GetEnv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, databaseName)
	logger.Logger.Info(dsn)
	var err error
	connection, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// Ping the database to ensure a connection is established
	err = connection.Ping()
	if err != nil {
		panic(err)
	}
	logger.Logger.Info("Connected to MariaDB!")

	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(5)
}

func Write(persons []models.Person) error {
	// Start a transaction
	tx, err := connection.BeginTx(context.Background(), nil)
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
	rows, err := connection.QueryContext(context.Background(), query)
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
