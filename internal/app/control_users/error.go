package control_users

import "errors"

var (
	UserNotFound              = errors.New("User not found")
	BDNotWorking              = errors.New("The database is down. Ping error")
	IncorrectParamsConnectBD  = errors.New("Error in database connection parameters")
	BadInsertUser             = errors.New("Failed to insert user")
	BadRoleUser               = errors.New("Failed get role user")
	BadUpdateUser             = errors.New("Failed update user")
	BadDeleteUser             = errors.New("Failed delete user")
	BadGetAllUsers            = errors.New("Failed get all users")
	ZeroLenUsername           = errors.New("Length username is zero")
	LenPasswordLessSixSymbols = errors.New("Length password must be more 6 symbols")
	FailedAddUser             = errors.New("Failed add user")
	FindUserDuplicate         = errors.New("Found user with same login")
	TokenExpired              = errors.New("Token expired")
	InvalidAccessToken        = errors.New("Invalid access token")
	BadRequest                = errors.New("Request is not valid")
	Forbidden                 = errors.New("No access rights to the resource")
	DeletingSelfError         = errors.New("You can't remove yourself")
	UpdatingSelfRoleError     = errors.New("You can't update role yourself")
	InvalidValueRole          = errors.New("Role value must be in the range [0,1]")
)
