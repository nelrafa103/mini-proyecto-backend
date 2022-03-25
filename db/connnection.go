package db
import (
    "github.com/go-pg/pg/v10"
"os"
"log"
"github.com/mini-proyecto-backend/db/schemas"
"fmt"
)
func ConnectToDb() {
  fmt.Println("Connecting to Database")
  os.Setenv("DB_HOST","ec2-3-222-204-187.compute-1.amazonaws.com")
  os.Setenv("DB_USER","maakakelloctet")
  os.Setenv("DB_NAME", "d3bb6bag4vpoec")
  os.Setenv("DB_PASSWORD","e8a852c5447ee40f04b2c3f8406f80984221f6daa79d0bb9d80fa66b0057b75e")
db:= pg.Connect(&pg.Options{
  Addr: os.Getenv("DB_HOST"),
  User: os.Getenv("DB_USER"),
  Password: os.Getenv("DB_PASSWORD"),
  Database: os.Getenv("DB_NAME"),
})
  defer func (){
    user1 := &schemas.User{
      Name:"admin",
      Email: "ndiaz@intellisys.com",
      Password: "GatoDeDosMetros",
    }
    db.Model(user1).Insert()
    fmt.Println("Inserting new User")
     db.Close()
 
  }()
  defer func() {
        if err := db.Close(); err != nil {
            log.Println(err.Error())
        }
    }()
 defer db.Close()
}
// ConnectToDb()