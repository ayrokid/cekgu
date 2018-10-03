package middlewares

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		role := session.Get("role")
		if user == nil || role == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid session token",
			})
		} else {
			c.Next()
		}
	}
}
