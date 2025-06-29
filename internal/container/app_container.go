package container

import (
	"kairusService/internal/loaders"
	"kairusService/internal/services"
	"kairusService/internal/storage"
)

type Container struct {
	UserService  *services.UserService
	TokenService *services.TokenService
	HashService  *services.HashService
	AuthService  *services.AuthService
	ApiResponse  *services.ApiResponseService
}

func NewContainer(
	secretKey string,
	userStorage *storage.MongoUserRepository,

) *Container {

	// Storages
	loaderStorage := storage.NewLoaderStorage()
	loaderService := loaders.LoaderService{}
	hashStorage, tokenStorage := loaderStorage.LoadTokenAndHashStorage(secretKey)

	apiResponseStorage := storage.NewApiResponseStorage()
	apiResponse := services.NewApiResponseService(apiResponseStorage)

	// injeta os storages nos services
	// Services:
	hashService, tokenService := loaderService.LoadTokenAndHashService(hashStorage, tokenStorage)

	userService := loaderService.LoadUserService(userStorage, hashService)

	authService := loaderService.LoadAuthService(hashService, tokenService, userService)

	return &Container{
		UserService:  userService,
		TokenService: tokenService,
		HashService:  hashService,
		AuthService:  authService,
		ApiResponse:  apiResponse,
	}
}
