package middleware

import (
	"fmt"
    "strings"
	"github.com/gin-gonic/gin"
    "github/tobiashjalmarsson/crypt-app/service"

)


func AuthenticateUser() gin.HandlerFunc{
    return func(c *gin.Context){
        fmt.Println("MIDDLEWARE")
        authHeader := c.Request.Header["Authorization"]
        parsedHeader := strings.Fields(authHeader[0])[1]
        fmt.Println("Header is: ", authHeader)
        fmt.Println("Header content is: ", parsedHeader)
        claims, uid, err := service.ValidateToken(parsedHeader)
        fmt.Println("ERROR IS :", err)
        fmt.Println("USERID IS IN MIDDLEWARE: ", uid)
        if err != nil {
            // IF error is not nil, break
            // else continue as authorized
            fmt.Println("Error: ", err.Error())
            c.Abort()
        } else {
            fmt.Println("Claims are: ", claims)
            c.Set("UID", uid)
        }
        c.Next()
    }
}
