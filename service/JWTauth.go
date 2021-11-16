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


func ValidateToken(tokenString string) (interface{}, error){
    SampleSecret := []byte("secret")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
        // Don't Forget to validate the alg is what you expect
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return SampleSecret, nil
    })
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        fmt.Println("WAS OK")
        fmt.Println(claims["foo"], claims["nbf"])
        return claims, nil
    } else {
        fmt.Println(err)
        return claims, nil
    }

}
