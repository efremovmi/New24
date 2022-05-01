package postgres

import (
	errorsCustom "News24/internal/app/control_users"
	"News24/internal/common/helpers_function"
	"News24/internal/models"

	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
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

	err = repo.CreateUser(context.Background(),
		&models.User{
			UserName: "test",
			Password: "test",
			Role:     0})
	if err != nil {
		log.Fatalf("Error: %v", err)
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

	err = repo.CreateUser(context.Background(), &models.User{UserName: "test", Password: "test", Role: 0})
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
			user, err := repo.GetUserForLogin(context.Background(), tc.payload["username"])
			if tc.expectedError != nil {
				assert.Equal(t, errorsCustom.UserNotFound, err)
			} else {
				user.ID = nil
				assert.Equal(t, tc.expectedUser, user)
			}
		})
	}
}

func TestUpdateUserRoleForLogin(t *testing.T) {
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

	err = repo.CreateUser(context.Background(), &models.User{UserName: "test", Password: "test", Role: 0})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	testCases := []struct {
		name          string
		payload       *models.User
		expectedError error
	}{
		{
			name:          "test 1: record updated",
			payload:       &models.User{UserName: "test", Role: 22},
			expectedError: nil,
		},
		{
			name:          "test 2: record not found",
			payload:       &models.User{UserName: "test2", Role: 22},
			expectedError: errorsCustom.UserNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			role, err := strconv.Atoi(fmt.Sprintf("%v", tc.payload.Role))
			assert.Equal(t, nil, err)

			err = repo.UpdateUserRoleForLogin(
				context.Background(),
				fmt.Sprintf("%v", tc.payload.UserName),
				role)

			assert.Equal(t, tc.expectedError, err)
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

	repo := UserRepository{}
	err = repo.NewUserRepository(psqlconn, fmt.Sprintf("%v_test", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	testCases := []struct {
		name             string
		expectedUserList []*models.User
		expectedError    error
		isAddNow         bool
	}{
		{
			name:             "test 1: len user list is 0",
			expectedUserList: []*models.User{},
			expectedError:    nil,
			isAddNow:         false,
		},
		{
			name: "test 2: found 1 user",
			expectedUserList: []*models.User{
				{
					UserName: "test_1",
					Password: "",
					Role:     0,
				},
			},
			expectedError: nil,
			isAddNow:      true,
		},
		{
			name: "test 3: found 2 users",
			expectedUserList: []*models.User{
				{
					UserName: "test_1",
					Password: "",
					Role:     0,
				},
				{
					UserName: "test_2",
					Password: "",
					Role:     0,
				},
			},
			expectedError: nil,
			isAddNow:      true,
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.isAddNow {
				err = repo.CreateUser(
					context.Background(),
					&models.User{
						UserName: fmt.Sprintf("%s_%d", "test", i),
						Password: "test",
						Role:     0})

				assert.Equal(t, nil, err)
			}

			actualUserList, err := repo.GetAllUsers(context.Background())

			assert.Equal(t, len(tc.expectedUserList), len(actualUserList))
			minLen := len(tc.expectedUserList)
			if minLen > len(actualUserList) {
				minLen = len(actualUserList)
			}

			for i := 0; i < minLen; i++ {
				assert.Equal(t, tc.expectedUserList[i].UserName, actualUserList[i].UserName)
				assert.Equal(t, tc.expectedUserList[i].Role, actualUserList[i].Role)
				assert.Equal(t, tc.expectedUserList[i].Password, actualUserList[i].Password)
			}

			assert.Equal(t, tc.expectedError, err)
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

	repo := UserRepository{}
	err = repo.NewUserRepository(psqlconn, fmt.Sprintf("%v_test", config.POSTGRES_TABLE_USERS))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = repo.CreateUser(context.Background(), &models.User{UserName: "test", Password: "test", Role: 0})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = repo.CreateUser(context.Background(), &models.User{UserName: "test1", Password: "test1", Role: 1})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	testCases := []struct {
		name               string
		payload            *models.User
		expectedError      error
		expectedCountUsers int
	}{
		{
			name:               "test 1: user was not deleted. Now count user is 2",
			payload:            &models.User{UserName: "user not in table"},
			expectedError:      errorsCustom.UserNotFound,
			expectedCountUsers: 2,
		},
		{
			name:               "test 2: user deleted. Now count user is 1",
			payload:            &models.User{UserName: "test"},
			expectedError:      nil,
			expectedCountUsers: 1,
		},
		{
			name:               "test 3: user deleted. Now count user is 1",
			payload:            &models.User{UserName: "test1"},
			expectedError:      nil,
			expectedCountUsers: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err = repo.DeleteUserForLogin(
				context.Background(),
				fmt.Sprintf("%v", tc.payload.UserName),
			)
			assert.Equal(t, tc.expectedError, err)

			userList, err := repo.GetAllUsers(context.Background())
			assert.Equal(t, nil, err)
			assert.Equal(t, tc.expectedCountUsers, len(userList))
		})
	}
}
