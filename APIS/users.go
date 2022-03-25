package APIS
import (
	 "github.com/gofiber/fiber/v2"
	 "github.com/goccy/go-json"
)

func UserApi(app *fiber.Ctx){
   app.Get("/user", func(c *fiber.Ctx) error {
   	type user struct {
   		Name string,
   		Email string
   	}
    user1 := &user{
      Name:"admin",
      Email: "ndiaz@intellisys.com",
    }
    return c.Json()
  })
  return 
}