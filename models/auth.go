package models

type User struct {
	ID       uint   `json:"id" gorm"primary_key`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
