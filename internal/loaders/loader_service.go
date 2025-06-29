package loaders

import (
	"kairusService/internal/middlewares"
	"kairusService/internal/services"
	"kairusService/internal/storage"
)

type LoaderService struct{}

func (loader *LoaderService) LoadTokenAndHashService(hashHepo *storage.HashStorage, tokenRepo *storage.TokenRepository) (*services.HashService, *services.TokenService) {
	hashService := services.NewHashService(hashHepo)
	tokenService := services.NewTokenService(tokenRepo)

	return hashService, tokenService
}

func (loader *LoaderService) LoadUserService(userStorage *storage.MongoUserRepository, hashHepo *services.HashService) *services.UserService {
	userService := services.NewUserService(userStorage, hashHepo)

	return userService

}

func (loader *LoaderService) LoadApiResponseService(apiResponseStorage *storage.ApiResponseStorage) *services.ApiResponseService {
	return services.NewApiResponseService(apiResponseStorage)
}

func (loader *LoaderService) LoadAuthMiddleware(tokenService *services.TokenService, apiResponseService *services.ApiResponseService) *middlewares.AuthMiddleware {

	return middlewares.NewAuthMiddleware(tokenService, apiResponseService)
}

func (loader *LoaderService) LoadAuthService(hashService *services.HashService, tokenService *services.TokenService, userService *services.UserService) *services.AuthService {
	return services.NewAuthServices(hashService, tokenService, userService)
}
