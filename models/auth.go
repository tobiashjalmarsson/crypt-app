package models

import (
    "fmt"
)


type User struct {
	ID       uint   `json:"id" gorm"primary_key`
	Email    string `json:"email"`
	Password string `json:"password"`
    Files   []File
}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) removeFileById(id uint){
    fmt.Println("Removing file")
    for _, file := range u.Files {
        fmt.Println("File with id: ", file.ID)
    }
}

func (u *User) addFile(file File){
    u.Files = append(u.Files, file)
}
