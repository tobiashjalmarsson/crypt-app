package controllers

import (
    //"github/tobiashjalmarsson/crypt-app/models"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github/tobiashjalmarsson/crypt-app/models"
)

// TODO Get/Add/Delete the files

func GetFiles(c *gin.Context){
    uid := c.MustGet("UID").(int)
    fmt.Println("USERID from Controller: ", uid)
    var files []models.File
    if err := models.DB.Find(&files).Where("UserID = ?", uid).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": files})
}

func AddFiles(c *gin.Context){
    uid := c.MustGet("UID").(int)
    var input models.CreateFileInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var user models.User
    if err := models.DB.Where("id = ?", uid).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

    file := models.File{UserID: uid, Content: input.Content, Owner: user}
    models.DB.Create(&file)

    c.JSON(http.StatusCreated, gin.H{"data" : file})
}

func DeleteFile(c *gin.Context){
    fmt.Println("Removing Item By ID")
    c.JSON(http.StatusOK, gin.H{"data": "Removing Item"})
}


