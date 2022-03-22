package db
import (
    "github.com/go-pg/pg/v10"
"os"
"log"
)
func ConnectToDb() {
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
