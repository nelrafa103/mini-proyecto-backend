package apis

import (
	"fmt"
	"math/rand"
	"github.com/gofiber/fiber/v2"
 
	"github.com/mini-proyecto-backend/db"
	"github.com/mini-proyecto-backend/schemas"
	//"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type Response struct {
	AuthType string
	User     schemas.AuthLogin
}

func getUser(c *fiber.Ctx) /*error*/ {
	// response := new(Response)
	/* if(response.AuthType == "LogIn"){

	   }
	   else if(response.AuthType == "Register") {

	   } */
	/* user := new(schemas.AuthLogin)
	   if err:= c.BodyParser(user); err != nil {
	   	return err
	   }
	   searchForAuth(user) */
	// return c.JSON()
	// return err // c.JSON()
}

func validate(c *fiber.Ctx, mail string) bool {
	// if()
	return true
}

/* func Auth(c *fiber.Ctx) error{
   // foundUser := searchUser("Jose")
	// return c.JSONP(foundUser)
} */

func returnUser(c *fiber.Ctx, id int) error {
	foundUser := searchUser("Manuel")
	return c.JSONP(foundUser)

}

func searchUser(email string) error {
	db := db.ConnectToDb()
	user := &schemas.User{Email: email}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}
	return err
}

func SearchForAuth(user *schemas.AuthLogin) int {
	db := db.ConnectToDb()
	whereToSearch := new(schemas.User)
	var id int
	db.Model(whereToSearch).Column("id").Where("email = ?", user.Email).Where("password = ?", user.Password).Select(&id)

	return id
}

func SearchEmail(user *schemas.AuthLogin) int {
	db := db.ConnectToDb()
	whereToSearch := new(schemas.User)
	var id int
	db.Model(whereToSearch).Column("id").Where("email = ?", user.Email).Select(&id)

	return id
}
func Register(c *fiber.Ctx) error {
	user := new(schemas.AuthLogin)
	if err := c.BodyParser(user); err != nil {
		return c.Send(c.Body())
	}
	searchResult := SearchEmail(user)
	type Respond struct {
		Status string
		Body   *schemas.AuthLogin
	}

	validRequest := Respond{Status: "Log in approved", Body: user}

	if searchResult == 0 {
		newUser := new(schemas.AuthRegister)
		c.BodyParser(newUser)
		InsertNewUser(newUser)
		return c.JSON(validRequest)
	}
	//return searchResult
	invalidRequest := Respond{Status: "Log in Invalid"}
	return c.JSON(invalidRequest)
}

func LogIn(c *fiber.Ctx) error {
	
	user := new(schemas.AuthLogin)

	type Respond struct {
		Status string
		Token  string
		Body   *schemas.AuthLogin
	}

	if err := c.BodyParser(user); err != nil {
		return c.Send(c.Body())
	}
   tokenString := Generate(71)
	searchResult := SearchForAuth(user)

	validRequest := Respond{Status: "Log in approved", Body: user, Token: tokenString}
	fmt.Println(validRequest.Token)
	if searchResult != 0 {
		SetOnTokens(tokenString,validRequest.Body.Email)
		return c.JSON(validRequest)
	}

	invalidRequest := Respond{Status: "Log in Invalid"}
	return c.JSON(invalidRequest)
}

func InsertNewUser(newUser *schemas.AuthRegister) error {
	db := db.ConnectToDb()
	whereToSearch := new(schemas.User)
	count, error := db.Model(whereToSearch).Count()
	whereToSearch.Name = newUser.Name
	whereToSearch.Email = newUser.Email
	whereToSearch.Password = newUser.Password
	whereToSearch.Id = uint32(count + 1)
	test, err := db.Model(whereToSearch).Insert()
	fmt.Println(err, test)
	return error

}

func searchingForEmail(email string) *schemas.AuthRegister {
	var name, password, mail string
	fmt.Println(email)
	db := db.ConnectToDb()
	whereToSearch := new(schemas.User)
	db.Model(whereToSearch).Column("name", "email", "password").Where("email = ?", email).Select(&name, &mail, &password)
	var accountFound = schemas.AuthRegister{Name: name, Password: password, Email: email}
	return &accountFound
}

func Account(c *fiber.Ctx) error {
	user := new(schemas.AccountRequest)
	fmt.Println(c.JSON(c.Body()))
	if err := c.BodyParser(user); err != nil {
		return c.Send(c.Body())
	}
	fmt.Println(user)
   type Respond struct {
   	Status string 
   }
   fmt.Println("Fetching account info")
   var respond = Respond{Status: "Invalid"}
   var email string = GetToken(user.Token)
	if(email ==  "") {
		return c.JSON(respond)
	}
	account := searchingForEmail(email)
	return c.JSON(account)
}
func GetToken(token string) string{
  db := db.ConnectToDb()
  Token := new(schemas.Token)
  var email string
  err:= db.Model(Token).Column("email").Where("token = ?",token).Select(&email)
  fmt.Println(err)
  return email
}
func SetOnTokens(token,email string) {
  db := db.ConnectToDb()
  NewToken := new(schemas.Token)
  NewToken.Token = token
  NewToken.Email = email
  test,err:= db.Model(NewToken).Insert()
  fmt.Println(test,err)
}


func Generate(n int) string {
    var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
    str := make([]rune, n)
    for i := range str {
        str[i] = chars[rand.Intn(len(chars))]
    }
    return string(str)
}