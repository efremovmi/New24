package control_users

import "errors"

var (
	BDNotWorking              = errors.New("The database is down. Ping error")
	IncorrectParamsConnectBD  = errors.New("Error in database connection parameters")
	BadInsertUser             = errors.New("Failed to insert user")
	BadRoleUser               = errors.New("Failed get role user")
	BadUpdateUser             = errors.New("Failed update user")
	BadDeleteUser             = errors.New("Failed delete user")
	BadGetAllUsers            = errors.New("Failed get all users")
	FailedAddUser             = errors.New("Failed add user")
	TokenExpired              = errors.New("Срок действия токена истек")
	InvalidAccessToken        = errors.New("Вы не авторизованы")
	Forbidden                 = errors.New("No access rights to the resource")
	DeletingSelfError         = errors.New("You can't remove yourself")
	UpdatingSelfRoleError     = errors.New("You can't update role yourself")
	InvalidValueRole          = errors.New("Role value must be in the range [0,1]")
	UserNotFound              = errors.New("Пользователь не найден")
	FindUserDuplicate         = errors.New("Пользователь с таким логином уже существует")
	LenUsernameLessSixSymbols = errors.New("Длина имени пользователя должна быть больше 5 символов")
	LenPasswordLessSixSymbols = errors.New("Длина пароля должна быть больше 5 символов")
	BadRequest                = errors.New("Некорректное тело запроса")
)
