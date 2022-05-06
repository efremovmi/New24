package http

import (
	ctrlUsers "News24/internal/app/control_users"
	errorsCustom "News24/internal/app/control_users"
	helpers "News24/internal/common/helpers_function"
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

func (h *Handler) AddUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	_, err := helpers.IsAdmin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	var inp struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Role     int    `json:"role"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error()})
		return
	}

	if inp.Role < 0 || inp.Role > 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.InvalidValueRole.Error()})
		return
	}

	err = h.useCase.AddUser(inp.UserName, inp.Password, inp.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) DeleteUserForLogin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	userAdmin, err := helpers.IsAdmin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	var inp struct {
		UserName string `json:"username"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error()})
		return
	}

	if userAdmin.UserName == inp.UserName {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.DeletingSelfError.Error()})
		return
	}

	err = h.useCase.DeleteUserForLogin(inp.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) UpdateRoleUserForLogin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	userAdmin, err := helpers.IsAdmin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	var inp struct {
		UserName string `json:"username"`
		Role     int    `json:"role"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error()})
		return
	}

	if userAdmin.UserName == inp.UserName {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.UpdatingSelfRoleError.Error()})
		return
	}

	if inp.Role < 0 || inp.Role > 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.InvalidValueRole.Error()})
		return
	}

	err = h.useCase.UpdateRoleUserForLogin(inp.UserName, inp.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) GetUserForLogin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	_, err := helpers.IsAdmin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	var inp struct {
		UserName string `json:"username"`
	}

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errorsCustom.BadRequest.Error(),
		})
		return
	}

	user, err := h.useCase.GetUserForLogin(inp.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	_, err := helpers.IsAdmin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	users, err := h.useCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *Handler) GetModeratorMenuHTML(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	_, err := helpers.IsAdmin(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "control-users/control-users.html", nil)
	return
}
