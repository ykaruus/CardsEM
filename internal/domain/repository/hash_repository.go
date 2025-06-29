package repository

type HashRepository interface {
	CreateHash(password string) (string, error)
	CheckHash(password string, hash string) bool
}
