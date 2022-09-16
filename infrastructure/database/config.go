package database

import (
	"os"
)

type config struct {
	host     string
	database string
	port     string
	user     string
	password string
	timezone string
}

func NewConfigPostgres() *config {
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DB"),
		port:     os.Getenv("POSTGRES_PORT"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		timezone: os.Getenv("POSTGRES_TIMEZONE"),
	}
}
