package postgres

import (
	errorsCustom "News24/internal/app/auth"
	"News24/internal/common/helpers_function"
	"News24/internal/models"

	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
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

	repo := UserRepository{}
	err = repo.NewUserRepository(psqlconn, fmt.Sprintf("%v_test", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = repo.CreateUser(&models.User{
		UserName: "test",
		Password: "test",
		Role:     0})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func TestGetUserForLoginAndPassword(t *testing.T) {
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

	repo := UserRepository{}
	err = repo.NewUserRepository(psqlconn, fmt.Sprintf("%v_test", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = repo.CreateUser(&models.User{UserName: "test", Password: "test", Role: 0})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	testCases := []struct {
		name          string
		payload       map[string]string
		expectedError error
		expectedUser  *models.User
	}{
		{
			name:          "test 1: record found",
			payload:       map[string]string{"username": "test", "password": "test"},
			expectedError: nil,
			expectedUser:  &models.User{UserName: "test", Password: "test", Role: 0},
		},
		{
			name:          "test 2: record not found",
			payload:       map[string]string{"username": "test", "password": "123"},
			expectedError: errorsCustom.UserNotFound,
			expectedUser:  nil,
		},
		{
			name:          "test 3: record not found 2",
			payload:       map[string]string{"username": "123", "password": "test"},
			expectedError: errorsCustom.UserNotFound,
			expectedUser:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := repo.GetUserForLoginAndPassword(tc.payload["username"], tc.payload["password"])
			if tc.expectedError != nil {
				assert.Equal(t, errorsCustom.UserNotFound, err)
			} else {
				assert.Equal(t, nil, err)
				user.ID = nil
				assert.Equal(t, tc.expectedUser, user)
			}
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

	repo := UserRepository{}
	err = repo.NewUserRepository(psqlconn, fmt.Sprintf("%v_test", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = repo.CreateUser(&models.User{UserName: "test", Password: "test", Role: 0})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	testCases := []struct {
		name          string
		payload       map[string]string
		expectedError error
		expectedUser  *models.User
	}{
		{
			name:          "test 1: record found",
			payload:       map[string]string{"username": "test"},
			expectedError: nil,
			expectedUser:  &models.User{UserName: "test", Password: "test", Role: 0},
		},
		{
			name:          "test 2: record not found",
			payload:       map[string]string{"username": "test12"},
			expectedError: errorsCustom.UserNotFound,
			expectedUser:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := repo.GetUserForLogin(tc.payload["username"])
			if tc.expectedError != nil {
				assert.Equal(t, errorsCustom.UserNotFound, err)
			} else {
				user.ID = nil
				assert.Equal(t, tc.expectedUser, user)
			}
		})
	}
}
