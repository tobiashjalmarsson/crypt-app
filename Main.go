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

	// Note routes
	r.GET("/notes", controllers.FindNotes)
	r.GET("/notes/:id", controllers.FindNote)
	r.PATCH("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteBook)
	r.POST("/notes", controllers.CreateNote)

	// Auth routes
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	models.ConnectDatabase()

	r.Run()
}
