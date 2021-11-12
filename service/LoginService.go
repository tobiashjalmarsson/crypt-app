package service

import (
	"github/tobiashjalmarsson/crypt-app/utils"
    "fmt"
)

// Struct for the login information
type LoginInfo struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Interface for the LoginFunction
type LoginService interface {
	LogIn(email string, password string) bool
}

func (info *LoginInfo) LogIn(email string, password string) bool {
	var key = []byte("0123456789012345")
    fmt.Println("Email: ",info.Email,", Password:", info.Password)
    fmt.Println("UEmail: ", email, ", UPassword:", utils.Decrypt(key, password))
	return info.Email == email && info.Password == utils.Decrypt(key, password)
}
