package storage

import (
	"fmt"
	"kairusService/internal/domain/entities"
	"kairusService/internal/domain/repository"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenRepository struct {
	secretkey []byte
}

var _ repository.TokenRepository = (*TokenRepository)(nil)

func NewTokenRepository(secretkey string) *TokenRepository {
	return &TokenRepository{
		secretkey: []byte(secretkey),
	}
}

func (t *TokenRepository) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return t.secretkey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func (t *TokenRepository) CreateToken(u *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Name,
		"role":     u.Role,
		"exp":      time.Now().Add(time.Hour * 10).Unix(),
	})

	tokenString, err := token.SignedString(t.secretkey)

	if err != nil {
		return "", nil
	}

	return tokenString, nil

}

func (t *TokenRepository) DecriptToken(token string) (map[string]string, error) {
	descriptedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secretkey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := descriptedToken.Claims.(jwt.MapClaims), descriptedToken.Valid

	if ok && descriptedToken.Valid {
		payload := make(map[string]string)
		payload["username"] = fmt.Sprint(claims["username"])
		payload["role"] = fmt.Sprint(claims["role"])
		return payload, nil
	}

	return nil, fmt.Errorf("token invalído")

}

func (t *TokenRepository) ExtractToken(token string) (string, error) {
	parts := strings.SplitN(token, " ", 2)

	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("token malformed")
	}

	fmt.Println(parts[1])

	return parts[1], nil
}
