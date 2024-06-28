package dto

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
