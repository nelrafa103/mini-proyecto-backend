package db

import (
	"github.com/go-pg/pg/v10"
	"os"
)

func ConnectToDb()(*pg.DB) {
	 
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_ADDR", ":5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PASSWORD", "marielm20")
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
	return db
}


