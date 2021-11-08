package controllers

import (
	"github/tobiashjalmarsson/crypt-app/models"

	"github/tobiashjalmarsson/crypt-app/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Variables for controllers.auth

// Move to enviromental variable, 16 bytes
var key = []byte("0123456789012345")

// GET /users
// For testing purposes only
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	for idx, _ := range users {
		users[idx].Password = utils.Decrypt(key, users[idx].Password)
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

/*
POST /users

body
{
	email: string,
	password: string
}
*/
func CreateUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Email: input.Email, Password: utils.Encrypt(key, input.Password)}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

/*
GET /users/:id

body
{
	EMPTY
}
*/
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

/*
PATCH /users/:id

body {
	email: string,
	password: string
}

*/
func UpdateUser(c *gin.Context) {

	// First we attempt to get the model
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Check if the data is correctly formated
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

/*
DELETE /users/:id

body {
	EMPTY
}
*/
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
