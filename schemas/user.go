package schemas

type User struct {
	Id       uint32
	Name     string
	Email    string
	Password string
}

type AuthLogin struct {
	Email string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
}

type AuthRegister struct {
	Email string `json:"email" xml:"email" form:"email"`
	Name string `json:"name" xml:"name" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
}

type AccountRequest struct {
 Email string `json:"email" xml:"email" form:"email"`
 Token string `json:"token" xml:"token" form:"token"`
}

