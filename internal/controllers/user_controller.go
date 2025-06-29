package controllers

import (
	"fmt"
	"kairusService/internal/domain/entities"
	"kairusService/internal/services"
	"kairusService/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service    *services.UserService
	apiService *services.ApiResponseService
}

func NewUserController(userService *services.UserService, apiService *services.ApiResponseService) *UserController {
	return &UserController{
		service:    userService,
		apiService: apiService,
	}
}

func (user *UserController) Create(c *gin.Context) {
	var newUser entities.UserRequest
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(400, user.apiService.CreateReponseModel(
			false,
			utils.StatusMessages.BadRequest,
			utils.StatusCodes.BadRequest,
			utils.INVALID_REQUEST_BODY,
			nil,
			nil,
		))
		return
	}

	insertId, err := user.service.CreateUser(newUser)

	if err != nil {
		if err.Error() == "duplicated key error" {
			c.JSON(400, user.apiService.CreateReponseModel(
				false,
				utils.StatusMessages.Conflict,
				utils.StatusCodes.Conflict,
				utils.RESOURCE_CONFLICT,
				nil,
				nil,
			))
			return
		}
		c.JSON(500, user.apiService.CreateReponseModel(
			false,
			utils.StatusMessages.InternalError,
			utils.StatusCodes.InternalError,
			utils.INTERNAL_ERROR,
			nil,
			nil,
		))
		log.Default().Println(err)
		return
	}

	c.JSON(200, user.apiService.CreateReponseModel(
		true,
		utils.StatusMessages.Success,
		utils.StatusCodes.Ok,
		utils.RESOURCE_SUCCESS_CREATED,
		gin.H{
			"user_id": insertId,
		},
		nil,
	))

}

func (user *UserController) GetUserFromId(c *gin.Context) {
	id := c.Param("id")

	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"message": "Id Invalido!, verifique o ID e tente novamente",
	// 	})
	// 	rurn
	// }

	err := user.service.CheckIdString(id)
	if err != nil {
		c.JSON(400, user.apiService.CreateReponseModel(
			true,
			utils.StatusMessages.InvalidId,
			utils.StatusCodes.BadRequest,
			utils.INVALID_REQUEST_BODY,
			nil,
			nil,
		))
		return
	}
	foundUser, err := user.service.GetUser(id)

	if err != nil {
		c.JSON(400, user.apiService.CreateReponseModel(
			true,
			utils.StatusMessages.BadRequest,
			utils.StatusCodes.BadRequest,
			utils.RESOURCE_NOT_FOUND,
			nil,
			nil,
		))
		return
	}

	if foundUser == nil {
		c.JSON(404, user.apiService.CreateReponseModel(
			true,
			utils.StatusMessages.NotFound,
			utils.StatusCodes.NotFound,
			utils.RESOURCE_NOT_FOUND,
			nil,
			nil,
		))
		return
	}

	c.JSON(200, user.apiService.CreateReponseModel(
		true,
		utils.StatusMessages.Success,
		utils.StatusCodes.Ok,
		utils.SUCCESS,
		gin.H{
			"user": gin.H{
				"name":    foundUser.Name,
				"role":    foundUser.Role,
				"user_id": foundUser.ID.Hex(),
			},
			"created_at": foundUser.CreatedAt,
			"updated_at": foundUser.UpdatedAt,
		},
		nil,
	))
}

func (user *UserController) GetUserFromName(c *gin.Context) {
	name := c.Query("name")

	fmt.Println(name)

	foundUser, err := user.service.GetUserFrom(name)

	if err != nil {
		c.JSON(400, user.apiService.CreateReponseModel(
			true,
			utils.StatusMessages.BadRequest,
			utils.StatusCodes.BadRequest,
			utils.RESOURCE_NOT_FOUND,
			nil,
			nil,
		))
		return
	}

	if foundUser == nil {
		c.JSON(404, user.apiService.CreateReponseModel(
			true,
			utils.StatusMessages.NotFound,
			utils.StatusCodes.NotFound,
			utils.RESOURCE_NOT_FOUND,
			nil,
			nil,
		))
		return
	}

	c.JSON(200, user.apiService.CreateReponseModel(
		true,
		utils.StatusMessages.Success,
		utils.StatusCodes.Ok,
		utils.SUCCESS,
		gin.H{
			"user": gin.H{
				"name":    foundUser.Name,
				"role":    foundUser.Role,
				"user_id": foundUser.ID.Hex(),
			},
			"created_at": foundUser.CreatedAt,
			"updated_at": foundUser.UpdatedAt,
		},
		nil,
	))
}

func (user *UserController) GetUsers(c *gin.Context) {
	users, err := user.service.GetAllUser()

	if err != nil {
		c.JSON(500, user.apiService.CreateReponseModel(
			true,
			utils.StatusMessages.InternalError,
			utils.StatusCodes.InternalError,
			utils.INTERNAL_ERROR,
			nil,
			nil,
		))
		return
	}

	if users == nil {
		c.JSON(404, user.apiService.CreateReponseModel(
			true,
			utils.StatusMessages.NotFound,
			utils.StatusCodes.NotFound,
			utils.RESOURCE_NOT_FOUND,
			nil,
			nil,
		))
		return
	}

	c.JSON(200, user.apiService.CreateReponseModel(
		true,
		utils.StatusMessages.Success,
		utils.StatusCodes.Ok,
		utils.SUCCESS,
		users,
		nil,
	))
}

func (user *UserController) UpdateUser(c *gin.Context) {
	var updatedUser entities.UserRequest

	id := c.Param("id")

	id_err := user.service.CheckIdString(id)

	if id_err != nil {
		c.JSON(400, user.apiService.CreateReponseModel(
			false,
			utils.StatusMessages.InvalidId,
			utils.StatusCodes.BadRequest,
			utils.INVALID_REQUEST_BODY,
			nil,
			nil,
		))
	}

	erro := c.ShouldBindJSON(&updatedUser)

	if erro != nil {
		c.JSON(400, user.apiService.CreateReponseModel(
			false,
			utils.StatusMessages.BadRequest,
			utils.StatusCodes.BadRequest,
			utils.INVALID_REQUEST_BODY,
			nil,
			nil,
		))
		return
	}

	updatedUserResultId, err := user.service.UpdateUser(id, updatedUser)

	if err != nil && err.Error() == "duplicated key error" {
		c.JSON(400, user.apiService.CreateReponseModel(
			false,
			utils.StatusMessages.Conflict,
			utils.StatusCodes.BadRequest,
			utils.RESOURCE_CONFLICT,
			nil,
			nil,
		))
		return
	}

	if err != nil {
		c.JSON(500, user.apiService.CreateReponseModel(
			false,
			utils.StatusMessages.InternalError,
			utils.StatusCodes.InternalError,
			utils.INTERNAL_ERROR,
			nil,
			nil,
		))

		log.Fatal(err)
		return
	}

	c.JSON(200, user.apiService.CreateReponseModel(
		false,
		utils.StatusMessages.ResourceUpdated,
		utils.StatusCodes.Ok,
		utils.SUCCESS,
		gin.H{
			"user_id": updatedUserResultId,
		},
		nil,
	))

}
