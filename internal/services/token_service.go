package services

import (
	"kairusService/internal/domain/entities"
	"kairusService/internal/domain/repository"
	"kairusService/internal/storage"
)

type TokenService struct {
	tokenRepo *storage.TokenRepository
}

var _ repository.TokenRepository = (*TokenService)(nil)

func NewTokenService(tokenRepo *storage.TokenRepository) *TokenService {
	return &TokenService{
		tokenRepo: tokenRepo,
	}
}

func (tokenService *TokenService) VerifyToken(token string) error {
	err := tokenService.tokenRepo.VerifyToken(token)

	if err != nil {
		return err
	}

	return nil
}

func (tokenService *TokenService) CreateToken(u *entities.User) (string, error) {
	tokenString, err := tokenService.tokenRepo.CreateToken(u)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (tokenService *TokenService) DecriptToken(token string) (map[string]string, error) {
	tokenClaims, err := tokenService.tokenRepo.DecriptToken(token)

	if err != nil {
		return nil, err
	}

	return tokenClaims, nil
}

func (tokenService *TokenService) ExtractToken(token string) (string, error) {
	tokenString, err := tokenService.tokenRepo.ExtractToken(token)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
