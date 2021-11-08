package models

type Note struct {
	ID      uint   `json:"id" gorm"primary_key`
	Title   string `json:"title"`
	Author  string `json:"Author"`
	Content string `json:"Content"`
}

type CreateNoteInput struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateNoteInput struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
}
