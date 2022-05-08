package http

import (
	"News24/internal/app/auth"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase, pathToViews string) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)

		// HTML
		authEndpoints.GET("", h.AuthPage)
		authEndpoints.Static("/static", pathToViews+"/auth/static")
	}
}
