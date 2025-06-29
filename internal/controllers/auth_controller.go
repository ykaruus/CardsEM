package controllers

import (
	"fmt"
	"kairusService/internal/domain/entities"
	"kairusService/internal/services"
	"kairusService/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService        *services.AuthService
	apiResponseService *services.ApiResponseService
}

func NewAuthController(auth *services.AuthService, apiResponseService *services.ApiResponseService) *AuthController {
	return &AuthController{
		authService:        auth,
		apiResponseService: apiResponseService,
	}
}

func (auth *AuthController) Login(ctx *gin.Context) {
	var user entities.UserLogin

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(400, auth.apiResponseService.CreateReponseModel(
			false,
			utils.StatusMessages.BadRequest,
			utils.StatusCodes.BadRequest,
			utils.INVALID_REQUEST_BODY,
			nil,
			gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	tokenString, err := auth.authService.Login(user.Name, user.Password)

	if err != nil {
		ctx.JSON(401, auth.apiResponseService.CreateReponseModel(
			false,
			utils.StatusMessages.Unauthorized,
			utils.StatusCodes.Unauthorized,
			utils.UNAUTHORIZED_ACCESS,
			nil,
			gin.H{
				"error": err.Error(),
			},
		))
		fmt.Println(err)
		return
	}

	ctx.JSON(200, auth.apiResponseService.CreateReponseModel(
		true,
		utils.StatusMessages.Authorized,
		utils.StatusCodes.Ok,
		utils.AUTHORIZED_ACCESS, gin.H{
			"token": tokenString,
		}, nil))

}

func (auth *AuthController) DecriptToken(ctx *gin.Context) {
	var token string = ctx.GetHeader("Authorization")

	tokenString, err := auth.authService.ExtractToken(token)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "erro : " + err.Error(),
		})
	}
	payload, err := auth.authService.DecriptToken(tokenString)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "erro : " + err.Error(),
		})
		return
	}

	if payload == nil {
		ctx.JSON(500, gin.H{
			"message": "Data não recebida",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": payload,
	})

}
