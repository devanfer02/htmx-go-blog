package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/devanfer02/go-todo/infra/env"
)

func NewPgsqlConn() *sqlx.DB {
	dbx, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=disable",
		env.AppEnv.DBUser, 
		env.AppEnv.DBPass, 
		env.AppEnv.DBHost, 
		env.AppEnv.DBName,
	))

	if err != nil {
		log.Fatal("ERR: " + err.Error())
	}

	return dbx 
}