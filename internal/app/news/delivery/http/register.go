package http

import (
	"News24/internal/app/news"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc news.UseCase) {

	h := NewHandler(uc)

	authEndpoints := router.Group("/news")
	{
		//authEndpoints.Use(middleware.JwtAuthMiddleware())
		authEndpoints.POST("/save-post", h.SaveNews)
		authEndpoints.POST("/delete-post", h.DeleteNewsForHeader)
		authEndpoints.POST("/get-preview-list", h.GetListPreviewNews)

		authEndpoints.GET("/get-post", h.GetNewsHTMLForHeader)
		authEndpoints.GET("", h.GetNewsByRoleHTML)
		authEndpoints.Static("/views", "/home/max/KURSOVAY/News24/views")
	}

}
