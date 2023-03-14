package auth

type Credentials struct {
	Email    string
	Password string
}

type Session struct {
	Email string
	Token string
}
