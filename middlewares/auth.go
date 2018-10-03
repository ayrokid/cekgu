package middlewares

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token == nil && err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		})
		c.Abort()
	}

	fmt.Println("token verfied")
}

func ExtractClaims(c *gin.Context) (jwt.MapClaims, bool) {
	tokenString := c.Request.Header.Get("token")
	hmacSecretString := "secret" // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
