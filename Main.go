package main

import (
	"github/tobiashjalmarsson/crypt-app/controllers"
	"github/tobiashjalmarsson/crypt-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/notes", controllers.FindNotes)
	r.GET("/notes/:id", controllers.FindNote)
	r.POST("/notes", controllers.CreateNote)

	models.ConnectDatabase()

	r.Run()
}
