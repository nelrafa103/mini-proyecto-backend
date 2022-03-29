package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mini-proyecto-backend/db"
	"github.com/mini-proyecto-backend/schemas"
)

func getProducts(c *fiber.Ctx) error {
	db := db.ConnectToDb()
	var products []schemas.Product
	err := db.Model(&products).Select()
	if err != nil {
		panic(err)
	}
	return err
}
