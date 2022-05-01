package usecase

import (
	errorsCustom "News24/internal/app/auth"
	"News24/internal/app/control_users/repository/postgres"
	"News24/internal/common/helpers_function"
	"News24/internal/models"

	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAddUser(t *testing.T) {
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
	controlUsersService := NewContolUserUseCase(&userRepository, fmt.Sprintf("%v", config.HASH_SALT))

	testCases := []struct {
		name         string
		InputUser    *models.User
		expectedResp *models.StandardResponses
	}{
		{
			name: "test 1: length username is zero",
			InputUser: &models.User{
				UserName: "",
				Password: "",
				Role:     0,
			},
			expectedResp: &models.StandardResponses{
				Ok:  "false",
				Err: errorsCustom.ZeroLenUsername.Error(),
			},
		},
		{
			name: "test 2: length password less 6 symbols",
			InputUser: &models.User{
				UserName: "123",
				Password: "12345",
				Role:     0,
			},
			expectedResp: &models.StandardResponses{
				Ok:  "false",
				Err: errorsCustom.LenPasswordLessSixSymbols.Error(),
			},
		},
		{
			name: "test 3: all is ok",
			InputUser: &models.User{
				UserName: "test",
				Password: "testtest",
				Role:     0,
			},
			expectedResp: &models.StandardResponses{
				Ok:  "true",
				Err: "",
			},
		},
		{
			name: "test 4: duplicate found",
			InputUser: &models.User{
				UserName: "test",
				Password: "testtest",
				Role:     0,
			},
			expectedResp: &models.StandardResponses{
				Ok:  "false",
				Err: errorsCustom.FindUserDuplicate.Error(),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := controlUsersService.AddUser(
				context.Background(),
				fmt.Sprintf("%v", tc.InputUser.UserName),
				fmt.Sprintf("%v", tc.InputUser.Password),
				tc.InputUser.Role.(int))

			assert.Equal(t, tc.expectedResp, resp)

		})
	}
}

func TestDeleteUserForLogin(t *testing.T) {
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
	controlUsersService := NewContolUserUseCase(&userRepository, fmt.Sprintf("%v", config.HASH_SALT))

	err = controlUsersService.userRepo.CreateUser(
		context.Background(),
		&models.User{
			UserName: "test",
			Password: "test123",
			Role:     0})

	assert.Equal(t, nil, err)

	testCases := []struct {
		name         string
		InputName    string
		expectedResp *models.StandardResponses
	}{
		{
			name:      "test 1: length username is zero",
			InputName: "",
			expectedResp: &models.StandardResponses{
				Ok:  "false",
				Err: errorsCustom.ZeroLenUsername.Error(),
			},
		},
		{
			name:      "test 2: user not found",
			InputName: "user not in bd",
			expectedResp: &models.StandardResponses{
				Ok:  "false",
				Err: errorsCustom.UserNotFound.Error(),
			},
		},
		{
			name:      "test 3: user found",
			InputName: "test",
			expectedResp: &models.StandardResponses{
				Ok:  "true",
				Err: "",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := controlUsersService.DeleteUserForLogin(
				context.Background(), tc.InputName)

			assert.Equal(t, tc.expectedResp.Err, resp.Err)
			assert.Equal(t, tc.expectedResp.Ok, resp.Ok)

		})
	}
}

func TestUpdateRoleUserForLogin(t *testing.T) {
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
	controlUsersService := NewContolUserUseCase(&userRepository, fmt.Sprintf("%v", config.HASH_SALT))

	err = controlUsersService.userRepo.CreateUser(
		context.Background(),
		&models.User{
			UserName: "test",
			Password: "test123",
			Role:     0})

	assert.Equal(t, nil, err)

	testCases := []struct {
		name         string
		InputName    string
		InputRole    int
		expectedResp *models.StandardResponses
	}{
		{
			name:      "test 1: length username is zero",
			InputName: "",
			InputRole: 0,
			expectedResp: &models.StandardResponses{
				Ok:  "false",
				Err: errorsCustom.ZeroLenUsername.Error(),
			},
		},
		{
			name:      "test 2: user not found",
			InputName: "user not in bd",
			InputRole: 0,
			expectedResp: &models.StandardResponses{
				Ok:  "false",
				Err: errorsCustom.UserNotFound.Error(),
			},
		},
		{
			name:      "test 3: user updated",
			InputName: "test",
			InputRole: 0,
			expectedResp: &models.StandardResponses{
				Ok:  "true",
				Err: "",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := controlUsersService.UpdateRoleUserForLogin(
				context.Background(), tc.InputName, tc.InputRole)

			assert.Equal(t, tc.expectedResp.Err, resp.Err)
			assert.Equal(t, tc.expectedResp.Ok, resp.Ok)

		})
	}
}

func TestGetUserForLogin(t *testing.T) {
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
	controlUsersService := NewContolUserUseCase(&userRepository, fmt.Sprintf("%v", config.HASH_SALT))

	err = controlUsersService.userRepo.CreateUser(
		context.Background(),
		&models.User{
			UserName: "test",
			Password: "test123",
			Role:     0})

	assert.Equal(t, nil, err)

	testCases := []struct {
		name         string
		InputName    string
		expectedResp *models.GetUserResponses
	}{
		{
			name:      "test 1: length username is zero",
			InputName: "",
			expectedResp: &models.GetUserResponses{
				Ok:   "false",
				Err:  errorsCustom.ZeroLenUsername.Error(),
				User: nil,
			},
		},
		{
			name:      "test 2: user not found",
			InputName: "user not in bd",
			expectedResp: &models.GetUserResponses{
				Ok:   "false",
				Err:  errorsCustom.UserNotFound.Error(),
				User: nil,
			},
		},
		{
			name:      "test 3: user found",
			InputName: "test",
			expectedResp: &models.GetUserResponses{
				Ok:  "true",
				Err: "",
				User: &models.User{
					UserName: "test",
					Role:     0,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := controlUsersService.GetUserForLogin(
				context.Background(), tc.InputName)

			assert.Equal(t, tc.expectedResp.Err, resp.Err)
			assert.Equal(t, tc.expectedResp.Ok, resp.Ok)
			if tc.expectedResp.Ok == "false" && tc.expectedResp.User != nil {
				assert.Equal(t, tc.expectedResp.User.UserName, resp.User.UserName)
				assert.Equal(t, tc.expectedResp.User.Role, resp.User.Role)
			}

		})
	}
}

func TestGetAllUsers(t *testing.T) {
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
	controlUsersService := NewContolUserUseCase(&userRepository, fmt.Sprintf("%v", config.HASH_SALT))

	testCases := []struct {
		name         string
		expectedResp *models.GetAllUsersResponses
		isAddNow     bool
	}{
		{
			name: "test 1: len user list is 0",
			expectedResp: &models.GetAllUsersResponses{
				Ok:    "true",
				Err:   "",
				Users: nil,
			},
			isAddNow: false,
		},
		{
			name: "test 2: len user list is 1",
			expectedResp: &models.GetAllUsersResponses{
				Ok:  "true",
				Err: "",
				Users: []*models.User{
					{
						UserName: "test_1",
						Password: "",
						Role:     0,
					},
				},
			},
			isAddNow: true,
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.isAddNow {
				err = controlUsersService.userRepo.CreateUser(
					context.Background(),
					&models.User{
						UserName: fmt.Sprintf("%s_%d", "test", i),
						Password: "test123",
						Role:     0})

				assert.Equal(t, nil, err)
			}

			resp := controlUsersService.GetAllUsers(context.Background())
			assert.Equal(t, tc.expectedResp.Err, resp.Err)
			assert.Equal(t, tc.expectedResp.Ok, resp.Ok)

			assert.Equal(t, len(tc.expectedResp.Users), len(resp.Users))

		})
	}
}
