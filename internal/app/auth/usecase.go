package auth

type UseCase interface {
	SignUp(username, password string) (token string, err error)
	SignIn(username, password string) (token string, err error)
}
