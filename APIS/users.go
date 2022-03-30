package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mini-proyecto-backend/db"
	"github.com/mini-proyecto-backend/schemas"
	//"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func getUser(c *fiber.Ctx) /*error*/ {

	return
	// return err // c.JSON()
}

func Auth(c *fiber.Ctx) error{
   foundUser := searchUser(2)
	return c.JSONP(foundUser)
}

func returnUser(c *fiber.Ctx,id int) error{
	foundUser := searchUser(2)
	return c.JSONP(foundUser)

}

func searchUser(id uint32) error {
	db := db.ConnectToDb()
	user := &schemas.User{Id: id}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}
	return err
}

func searchForAuth(user schemas.AuthLogin){
 /*db := db.ConnectToDb()
 err := db.Model(*schemas.User).Where("name",schemas.User.Name).Where("password",schemas.User.Password).Select()
 if err != nil {
		panic(err)
	}
	return err*/
}

