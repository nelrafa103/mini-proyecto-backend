package main

import (
	"github.com/gofiber/fiber/v2"
	 "github.com/mini-proyecto-backend/apis"
	"github.com/mini-proyecto-backend/db"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	db.ConnectToDb()
	app.Post("/login",apis.Auth)
	app.Listen(":3000")
}
