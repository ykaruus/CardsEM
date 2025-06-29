package storage

import (
	"kairusService/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

type ApiResponseStorage struct{}

func NewApiResponseStorage() *ApiResponseStorage {
	return &ApiResponseStorage{}
}

var _ repository.ApiResponseRepository = (*ApiResponseStorage)(nil)

func (apiResponse *ApiResponseStorage) CreateReponseModel(sucess bool, message string, status_code int, code string, data interface{}, errors interface{}) *gin.H {
	return &gin.H{
		"success": sucess,
		"message": message,
		"status":  status_code,
		"code":    code,
		"data":    data,
		"errors":  errors,
	}
}
