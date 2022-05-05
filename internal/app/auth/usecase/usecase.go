package usecase

import (
	"News24/internal/app/auth"
	errorsCustom "News24/internal/app/auth"
	"News24/internal/common/helpers_function"
	"News24/internal/models"

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

func (a *AuthUseCase) SignUp(username, password string) (token string, err error) {

	if len(username) < 6 {
		return "", errorsCustom.LenUsernameLessSixSymbols
	}

	if len(password) < 6 {
		return "", errorsCustom.LenPasswordLessSixSymbols
	}

	userInDB, err := a.userRepo.GetUserForLogin(username)
	if userInDB != nil {
		return "", errorsCustom.FindUserDuplicate
	}

	if err != nil && err != errorsCustom.UserNotFound {
		return "", errorsCustom.FailedSignUp
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		UserName: username,
		Password: hex.EncodeToString(pwd.Sum(nil)),
		Role:     0,
	}

	err = a.userRepo.CreateUser(user)
	if err != nil {
		return "", errorsCustom.FailedSignUp
	}

	token, err = helpers_function.GetTokenByUser(user)
	if err != nil {
		return "", errorsCustom.FailedGenToken
	}

	return token, nil
}

func (a *AuthUseCase) SignIn(username, password string) (token string, err error) {

	if len(username) < 6 {
		return "", errorsCustom.LenUsernameLessSixSymbols
	}

	if len(password) < 6 {
		return "", errorsCustom.LenPasswordLessSixSymbols
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	userInDB, err := a.userRepo.GetUserForLoginAndPassword(username, hex.EncodeToString(pwd.Sum(nil)))
	if userInDB == nil {
		return "", errorsCustom.UserNotFound
	}
	if err != nil {
		return "", errorsCustom.FailedSignIn
	}

	token, err = helpers_function.GetTokenByUser(userInDB)
	if err != nil {
		return token, errorsCustom.FailedGenToken
	}

	return token, nil
}
