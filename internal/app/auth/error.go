package auth

import "errors"

var (
	UserNotFound              = errors.New("Пользователь не найден")
	BDNotWorking              = errors.New("The database is down. Ping error")
	IncorrectParamsConnectBD  = errors.New("Error in database connection parameters")
	BadInsertUser             = errors.New("Failed to insert user")
	BadRoleUser               = errors.New("Failed get role user")
	BadSqlRequest             = errors.New("Bad sql request")
	FailedSignUp              = errors.New("Ошибка сервера: не удалось зарегистрироваться, попробуйте позже")
	FailedSignIn              = errors.New("Ошибка сервера: не удалось войти, попробуйте позже")
	FindUserDuplicate         = errors.New("Пользователь с таким логином уже существует")
	FailedGenToken            = errors.New("Ошибка генерации токена")
	LenUsernameLessSixSymbols = errors.New("Длина имени пользователя даолжна быть больше 5 символов")
	LenPasswordLessSixSymbols = errors.New("Длина пароля должна быть больше 5 символов")
	BadRequest                = errors.New("Некорректное тело запроса")
	InvalidLoginOrPassword    = errors.New("Неверный логин или пароль")
)
