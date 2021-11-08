package controllers

import (
	"github/tobiashjalmarsson/crypt-app/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /notes
// Returns all the currently available notes for the user to fetch
// Add authentication based on creator
/*
{
	EMPTY
}
*/
func FindNotes(c *gin.Context) {
	var notes []models.Note
	models.DB.Find(&notes)

	c.JSON(http.StatusOK, gin.H{"data": notes})
}

// POST /notes
// Creates a note
// add authomatic author based on authentication
/*
{
	title: string
	author: string
	content: string
}
*/
func CreateNote(c *gin.Context) {
	var input models.CreateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the new note

	note := models.Note{Title: input.Title, Author: input.Author, Content: input.Content}
	models.DB.Create(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})

}

// GET /notes/:id
// Get a specific note based on id
/*
{
	Empty
}
*/
func FindNote(c *gin.Context) {
	var note models.Note

	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": note})
}

/*
PATCH /notes/:id
Update note with new title, author and content
{
	"title": string,
	"author": string,
	"content": string,
}
*/

func UpdateNote(c *gin.Context) {

	// First we attempt to get the model
	var note models.Note
	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Note not found"})
		return
	}

	// Check if the data is correctly formated
	var input models.UpdateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&note).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": note})
}

/*
DELETE /books/:id
remove book with id == :id
*/
func DeleteBook(c *gin.Context) {
	// Get the note if there is one in the database
	// with the corresponding ID
	var note models.Note
	if err := models.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Note not found"})
		return
	}

	models.DB.Delete(&note)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
