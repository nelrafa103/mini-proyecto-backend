package schemas

type User struct {
	Id       uint32
	Name     string
	Email    string
	Password string
}

type AuthLogin struct {
	Name string
	Password string
}

type AuthSignUp struct {
	Name string
	Password string
	Email string
}