package main
 
import (
  "github.com/gofiber/fiber/v2"
  "github.com/mini-proyecto-backend/db"
  "github.com/mini-proyecto-backend/APIS"
)

func main() {
  app := fiber.New()
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })
  APIS.UserApi(app)
  db.ConnectToDb()
  app.Listen(":3000")
}