package repository

import (
	"kairusService/internal/domain/entities"
)

type TokenRepository interface {
	VerifyToken(token string) error
	CreateToken(u *entities.User) (string, error)
	DecriptToken(token string) (map[string]string, error)
	ExtractToken(token string) (string, error)
}
