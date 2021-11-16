package main

import (
	"github/tobiashjalmarsson/crypt-app/controllers"
	"github/tobiashjalmarsson/crypt-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
TODO Add JWT authentication
TODO Rework notes model to be associated with a User
TODO Add possibility of shared notes between groups
TODO Add email to send invites to users
TODO Start working on React<Typescript> Frontend
*/

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	// Note routes
    notesGroup := r.Group("/notes")
    {
        notesGroup.GET("", controllers.FindNotes)
        notesGroup.GET("/:id", controllers.FindNote)
        notesGroup.PATCH("/:id", controllers.UpdateNote)
        notesGroup.DELETE("/:id", controllers.DeleteBook)
        notesGroup.POST("", controllers.CreateNote)
    }
	// User routes
    userGroup := r.Group("/users")
    {
        userGroup.GET("", controllers.FindUsers)
        userGroup.GET("/:id", controllers.FindUser)
        userGroup.POST("", controllers.CreateUser)
        userGroup.PATCH("/:id", controllers.UpdateUser)
        userGroup.DELETE("/:id", controllers.DeleteUser)
    }

	// Login routes
    r.POST("/login", controllers.LoginUser)


    // Testroutes
    r.POST("/sign", controllers.SignToken)
    r.POST("/validate", controllers.ValidateToken)
	models.ConnectDatabase()

	r.Run()
}
