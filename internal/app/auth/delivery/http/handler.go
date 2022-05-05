package http

import (
	"News24/internal/app/auth"
	errorsCustom "News24/internal/app/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) SignUp(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var inp struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error()})
		return
	}

	token, err := h.useCase.SignUp(inp.UserName, inp.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token})
}

func (h *Handler) SignIn(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var inp struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error()})
		return
	}

	token, err := h.useCase.SignIn(inp.UserName, inp.Password)
	if err != nil {
		if err == errorsCustom.UserNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": errorsCustom.InvalidLoginOrPassword.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token})
}

func (h *Handler) AuthPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/auth.html", nil)
}
