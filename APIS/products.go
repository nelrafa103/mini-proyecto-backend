package apis

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mini-proyecto-backend/db"
	"github.com/mini-proyecto-backend/schemas"
)

func GetProducts(c *fiber.Ctx) error {
	db := db.ConnectToDb()
	products := new(schemas.Products)
	type Respond struct {
		Status string
		Body   *schemas.Products
	}
	validResquest := Respond{Status: "Aproved", Body: products}

	if err := c.BodyParser(products); err != nil {
		return c.JSON(products)
	}

	fmt.Println(products.Action)
	if products.Action == "Adding" {
		products.Products = UpdateProduct(products)
		for i := 0; i < len(products.Products); i++ {
			var email string = GetToken(products.Products[i].Bag)
			fmt.Println(email)
			product := &schemas.Product{Title: products.Products[i].Title, Image: products.Products[i].Image, Description: products.Products[i].Description, Category: products.Products[i].Category, Bag: email, Price: products.Products[i].Price, Quantity: products.Products[i].Quantity}
			ok, err := db.Model(product).Insert()
			if err != nil {
				fmt.Println(ok)
				panic(err)
			}
		}
	} else {
		DeleteProduct(products)
	}
	return c.JSON(validResquest)
}

func DeleteProduct(products *schemas.Products) {
	db := db.ConnectToDb()
	product := new(schemas.Product)
	var Email string = GetToken(products.Products[0].Bag)
	if products.Products[0].Quantity == 0 {
	 db.Model(product).Where("title = ?", products.Products[0].Title).Where("bag = ?", Email).Delete()
	 }
	 productToDelete := &schemas.Product{Title: products.Products[0].Title, Image: products.Products[0].Image, Description: products.Products[0].Description, Category: products.Products[0].Category, Bag: Email, Price: products.Products[0].Price, Quantity: products.Products[0].Quantity}
     ok,err := db.Model(productToDelete).Where("title = ?",  products.Products[0].Title).Where("bag = ?", Email).Update()
	 fmt.Println(ok,err) 
	
	

}

func UpdateProduct(products *schemas.Products) []schemas.ProductOnBag {
	db := db.ConnectToDb()
	var filter []schemas.ProductOnBag
	var counter int = len(products.Products)
	var indexes []int 
	var title string
	filter = products.Products
	for i := 0; i < counter; i++ {
		var Email string = GetToken(products.Products[0].Bag)
		product := &schemas.Product{Title: products.Products[i].Title, Image: products.Products[i].Image, Description: products.Products[i].Description, Category: products.Products[i].Category, Bag: Email, Price: products.Products[i].Price, Quantity: products.Products[i].Quantity}
		ok,err := db.Model(product).Where("title = ?", product.Title).Where("bag = ?", product.Bag).Returning("title").Update(&title)
	    fmt.Println(ok,err) 
	    if err != nil {
	     // Remove an element from the array
	      
           indexes = append(indexes,i)
	    }
	}
	//fmt.Println(indexes)
	filter = RemoveIndex(filter,indexes)
	fmt.Println(filter,indexes)
	return filter
}

func RemoveIndex(s []schemas.ProductOnBag, indexes []int) []schemas.ProductOnBag {
	var filter []schemas.ProductOnBag
    var indexesLen int = len(indexes)
    var isOnTable bool = false
	
		for i := 0; i < len(s); i++ {
		  for x := 0; x < indexesLen; x++ {
            if(i == indexes[x]) {
              isOnTable = true
            }
		  }
		  if(isOnTable == true) {
		  	fmt.Println("The product is not already on the bag")
		  	filter = append(filter,s[i])
		  } 
		}
		
	
	return filter
}


/* 	if i != index {
	//fmt.Println(filter)
				fmt.Println(i, index)
				//fmt.Println(s[i])
				filter = append(filter, s[i])
			}*/
/*Utilizare el fakeapi que envio Leobardo (Link: https://fakestoreapi.com/)*/
