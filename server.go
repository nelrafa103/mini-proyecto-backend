package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mini-proyecto-backend/apis"
	"github.com/mini-proyecto-backend/db"
	// "github.com/mini-proyecto-backend/schemas"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept"}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	db.ConnectToDb()
	app.Post("/login", apis.LogIn)
	app.Post("/signup", apis.Register)
	app.Post("/account", apis.Account)
	app.Post("/bag", apis.GetProducts)
	app.Listen(":3000")

}
