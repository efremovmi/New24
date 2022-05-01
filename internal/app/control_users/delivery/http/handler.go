package http

import (
	ctrlUsers "News24/internal/app/control_users"
	errorsCustom "News24/internal/app/control_users"
	"fmt"
	"strconv"

	helpersFunction "News24/internal/common/helpers_function"

	"News24/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase ctrlUsers.UseCase
}

func NewHandler(useCase ctrlUsers.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// TODO try use token
func (h *Handler) AddUser(c *gin.Context) {
	var inp struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Role     int    `json:"role"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.BadRequest.Error()})
		return
	}

	resp := h.useCase.AddUser(c.Request.Context(), inp.UserName, inp.Password, inp.Role)
	if resp.Err != "" {
		c.AbortWithError(http.StatusBadRequest, errors.New(resp.Err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, models.StandardResponses{
		Ok:  "true",
		Err: resp.Err,
	})
}

func (h *Handler) DeleteUserForLogin(c *gin.Context) {
	var inp struct {
		UserName string `json:"username"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.BadRequest.Error()})
		return
	}

	resp := h.useCase.DeleteUserForLogin(c.Request.Context(), inp.UserName)
	if resp.Err != "" {
		c.AbortWithError(http.StatusBadRequest, errors.New(resp.Err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, models.StandardResponses{
		Ok:  "true",
		Err: resp.Err,
	})
}

func (h *Handler) UpdateRoleUserForLogin(c *gin.Context) {
	var inp struct {
		UserName string `json:"username"`
		Role     int    `json:"role"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.BadRequest.Error()})
		return
	}

	resp := h.useCase.UpdateRoleUserForLogin(c.Request.Context(), inp.UserName, inp.Role)
	if resp.Err != "" {
		c.AbortWithError(http.StatusBadRequest, errors.New(resp.Err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, models.StandardResponses{
		Ok:  "true",
		Err: resp.Err,
	})
}

func (h *Handler) GetUserForLogin(c *gin.Context) {
	var inp struct {
		UserName string `json:"username"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, models.GetUserResponses{
			Ok:   "false",
			Err:  errorsCustom.BadRequest.Error(),
			User: nil,
		})
		return
	}

	resp := h.useCase.GetUserForLogin(c.Request.Context(), inp.UserName)
	if resp.Err != "" {
		c.AbortWithError(http.StatusBadRequest, errors.New(resp.Err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, models.GetUserResponses{
		Ok:   "true",
		Err:  resp.Err,
		User: resp.User,
	})
}

func (h *Handler) GetAllUsers(c *gin.Context) {

	user, err := helpersFunction.GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.GetAllUsersResponses{
			Ok:    "false",
			Err:   err.Error(),
			Users: nil,
		})
	}

	role, err := strconv.Atoi(fmt.Sprintf("%v", user.Role))
	if role != 1 || err != nil {
		c.JSON(http.StatusForbidden, models.GetAllUsersResponses{
			Ok:    "false",
			Err:   errorsCustom.Forbidden.Error(),
			Users: nil,
		})
		return
	}

	resp := h.useCase.GetAllUsers(c.Request.Context())
	if resp.Err != "" {
		c.AbortWithError(http.StatusBadRequest, errors.New(resp.Err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, models.GetAllUsersResponses{
		Ok:    "true",
		Err:   resp.Err,
		Users: resp.Users,
	})
}
