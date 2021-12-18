package models

type File struct {
    ID      uint `json:"id" gorm"primary_key`
    Content string `json:"content"`
}

type CreateFileInput struct {
    Content string `json:"content" binding:"required"`
}
