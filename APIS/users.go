package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mini-proyecto-backend/db"
	"github.com/mini-proyecto-backend/schemas"
)

func getUser(c *fiber.Ctx) /*error*/ {

	return
	// return err // c.JSON()
}

func auth() {

}

func returnUser() schemas.User {
	user := new(schemas.User)
	return *user
}
func searchUser(id uint32) {
	db := db.ConnectToDb()
	user := &schemas.User{Id: id}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}
	return
}
