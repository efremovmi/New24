package usecase

import (
	ctrlUsers "News24/internal/app/control_users"
	errorsCustom "News24/internal/app/control_users"
	"News24/internal/models"

	"context"
	"crypto/sha1"
	"encoding/hex"
)

type ContolUserUseCase struct {
	userRepo ctrlUsers.UserRepository
	hashSalt string
}

func NewContolUserUseCase(userRepo ctrlUsers.UserRepository, hashSalt string) *ContolUserUseCase {

	return &ContolUserUseCase{
		userRepo: userRepo,
		hashSalt: hashSalt,
	}
}

func (c *ContolUserUseCase) AddUser(ctx context.Context, username,
	password string, role int) (responses *models.StandardResponses) {

	if len(username) == 0 {
		return &models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.ZeroLenUsername.Error(),
		}
	}
	if len(password) < 6 {
		return &models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.LenPasswordLessSixSymbols.Error(),
		}
	}

	userInDB, _ := c.userRepo.GetUserForLogin(ctx, username)
	if userInDB != nil {
		return &models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.FindUserDuplicate.Error(),
		}
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(c.hashSalt))

	user := &models.User{
		UserName: username,
		Password: hex.EncodeToString(pwd.Sum(nil)),
		Role:     role,
	}

	err := c.userRepo.CreateUser(ctx, user)
	if err != nil {
		return &models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.FailedSignUp.Error(),
		}
	}
	return &models.StandardResponses{
		Ok:  "true",
		Err: "",
	}
}

func (c *ContolUserUseCase) DeleteUserForLogin(ctx context.Context, username string) (responses *models.StandardResponses) {
	if len(username) == 0 {
		return &models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.ZeroLenUsername.Error(),
		}
	}

	err := c.userRepo.DeleteUserForLogin(ctx, username)
	if err != nil {
		return &models.StandardResponses{
			Ok:  "false",
			Err: err.Error(),
		}
	}

	return &models.StandardResponses{
		Ok:  "true",
		Err: "",
	}
}

func (c *ContolUserUseCase) UpdateRoleUserForLogin(ctx context.Context,
	username string, role int) (responses *models.StandardResponses) {

	if len(username) == 0 {
		return &models.StandardResponses{
			Ok:  "false",
			Err: errorsCustom.ZeroLenUsername.Error(),
		}
	}

	err := c.userRepo.UpdateUserRoleForLogin(ctx, username, role)
	if err != nil {
		return &models.StandardResponses{
			Ok:  "false",
			Err: err.Error(),
		}
	}

	return &models.StandardResponses{
		Ok:  "true",
		Err: "",
	}
}

func (c *ContolUserUseCase) GetUserForLogin(ctx context.Context, username string) (responses *models.GetUserResponses) {

	if len(username) == 0 {
		return &models.GetUserResponses{
			Ok:   "false",
			Err:  errorsCustom.ZeroLenUsername.Error(),
			User: nil,
		}
	}

	user, err := c.userRepo.GetUserForLogin(ctx, username)
	if err != nil {
		return &models.GetUserResponses{
			Ok:   "false",
			Err:  err.Error(),
			User: nil,
		}
	}

	return &models.GetUserResponses{
		Ok:   "true",
		Err:  "",
		User: user,
	}
}

func (c *ContolUserUseCase) GetAllUsers(ctx context.Context) (responses *models.GetAllUsersResponses) {

	users, err := c.userRepo.GetAllUsers(ctx)
	if err != nil {
		return &models.GetAllUsersResponses{
			Ok:    "false",
			Err:   err.Error(),
			Users: nil,
		}
	}

	return &models.GetAllUsersResponses{
		Ok:    "true",
		Err:   "",
		Users: users,
	}
}
