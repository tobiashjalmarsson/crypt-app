package middleware

import (
	"fmt"
    "strings"
	"github.com/gin-gonic/gin"
    "github/tobiashjalmarsson/crypt-app/service"

)


func AuthenticateUser() gin.HandlerFunc{
    return func(c *gin.Context){
        authHeader := c.Request.Header["Authorization"]
        parsedHeader := strings.Fields(authHeader[0])[1]
        claims, uid, err := service.ValidateToken(parsedHeader)
        if err != nil {
            // IF error is not nil, break
            // else continue as authorized
            fmt.Println(claims)
            c.Abort()
        } else {
            c.Set("UID", uid)
        }
        c.Next()
    }
}
