package middleware

import (
	helpers "News24/internal/common/helpers_function"
	"News24/internal/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.StandardResponses{
				Ok:  "false",
				Err: err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
