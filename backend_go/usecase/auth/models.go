package auth

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

type RegisterOutput struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token string `json:"token"`
}
