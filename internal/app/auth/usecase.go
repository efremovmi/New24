package auth

import (
	"News24/internal/models"

	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string, role int) (responses *models.AuthResponses)
	SignIn(ctx context.Context, username, password string) (responses *models.AuthResponses)
}
