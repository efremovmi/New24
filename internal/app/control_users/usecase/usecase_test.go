package usecase

import (
	errorsCustom "News24/internal/app/auth"
	"News24/internal/app/control_users/repository/postgres"
	"News24/internal/common/helpers_function"
	"News24/internal/models"

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
		name        string
		InputUser   *models.User
		expectedErr error
	}{
		{
			name: "test 1: length username is zero",
			InputUser: &models.User{
				UserName: "",
				Password: "",
				Role:     0,
			},
			expectedErr: errorsCustom.ZeroLenUsername,
		},
		{
			name: "test 2: length password less 6 symbols",
			InputUser: &models.User{
				UserName: "123",
				Password: "12345",
				Role:     0,
			},
			expectedErr: errorsCustom.LenPasswordLessSixSymbols,
		},
		{
			name: "test 3: all is ok",
			InputUser: &models.User{
				UserName: "test",
				Password: "testtest",
				Role:     0,
			},
			expectedErr: nil,
		},
		{
			name: "test 4: duplicate found",
			InputUser: &models.User{
				UserName: "test",
				Password: "testtest",
				Role:     0,
			},
			expectedErr: errorsCustom.FindUserDuplicate,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = controlUsersService.AddUser(
				fmt.Sprintf("%v", tc.InputUser.UserName),
				fmt.Sprintf("%v", tc.InputUser.Password),
				tc.InputUser.Role.(int))

			assert.Equal(t, tc.expectedErr, err)

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
		&models.User{
			UserName: "test",
			Password: "test123",
			Role:     0})

	assert.Equal(t, nil, err)

	testCases := []struct {
		name        string
		InputName   string
		expectedErr error
	}{
		{
			name:        "test 1: length username is zero",
			InputName:   "",
			expectedErr: errorsCustom.ZeroLenUsername,
		},
		{
			name:        "test 2: user not found",
			InputName:   "user not in bd",
			expectedErr: errorsCustom.UserNotFound,
		},
		{
			name:        "test 3: user found",
			InputName:   "test",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = controlUsersService.DeleteUserForLogin(tc.InputName)

			assert.Equal(t, tc.expectedErr, err)
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
		&models.User{
			UserName: "test",
			Password: "test123",
			Role:     0})

	assert.Equal(t, nil, err)

	testCases := []struct {
		name        string
		InputName   string
		InputRole   int
		expectedErr error
	}{
		{
			name:        "test 1: length username is zero",
			InputName:   "",
			InputRole:   0,
			expectedErr: errorsCustom.ZeroLenUsername,
		},
		{
			name:        "test 2: user not found",
			InputName:   "user not in bd",
			InputRole:   0,
			expectedErr: errorsCustom.UserNotFound,
		},
		{
			name:        "test 3: user updated",
			InputName:   "test",
			InputRole:   0,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err = controlUsersService.UpdateRoleUserForLogin(tc.InputName, tc.InputRole)

			assert.Equal(t, tc.expectedErr, err)
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
		&models.User{
			UserName: "test",
			Password: "test123",
			Role:     0})

	assert.Equal(t, nil, err)

	testCases := []struct {
		name         string
		InputName    string
		expectedErr  error
		expectedUser *models.User
	}{
		{
			name:         "test 1: length username is zero",
			InputName:    "",
			expectedErr:  errorsCustom.ZeroLenUsername,
			expectedUser: nil,
		},
		{
			name:         "test 2: user not found",
			InputName:    "user not in bd",
			expectedErr:  errorsCustom.UserNotFound,
			expectedUser: nil,
		},
		{
			name:         "test 3: user found",
			InputName:    "test",
			expectedErr:  nil,
			expectedUser: &models.User{UserName: "test", Role: 0},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := controlUsersService.GetUserForLogin(tc.InputName)

			assert.Equal(t, tc.expectedErr, err)

			if tc.expectedErr == nil {
				assert.Equal(t, tc.expectedUser.UserName, user.UserName)
				assert.Equal(t, tc.expectedUser.Role, user.Role)
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
		name          string
		expectedErr   error
		expectedUsers []*models.User
		isAddNow      bool
	}{
		{
			name:          "test 1: len user list is 0",
			expectedErr:   nil,
			expectedUsers: nil,
			isAddNow:      false,
		},
		{
			name:        "test 2: len user list is 1",
			expectedErr: nil,
			expectedUsers: []*models.User{
				{
					UserName: "test_1",
					Password: "",
					Role:     0,
				},
			},
			isAddNow: true,
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.isAddNow {
				err = controlUsersService.userRepo.CreateUser(
					&models.User{
						UserName: fmt.Sprintf("%s_%d", "test", i),
						Password: "test123",
						Role:     0})

				assert.Equal(t, nil, err)
			}

			users, err := controlUsersService.GetAllUsers()
			assert.Equal(t, tc.expectedErr, err)

			assert.Equal(t, len(tc.expectedUsers), len(users))

		})
	}
}
