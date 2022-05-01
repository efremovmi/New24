package control_users

import (
	"News24/internal/models"

	"context"
)

type UseCase interface {
	AddUser(ctx context.Context, username, password string, role int) (responses *models.StandardResponses)
	DeleteUserForLogin(ctx context.Context, username string) (responses *models.StandardResponses)
	UpdateRoleUserForLogin(ctx context.Context, username string, role int) (responses *models.StandardResponses)
	GetUserForLogin(ctx context.Context, username string) (responses *models.GetUserResponses)
	GetAllUsers(ctx context.Context) (responses *models.GetAllUsersResponses)
}
