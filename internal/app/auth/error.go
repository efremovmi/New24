package auth

import "errors"

var (
	UserNotFound              = errors.New("User not found")
	BDNotWorking              = errors.New("The database is down. Ping error")
	IncorrectParamsConnectBD  = errors.New("Error in database connection parameters")
	BadInsertUser             = errors.New("Failed to insert user")
	BadRoleUser               = errors.New("Failed get role user")
	BadSqlRequest             = errors.New("Bad sql request")
	FailedSignUp              = errors.New("Failed sign up")
	FindUserDuplicate         = errors.New("Found user with same login")
	FailedGenToken            = errors.New("Failed generation token")
	ZeroLenUsername           = errors.New("Length username is zero")
	LenPasswordLessSixSymbols = errors.New("Length password must be more 6 symbols")
	BadRequest                = errors.New("Request is not valid")
)
