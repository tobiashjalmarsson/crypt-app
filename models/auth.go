package models



type User struct {
    ID       uint   `json:"id" gorm:"primary_key;auto_increment"`
    Email    string `json:"email" gorm:"type:varchar()"`
	Password string `json:"password" gorm:"type:varchar()"`
}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type File struct {
    ID      uint `json:"id" gorm:"primary_key;auto_increment"`
    UserID  int `json:"-"`
    Content string `json:"content" gorm:"type:varchar()"`
    Owner   User `json:"owner" binding:"required" gorm:"foreignkey:UserID"`
}

type CreateFileInput struct {
    Content string `json:"content" binding:"required"`
}

