package usecase

import (
	"News24/internal/app/auth"
	errorsCustom "News24/internal/app/auth"
	"News24/internal/common/helpers_function"
	"News24/internal/models"

	"context"
	"crypto/sha1"
	"encoding/hex"
	"time"
)

type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	tokenTTLРHours time.Duration) *AuthUseCase {

	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		expireDuration: time.Hour * tokenTTLРHours,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, username,
	password string, role int) (responses *models.AuthResponses) {

	if len(username) == 0 {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.ZeroLenUsername.Error(),
			Token: "",
		}
	}
	if len(password) < 6 {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.LenPasswordLessSixSymbols.Error(),
			Token: "",
		}
	}

	userInDB, err := a.userRepo.GetUserForLogin(ctx, username)
	if userInDB != nil {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.FindUserDuplicate.Error(),
			Token: "",
		}
	}
	if err != nil && err != errorsCustom.UserNotFound {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.FailedSignUp.Error(),
			Token: "",
		}
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		UserName: username,
		Password: hex.EncodeToString(pwd.Sum(nil)),
		Role:     role,
	}

	err = a.userRepo.CreateUser(ctx, user)
	if err != nil {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.FailedSignUp.Error(),
			Token: "",
		}
	}

	token, err := helpers_function.GetTokenByUser(user)
	if err != nil {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.FailedGenToken.Error(),
			Token: token,
		}
	}

	return &models.AuthResponses{
		Ok:    "true",
		Err:   "",
		Token: token,
	}
}

func (a *AuthUseCase) SignIn(ctx context.Context,
	username, password string) (responses *models.AuthResponses) {

	if len(username) == 0 {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.ZeroLenUsername.Error(),
			Token: "",
		}
	}
	if len(password) < 6 {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.LenPasswordLessSixSymbols.Error(),
			Token: "",
		}
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	userInDB, err := a.userRepo.GetUserForLoginAndPassword(ctx, username, hex.EncodeToString(pwd.Sum(nil)))
	if userInDB == nil {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.UserNotFound.Error(),
			Token: "",
		}
	}
	if err != nil {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.FailedSignUp.Error(),
			Token: "",
		}
	}

	token, err := helpers_function.GetTokenByUser(userInDB)
	if err != nil {
		return &models.AuthResponses{
			Ok:    "false",
			Err:   errorsCustom.FailedGenToken.Error(),
			Token: token,
		}
	}

	return &models.AuthResponses{
		Ok:    "true",
		Err:   "",
		Token: token,
	}
}
