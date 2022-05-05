package usecase

import (
	errorsCustom "News24/internal/app/auth"
	"News24/internal/app/auth/repository/postgres"
	"News24/internal/common/helpers_function"
	"News24/internal/models"
	"crypto/sha1"
	"encoding/hex"
	"os"

	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestSignUp(t *testing.T) {
	config := helpers_function.GetEnvParams()
	psqlconn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_BD_NAME)

	err := helpers_function.CreateTestTable(psqlconn, fmt.Sprintf("%v", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer helpers_function.DropTestTable(psqlconn, fmt.Sprintf("%v", config.POSTGRES_TABLE_USERS))

	userRepository := postgres.UserRepository{}
	err = userRepository.NewUserRepository(psqlconn, fmt.Sprintf("%v_test", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	authService := NewAuthUseCase(&userRepository, fmt.Sprintf("%v", config.HASH_SALT), time.Duration(5))

	testCases := []struct {
		name          string
		InputUser     *models.User
		expectedError error
	}{
		{
			name: "test 1: success sign up",
			InputUser: &models.User{
				UserName: "test",
				Password: "test123",
				Role:     0,
			},
			expectedError: nil,
		},
		{
			name: "test 2: not success sign up",
			InputUser: &models.User{
				UserName: "test",
				Password: "test12",
				Role:     0,
			},
			expectedError: errorsCustom.FindUserDuplicate,
		},
		{
			name: "test 3: success sign up",
			InputUser: &models.User{
				UserName: "test123",
				Password: "test123",
				Role:     0,
			},
			expectedError: nil,
		},
		{
			name: "test 4: Length password less 6 symbols",
			InputUser: &models.User{
				UserName: "test1234",
				Password: "test",
				Role:     0,
			},
			expectedError: errorsCustom.LenPasswordLessSixSymbols,
		},
		{
			name: "test 5: Length username is zero",
			InputUser: &models.User{
				UserName: "",
				Password: "test123",
				Role:     0,
			},
			expectedError: errorsCustom.ZeroLenUsername,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expectedToken, err := getTokenByUser(fmt.Sprintf("%v", tc.InputUser.UserName),
				fmt.Sprintf("%v", tc.InputUser.Password),
				tc.InputUser.Role.(int))
			assert.Equal(t, nil, err)

			token, err := authService.SignUp(fmt.Sprintf("%v", tc.InputUser.UserName),
				fmt.Sprintf("%v", tc.InputUser.Password))

			assert.Equal(t, tc.expectedError, err)
			if tc.expectedError == nil {
				assert.Equal(t, expectedToken, token)
			}

		})
	}
}

func TestSignIn(t *testing.T) {
	config := helpers_function.GetEnvParams()
	psqlconn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_BD_NAME)

	err := helpers_function.CreateTestTable(psqlconn, fmt.Sprintf("%v", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer helpers_function.DropTestTable(psqlconn, fmt.Sprintf("%v", config.POSTGRES_TABLE_USERS))

	userRepository := postgres.UserRepository{}
	err = userRepository.NewUserRepository(psqlconn, fmt.Sprintf("%v_test", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	authService := NewAuthUseCase(&userRepository, fmt.Sprintf("%v", config.HASH_SALT), time.Duration(5))

	authService.SignUp("test", "test1234")
	authService.SignUp("test1", "test1234")

	testCases := []struct {
		name          string
		InputUser     *models.User
		expectedError error
	}{
		{
			name: "test 1: success sign in",
			InputUser: &models.User{
				UserName: "test",
				Password: "test1234",
			},
			expectedError: nil,
		},
		{
			name: "test 2: not success sign in",
			InputUser: &models.User{
				UserName: "test",
				Password: "test112414",
			},
			expectedError: errorsCustom.UserNotFound,
		},
		{
			name: "test 3: not success sign in",
			InputUser: &models.User{
				UserName: "test1",
				Password: "test12345",
			},
			expectedError: errorsCustom.UserNotFound,
		},
		{
			name: "test 4: success sign in",
			InputUser: &models.User{
				UserName: "test1",
				Password: "test1234",
			},
			expectedError: nil,
		},
		{
			name: "test 5: Length password less 6 symbols",
			InputUser: &models.User{
				UserName: "test1234",
				Password: "test",
				Role:     0,
			},
			expectedError: errorsCustom.LenPasswordLessSixSymbols,
		},
		{
			name: "test 6: Length username is zero",
			InputUser: &models.User{
				UserName: "",
				Password: "test123",
				Role:     0,
			},
			expectedError: errorsCustom.ZeroLenUsername,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			token, err := authService.SignIn(
				fmt.Sprintf("%v", tc.InputUser.UserName),
				fmt.Sprintf("%v", tc.InputUser.Password))

			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.NotEqual(t, 0, len(token))
			} else {
				assert.Equal(t, 0, len(token))
			}

		})
	}
}

func getTokenByUser(username, password string, role int) (token string, err error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(os.Getenv("HASH_SALT")))

	user := &models.User{
		UserName: username,
		Password: hex.EncodeToString(pwd.Sum(nil)),
		Role:     role,
	}
	token, err = helpers_function.GetTokenByUser(user)
	return token, err
}
