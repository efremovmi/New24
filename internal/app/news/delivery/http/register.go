package http

import (
	"News24/internal/app/news"
	"News24/internal/common/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc news.UseCase, pathToViews string) {

	h := NewHandler(uc)

	authEndpoints := router.Group("/news")
	{
		authEndpoints.Use(middleware.JwtAuthMiddleware())
		authEndpoints.POST("/save-post", h.SaveNews)
		authEndpoints.POST("/update-post", h.UpdateNewsForId)
		authEndpoints.POST("/delete-post", h.DeleteNewsForHeader)
		authEndpoints.POST("/get-preview-list", h.GetListPreviewNews)
		authEndpoints.GET("/get-all-news", h.GetAllNews)
		authEndpoints.GET("/get-moder-menu", h.GetModeratorMenu)

		// HTML
		authEndpoints.GET("/get-post", h.GetNewsHTMLForHeader)
		authEndpoints.GET("", h.GetNewsByRoleHTML)
		authEndpoints.Static("/views", pathToViews)
	}

}
