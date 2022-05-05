package auth

import (
	"News24/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (err error)
	GetUserForLoginAndPassword(username, password string) (user *models.User, err error)
	GetUserForLogin(username string) (user *models.User, err error)
}
