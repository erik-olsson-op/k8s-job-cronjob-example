package database

import (
	"database/sql"
	"fmt"
	"github.com/erik-olsson-op/shared/logger"
	"github.com/erik-olsson-op/shared/utils"
	_ "github.com/go-sql-driver/mysql"
)

var Connection *sql.DB

func init() {
	user := utils.GetEnv("DATABASE_USER")
	port := utils.GetEnv("DATABASE_PORT")
	password := utils.GetEnv("DATABASE_PASSWORD")
	host := utils.GetEnv("DATABASE_HOST")
	databaseName := utils.GetEnv("DATABASE_NAME")
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
