package http

import (
	"News24/internal/app/auth"
	errorsCustom "News24/internal/app/auth"
	"News24/internal/models"

	"errors"
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
	var inp struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Role     int    `json:"role"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.BadRequest.Error(),
			Token: ""})
		return
	}

	resp := h.useCase.SignUp(c.Request.Context(), inp.UserName, inp.Password, inp.Role)
	if resp.Err != "" {
		c.AbortWithError(http.StatusBadRequest, errors.New(resp.Err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, models.AuthResponses{
		Ok:    "true",
		Err:   resp.Err,
		Token: resp.Token})
}

func (h *Handler) SignIn(c *gin.Context) {
	var inp struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.BadRequest.Error(),
			Token: ""})
		return
	}

	resp := h.useCase.SignIn(c.Request.Context(), inp.UserName, inp.Password)
	if resp.Err != "" {
		if resp.Err == errorsCustom.UserNotFound.Error() {
			c.AbortWithError(http.StatusUnauthorized, errorsCustom.UserNotFound)
			c.JSON(http.StatusUnauthorized, resp)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, errors.New(resp.Err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, models.AuthResponses{
		Ok:    "true",
		Err:   resp.Err,
		Token: resp.Token})
}
