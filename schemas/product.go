package schemas

type Product struct {
	Id    uint32
	Title  string
	Image string
	Description string
	Price float32
    Category string
    Bag string
    Quantity uint32
}


type ProductOnBag struct {
	Title string `json:"title" xml:"title" form:"title"`
	Description string `json:"description" xml:"description" form:"description"`
	Category string `json:"category" xml:"category" form:"category"`
	Image string `json:"image" xml:"image" form:"image"`
	Price float32 `json:"price" xml:"price" form:"price"`
    Bag string `json:"bag" xml:"bag" form:"bag"`
    Quantity uint32 `json:"quantity" xml:"quantity" form:"quantity"`
}

type Products struct {
	Products []ProductOnBag `json:"products" xml:"products" form:"products"`
	Action string `json:"action" xml:"action" form:"action"`
}

