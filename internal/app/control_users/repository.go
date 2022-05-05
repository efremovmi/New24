package control_users

import (
	"News24/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (err error)
	GetUserForLogin(username string) (user *models.User, err error)
	UpdateUserRoleForLogin(username string, role int) (err error)
	GetAllUsers() (user []*models.User, err error)
	DeleteUserForLogin(username string) (err error)
}
