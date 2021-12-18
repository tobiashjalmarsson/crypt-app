package controllers

import (
	"github/tobiashjalmarsson/crypt-app/models"

	"github/tobiashjalmarsson/crypt-app/utils"

    "strings"

	"fmt"

	"net/http"

	"github/tobiashjalmarsson/crypt-app/service"

	"github.com/gin-gonic/gin"

)

// Variables for controllers.auth

// Move to enviromental variable, 16 bytes
var key = []byte("0123456789012345")

// Controllers to test the JWT functions

func SignToken(c *gin.Context){
    token, err := service.CreateToken(999999)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"data": "couldent create token"})
        return
    }
    c.JSON(http.StatusAccepted, gin.H{"token" : token})
}

func ValidateToken(c *gin.Context){
    fmt.Println("Inside ValidateToken")
    authHeader := c.Request.Header["Authorization"]
    parsedHeader := strings.Fields(authHeader[0])[1]
    fmt.Println("Header is: ", authHeader)
    fmt.Println("Header content is: ", parsedHeader)
    claims, uid,err := service.ValidateToken(parsedHeader)
    fmt.Println("ERROR IS :", err)
    fmt.Println("UID is: ", uid)
    if err != nil {
        // IF error is not nil, break
        // else continue as authorized
        fmt.Println("Error: ", err.Error())
    } else {
        fmt.Println("Claims are: ", claims)
    }
    c.JSON(http.StatusAccepted, gin.H{"data" : "From validate token"})
}

func LoginUser(c *gin.Context){
    var submittedInfo service.LoginInfo
    if err := c.ShouldBindJSON(&submittedInfo); err != nil {
        fmt.Println("ERROR: ", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
        return
    }
    fmt.Println("Submitted Info: ", submittedInfo)
    var user models.User
    if err := models.DB.Where("email = ?", submittedInfo.Email).First(&user).Error; err != nil {
        fmt.Println("NO USER")
        fmt.Println(err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": "User Not Found"})
        return
    }

    if submittedInfo.LogIn(user.Email, user.Password) {
        token, err := service.CreateToken(int(user.ID))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"data": "Coulden't create token"})
            return
        }
        c.JSON(http.StatusAccepted, gin.H{
            "id" : user.ID,
            "email": user.Email,
            "token": token,
        })
        return

    } else {
        c.JSON(http.StatusBadRequest, gin.H{"data" : false})
    }
}


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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
    var oldUser models.User
    if err := models.DB.Where("Email = ?", input.Email).First(&oldUser).Error; err != nil {
        user := models.User{Email: input.Email, Password: utils.Encrypt(key, input.Password)}
	    models.DB.Create(&user)
	    c.JSON(http.StatusOK, gin.H{"data": user})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error" : "username already exists"})
    }
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

    // If the password is being updated we have to make sure to encrypt it
    if input.Password != "" {
        input.Password = utils.Encrypt(key, input.Password)
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
