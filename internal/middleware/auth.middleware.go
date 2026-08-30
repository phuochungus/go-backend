package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString != "valid-token" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	c.Next()
}
