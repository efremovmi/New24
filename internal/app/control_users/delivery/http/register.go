package http

import (
	ctrlUsers "News24/internal/app/control_users"
	"News24/internal/common/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc ctrlUsers.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/control-users")
	{
		authEndpoints.Use(middleware.JwtAuthMiddleware())
		authEndpoints.POST("/delete-user", h.DeleteUserForLogin)
		authEndpoints.POST("/add-user", h.AddUser)
		authEndpoints.POST("/get-user", h.GetUserForLogin)
		authEndpoints.POST("/update-user-role", h.UpdateRoleUserForLogin)
		authEndpoints.GET("/get-all-users", h.GetAllUsers)

		// HTML
		authEndpoints.GET("", h.GetModeratorMenuHTML)
		authEndpoints.Static("/static", "/home/max/KURSOVAY/News24/views/control-users/static")

	}

}
