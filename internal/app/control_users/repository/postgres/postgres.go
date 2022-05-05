package postgres

import (
	errorsCustom "News24/internal/app/control_users"
	"News24/internal/models"

	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type UserRepository struct {
	psqlconn  string
	tableName string
}

func (r *UserRepository) NewUserRepository(psqlconn, tableName string) (err error) {
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return errorsCustom.BDNotWorking
	}
	r.tableName = tableName
	r.psqlconn = psqlconn
	return nil
}

func (r *UserRepository) CreateUser(user *models.User) (err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	userPostgres, err := toPostgresUser(user)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (username, password, role) VALUES ('%s', '%s', %d) RETURNING id;",
		r.tableName,
		userPostgres.Username,
		userPostgres.Password,
		userPostgres.Role)
	var id int
	if err = db.QueryRow(query).Scan(&id); err != nil {
		return errorsCustom.BadInsertUser
	}
	return nil
}

func (r *UserRepository) GetUserForLogin(username string) (user *models.User, err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return user, errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id, username, password, role FROM %s WHERE username = '%s';",
		r.tableName,
		username)

	var id, role int
	var password string

	if err = db.QueryRow(query).Scan(&id, &username, &password, &role); err != nil {
		return user, errorsCustom.UserNotFound
	}

	user = toModel(
		&User{
			ID:       id,
			Username: username,
			Password: password,
			Role:     role,
		})

	return user, nil
}

func (r *UserRepository) GetAllUsers() (userList []*models.User, err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	userList = make([]*models.User, 0)

	if err != nil {
		return userList, errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("select id, username, role from %s;", r.tableName)

	rows, err := db.Query(query)

	for rows.Next() {
		user := models.User{}
		var role64 int64
		err = rows.Scan(&user.ID, &user.UserName, &role64)
		if err != nil {
			return make([]*models.User, 0), errorsCustom.BadGetAllUsers
		}
		user.Role = int(role64)
		user.Password = ""
		userList = append(userList, &user)
	}

	return userList, nil
}

func (r *UserRepository) UpdateUserRoleForLogin(username string, role int) (err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("UPDATE %s set role = %v where username = '%s' returning id;", r.tableName, role, username)

	var id int
	if err = db.QueryRow(query).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return errorsCustom.UserNotFound
		}
		return errorsCustom.BadUpdateUser
	}
	return nil
}

func (r *UserRepository) DeleteUserForLogin(username string) (err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("delete from %s where username = '%s' returning id;", r.tableName, username)

	var id int
	if err = db.QueryRow(query).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return errorsCustom.UserNotFound
		}
		return errorsCustom.BadDeleteUser
	}
	return nil
}

func toPostgresUser(u *models.User) (user *User, err error) {
	role, ok := u.Role.(int)
	if !ok {
		return user, errorsCustom.BadRoleUser
	}

	user = &User{
		Username: fmt.Sprintf("%v", u.UserName),
		Password: fmt.Sprintf("%v", u.Password),
		Role:     role,
	}
	return user, nil
}

func toModel(u *User) *models.User {
	return &models.User{
		ID:       u.ID,
		UserName: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
}
