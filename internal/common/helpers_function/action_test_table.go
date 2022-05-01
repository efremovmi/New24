package helpers_function

import (
	errorsCustom "News24/internal/app/auth"

	"database/sql"
	"fmt"
)

func CreateTestTable(psqlconn, tableName string) (err error) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s_test (LIKE %s INCLUDING ALL);", tableName, tableName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	if _, err = db.Exec(query); err != nil {
		return errorsCustom.BadSqlRequest
	}
	return nil
}

func DropTestTable(psqlconn, tableName string) (err error) {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s_test;", tableName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	if _, err = db.Exec(query); err != nil {
		return errorsCustom.BadSqlRequest
	}
	return nil
}
