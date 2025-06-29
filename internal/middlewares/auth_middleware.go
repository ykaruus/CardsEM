package middlewares

import (
	"fmt"
	"kairusService/internal/services"
	"kairusService/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenService       *services.TokenService
	apiResponseService *services.ApiResponseService
}

func NewAuthMiddleware(tokenService *services.TokenService, apiResponseService *services.ApiResponseService) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService:       tokenService,
		apiResponseService: apiResponseService,
	}
}

func (middle *AuthMiddleware) CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		tokenString, err := middle.tokenService.ExtractToken(token)

		if err != nil {
			ctx.JSON(401, middle.apiResponseService.CreateReponseModel(
				false,
				utils.StatusMessages.Token_malformed,
				utils.StatusCodes.Unauthorized,
				utils.INVALID_TOKEN,
				nil,
				nil,
			))

			return
		}

		token_error := middle.tokenService.VerifyToken(tokenString)

		if token_error != nil {
			ctx.AbortWithStatusJSON(401, middle.apiResponseService.CreateReponseModel(
				false,
				utils.StatusMessages.Unauthorized,
				utils.StatusCodes.Unauthorized,
				utils.INVALID_TOKEN,
				nil,
				nil,
			))

			return
		}

		ctx.Next()

	}
}

func (middle *AuthMiddleware) CheckAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		tokenString, err := middle.tokenService.ExtractToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(401, middle.apiResponseService.CreateReponseModel(
				false,
				utils.StatusMessages.Token_malformed,
				utils.StatusCodes.Unauthorized,
				utils.INVALID_TOKEN,
				nil,
				nil,
			))

			return
		}

		token_error := middle.tokenService.VerifyToken(tokenString)

		if token_error != nil {
			ctx.AbortWithStatusJSON(401, middle.apiResponseService.CreateReponseModel(
				false,
				utils.StatusMessages.Unauthorized,
				utils.StatusCodes.Unauthorized,
				utils.INVALID_TOKEN,
				nil,
				nil,
			))

			return
		}

		payload, err := middle.tokenService.DecriptToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(401, middle.apiResponseService.CreateReponseModel(
				false,
				utils.StatusMessages.Unauthorized,
				utils.StatusCodes.Unauthorized,
				utils.INVALID_TOKEN,
				nil,
				nil,
			))

			return
		}

		fmt.Println(payload)

		if payload["role"] != "admin" {
			ctx.AbortWithStatusJSON(401, middle.apiResponseService.CreateReponseModel(
				false,
				utils.StatusMessages.Unauthorized,
				utils.StatusCodes.Unauthorized,
				utils.UNAUTHORIZED_ACCESS,
				nil,
				nil,
			))

			return
		}

		ctx.Next()

	}
}
