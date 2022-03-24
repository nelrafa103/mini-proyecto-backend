package db
import (
    "github.com/go-pg/pg/v10"
"os"
"fmt"
"log"
)
func ConnectToDb() {
  os.Setenv("DB_HOST",)
 //fmt.Println(os.Getenv("DB_HOST"),"ec2-3-222-204-187.compute-1.amazonaws.com")
db:= pg.Connect(&pg.Options{
  Addr: os.Getenv("DB_HOST"),
  User: os.Getenv("DB_USER"),
  Password: os.Getenv("DB_PASSWORD"),
  Database: os.Getenv("DB_NAME"),
})
  defer func() {
        if err := db.Close(); err != nil {
            log.Println(err.Error())
        }
    }()
}
