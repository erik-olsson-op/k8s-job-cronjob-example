package database

import (
	"database/sql"
	"fmt"
	"github.com/erik-olsson-op/shared/logger"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var Connection *sql.DB

func init() {
	user := getEnv("DATABASE_USER")
	port := getEnv("DATABASE_PORT")
	password := getEnv("DATABASE_PASSWORD")
	host := getEnv("DATABASE_HOST")
	databaseName := getEnv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, databaseName)
	logger.Logger.Info(dsn)
	var err error
	Connection, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// Ping the database to ensure a connection is established
	err = Connection.Ping()
	if err != nil {
		panic(err)
	}
	logger.Logger.Info("Connected to MariaDB!")

	Connection.SetMaxOpenConns(10)
	Connection.SetMaxIdleConns(5)
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(key)
	}
	return value
}
