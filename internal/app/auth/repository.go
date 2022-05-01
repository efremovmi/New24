package auth

import (
	"News24/internal/models"

	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (err error)
	GetUserForLoginAndPassword(ctx context.Context, username, password string) (user *models.User, err error)
	GetUserForLogin(ctx context.Context, username string) (user *models.User, err error)
}
