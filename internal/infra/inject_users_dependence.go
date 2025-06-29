package infra

import (
	"kairusService/internal/controllers"
	"kairusService/internal/services"
)

func InjectUsersServices(userService *services.UserService, apiResponseService *services.ApiResponseService, authService *services.AuthService) (*controllers.UserController, *controllers.AuthController) {
	return controllers.NewUserController(userService, apiResponseService), controllers.NewAuthController(authService, apiResponseService)
}
