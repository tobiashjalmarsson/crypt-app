package service

import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt"
)


func CreateToken() (string, error) {
    SampleSecret := []byte("secret")

    // First we create the token,
    // TODO Add custom claims later
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "foo" : "bar",
        "nbf" : time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
    })

    // Sign the token and get the completeString using the secret
    // TODO Move the secret to an enviroment variable
    tokenString, err := token.SignedString(SampleSecret)
    fmt.Println(tokenString, err)
    return tokenString, err
}

