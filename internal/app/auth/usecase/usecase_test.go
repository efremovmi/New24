package usecase

import (
	errorsCustom "News24/internal/app/auth"
	"News24/internal/app/auth/repository/postgres"
	"News24/internal/common/helpers_function"
	"News24/internal/models"

	"context"
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
		expectedResp  *models.AuthResponses
	}{
		{
			name: "test 1: success sign up",
			InputUser: &models.User{
				UserName: "test",
				Password: "test",
				Role:     0,
			},
			expectedResp: &models.AuthResponses{
				Ok:  "true",
				Err: ""},
		},
		{
			name: "test 2: not success sign up",
			InputUser: &models.User{
				UserName: "test",
				Password: "test123",
				Role:     0,
			},
			expectedResp: &models.AuthResponses{
				Ok:  "false",
				Err: errorsCustom.FindUserDuplicate.Error()},
		},
		{
			name: "test 3: success sign up",
			InputUser: &models.User{
				UserName: "test123",
				Password: "test",
				Role:     0,
			},
			expectedResp: &models.AuthResponses{
				Ok:  "true",
				Err: ""},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := authService.SignUp(
				context.Background(),
				fmt.Sprintf("%v", tc.InputUser.UserName),
				fmt.Sprintf("%v", tc.InputUser.Password),
				tc.InputUser.Role.(int))

			assert.Equal(t, tc.expectedResp.Err, resp.Err)
			assert.Equal(t, tc.expectedResp.Ok, resp.Ok)

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

	authService.SignUp(context.Background(), "test", "test", 0)
	authService.SignUp(context.Background(), "test1", "test1", 0)

	testCases := []struct {
		name          string
		InputUser     *models.User
		expectedError error
		expectedResp  *models.AuthResponses
	}{
		{
			name: "test 1: success sign in",
			InputUser: &models.User{
				UserName: "test",
				Password: "test",
			},
			expectedResp: &models.AuthResponses{
				Ok:  "true",
				Err: ""},
		},
		{
			name: "test 2: not success sign in",
			InputUser: &models.User{
				UserName: "test",
				Password: "test1",
			},
			expectedResp: &models.AuthResponses{
				Ok:  "false",
				Err: errorsCustom.UserNotFound.Error()},
		},
		{
			name: "test 3: not success sign in",
			InputUser: &models.User{
				UserName: "test1",
				Password: "test",
			},
			expectedResp: &models.AuthResponses{
				Ok:  "false",
				Err: errorsCustom.UserNotFound.Error()},
		},
		{
			name: "test 4: success sign in",
			InputUser: &models.User{
				UserName: "test1",
				Password: "test1",
			},
			expectedResp: &models.AuthResponses{
				Ok:  "true",
				Err: ""},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := authService.SignIn(
				context.Background(),
				fmt.Sprintf("%v", tc.InputUser.UserName),
				fmt.Sprintf("%v", tc.InputUser.Password))

			assert.Equal(t, tc.expectedResp.Err, resp.Err)
			assert.Equal(t, tc.expectedResp.Ok, resp.Ok)

		})
	}
}
