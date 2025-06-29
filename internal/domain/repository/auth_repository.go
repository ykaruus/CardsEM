package repository

type AuthRepository interface {
	Login(username, password string) (string, error)
}
