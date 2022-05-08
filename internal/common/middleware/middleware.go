package middleware

import (
	helpers "News24/internal/common/helpers_function"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
