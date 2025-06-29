package services

import (
	"kairusService/internal/domain/repository"
	"kairusService/internal/storage"

	"github.com/gin-gonic/gin"
)

type ApiResponseService struct {
	apiResponseStorage *storage.ApiResponseStorage
}

func NewApiResponseService(apiResponseStorage *storage.ApiResponseStorage) *ApiResponseService {
	return &ApiResponseService{
		apiResponseStorage: apiResponseStorage,
	}
}

var _ repository.ApiResponseRepository = (*ApiResponseService)(nil)

func (apiService *ApiResponseService) CreateReponseModel(sucess bool, message string, status_code int, code string, data interface{}, errors interface{}) *gin.H {
	return apiService.apiResponseStorage.CreateReponseModel(sucess, message, status_code, code, data, errors)
}
