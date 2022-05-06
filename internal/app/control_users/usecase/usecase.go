package usecase

import (
	ctrlUsers "News24/internal/app/control_users"
	errorsCustom "News24/internal/app/control_users"
	"News24/internal/models"
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

func (c *ContolUserUseCase) AddUser(username, password string, role int) (err error) {

	if len(username) < 6 {
		return errorsCustom.LenUsernameLessSixSymbols
	}

	if len(password) < 6 {
		return errorsCustom.LenPasswordLessSixSymbols
	}

	userInDB, _ := c.userRepo.GetUserForLogin(username)
	if userInDB != nil {
		return errorsCustom.FindUserDuplicate
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(c.hashSalt))

	user := &models.User{
		UserName: username,
		Password: hex.EncodeToString(pwd.Sum(nil)),
		Role:     role,
	}

	err = c.userRepo.CreateUser(user)
	if err != nil {
		return errorsCustom.FailedAddUser
	}

	return nil
}

func (c *ContolUserUseCase) DeleteUserForLogin(username string) (err error) {
	if len(username) < 6 {
		return errorsCustom.LenUsernameLessSixSymbols
	}

	err = c.userRepo.DeleteUserForLogin(username)
	if err != nil {
		return err
	}

	return nil

}

func (c *ContolUserUseCase) UpdateRoleUserForLogin(username string, role int) (err error) {

	if len(username) < 6 {
		return errorsCustom.LenUsernameLessSixSymbols
	}

	err = c.userRepo.UpdateUserRoleForLogin(username, role)
	if err != nil {
		return err
	}

	return nil

}

func (c *ContolUserUseCase) GetUserForLogin(username string) (user *models.User, err error) {

	if len(username) < 6 {
		return nil, errorsCustom.LenUsernameLessSixSymbols
	}

	user, err = c.userRepo.GetUserForLogin(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *ContolUserUseCase) GetAllUsers() (users []*models.User, err error) {

	users, err = c.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
