package services

import (
	"fmt"
	"kairusService/internal/domain/repository"
)

type AuthService struct {
	hashService  *HashService
	tokenService *TokenService
	userService  *UserService
}

func NewAuthServices(hashService *HashService, tokenService *TokenService, userService *UserService) *AuthService {
	return &AuthService{
		hashService:  hashService,
		tokenService: tokenService,
		userService:  userService,
	}
}

var _ repository.AuthRepository = (*AuthService)(nil) // implementa repository.AuthRepository

func (auth *AuthService) Login(username string, password string) (string, error) {
	foundedUser, err := auth.userService.GetUserFrom(username)

	if err != nil {
		return "", err
	}

	if foundedUser == nil {
		return "", fmt.Errorf("user_not_found")
	}

	isPasswordValid := auth.hashService.CheckHash(password, foundedUser.PasswordHash)

	fmt.Println(isPasswordValid)
	if !isPasswordValid {
		return "", fmt.Errorf("user_credential_invalid")
	}

	tokenString, err := auth.tokenService.CreateToken(foundedUser)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (auth *AuthService) DecriptToken(token string) (map[string]string, error) {
	payload, err := auth.tokenService.DecriptToken(token)

	if err != nil {
		return nil, err
	}

	if payload == nil {
		return nil, nil
	}

	return payload, nil
}

func (auth *AuthService) ExtractToken(token string) (string, error) {
	tokenString, err := auth.tokenService.ExtractToken(token)

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
