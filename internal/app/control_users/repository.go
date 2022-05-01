package control_users

import (
	"News24/internal/models"

	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (err error)
	GetUserForLogin(ctx context.Context, username string) (user *models.User, err error)
	UpdateUserRoleForLogin(ctx context.Context, username string, role int) (err error)
	GetAllUsers(ctx context.Context) (user []*models.User, err error)
	DeleteUserForLogin(ctx context.Context, username string) (err error)
}
