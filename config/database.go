package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"todolist/helper"
)

const (
	host     = "localhost"
	port     = "5432"
	username = "postgres"
	password = "12345"
	dbname   = "todolist"
)

func InitializedDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	return db
}
