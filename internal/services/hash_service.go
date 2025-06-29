package services

import (
	"kairusService/internal/domain/repository"
	"kairusService/internal/storage"
)

type HashService struct {
	hashRepository *storage.HashStorage
}

var _ repository.HashRepository = (*HashService)(nil)

func NewHashService(auth *storage.HashStorage) *HashService {
	return &HashService{
		hashRepository: auth,
	}
}

func (auth *HashService) CheckHash(password string, passwordHash string) bool {
	err := auth.hashRepository.CheckHash(password, passwordHash)
	return err
}

func (auth *HashService) CreateHash(password string) (string, error) {
	passwordHash, err := auth.hashRepository.CreateHash(password)

	return passwordHash, err
}
