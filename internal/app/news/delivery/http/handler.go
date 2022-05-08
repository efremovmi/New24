package http

import (
	"News24/internal/app/news"
	errorsCustom "News24/internal/app/news"
	helpers "News24/internal/common/helpers_function"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	useCase news.UseCase
}

func NewHandler(useCase news.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) SaveNews(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	err := helpers.IsModerator(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	header := c.PostForm("header")
	news := c.PostForm("news")
	image, headerImage, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadImage.Error(),
		})
		return
	}

	statusCode, err := h.useCase.SaveNews(image, header, news, headerImage.Filename)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) DeleteNewsForHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	err := helpers.IsModerator(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	header := c.PostForm("header")

	statusCode, err := h.useCase.DeleteNewsForHeader(header)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) UpdateNewsForId(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	err := helpers.IsModerator(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadId.Error()})
		return
	}

	header := c.PostForm("header")
	news := c.PostForm("news")
	image, headerImage, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadImage.Error(),
		})
		return
	}

	statusCode, err := h.useCase.UpdateNewsForId(image, header, news, headerImage.Filename, id)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) GetNewsHTMLForHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	header, ok := c.GetQuery("header")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error()})
		return
	}

	newsModel, statusCode, err := h.useCase.GetNewsHTMLForHeader(header)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}
	splitPathToHTML := strings.Split(newsModel.PathToHTML, header+"/")
	pathToHTML := ""
	if len(splitPathToHTML) == 2 {
		pathToHTML = header + "/" + splitPathToHTML[1]
	}
	c.HTML(http.StatusOK, pathToHTML, nil)
}

func (h *Handler) GetNewsByRoleHTML(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	_, err := helpers.IsAdmin(c)
	if err == nil {
		c.HTML(http.StatusOK, "news/admin_news.html", nil)
		return
	}

	err = helpers.IsModerator(c)
	if err == nil {
		c.HTML(http.StatusOK, "news/moder_news.html", nil)
		return
	}

	c.HTML(http.StatusOK, "news/client_news.html", nil)
}

func (h *Handler) GetListPreviewNews(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var inp struct {
		LastId int `json:"last_id"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error()})
		return
	}

	previewNewsList, statusCode, err := h.useCase.GetListPreviewNews(inp.LastId)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"previewList": previewNewsList,
	})

}

func (h *Handler) GetAllNews(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	err := helpers.IsModerator(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	newsList, _, statusCode, err := h.useCase.GetAllNews()
	if err != nil {
		c.JSON(statusCode, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"news_list": newsList,
	})

}

func (h *Handler) GetModeratorMenu(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	err := helpers.IsModerator(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "news/moder-menu/moder-menu.html", nil)
}
