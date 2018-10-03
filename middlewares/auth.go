package middlewares

import (
	"fmt"
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
	UserRole := token.Claims.(jwt.MapClaims)["role"]
	if token.Valid {
		c.Set("UserID", int64(token.Claims.(jwt.MapClaims)["id"].(float64)))
		c.Set("Role", UserRole)
		c.Set("AuthToken", token)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization token",
			"status":  false,
		})
		c.Abort()
	}
}

func GetId(c *gin.Context) int {
	id, err := c.Get("UserID")
	if !err {
		return 0
	}
	return id.(int)
}

func GetRole(c *gin.Context) string {
	role, err := c.Get("Role")
	if !err {
		return ""
	}
	return role.(string)
}
