package model

type NewUserInput struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
