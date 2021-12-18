package controllers

import (
    //"github/tobiashjalmarsson/crypt-app/models"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// TODO Get/Add/Delete the files

func GetFiles(c *gin.Context){
    fmt.Println("Getting all user files")
    fmt.Println("USERID from Controller: ", c.MustGet("UID"))
    c.JSON(http.StatusOK, gin.H{"data": "Getting items"})
}

func AddFiles(c *gin.Context){
    fmt.Println("Adding Item")
    c.JSON(http.StatusCreated, gin.H{"data" : "Added Item"})
}

func DeleteFile(c *gin.Context){
    fmt.Println("Removing Item By ID")
    c.JSON(http.StatusOK, gin.H{"data": "Removing Item"})
}


