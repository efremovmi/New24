package control_users

import (
	"News24/internal/models"
)

type UseCase interface {
	AddUser(username, password string, role int) (err error)
	DeleteUserForLogin(username string) (err error)
	UpdateRoleUserForLogin(username string, role int) (err error)
	GetUserForLogin(username string) (user *models.User, err error)
	GetAllUsers() (users []*models.User, err error)
}
