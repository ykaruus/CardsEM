package repository

import "github.com/gin-gonic/gin"

type ApiResponseRepository interface {
	CreateReponseModel(sucess bool, message string, status_code int, code string, data interface{}, errors interface{}) *gin.H
}
