
package main


import (
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
)


func main(){


    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
            jwt.MapClaims{"foo":"bar"})

    tokenString,err := token.SignedString([]byte("AllYourBase"))

    fmt.Println(tokenString,err)

}
