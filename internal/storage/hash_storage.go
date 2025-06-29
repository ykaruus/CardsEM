package storage

import (
	"kairusService/internal/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type HashStorage struct{}

var _ repository.HashRepository = (*HashStorage)(nil)

func NewHashRepo() *HashStorage {
	return &HashStorage{}
}

func (hash *HashStorage) CreateHash(password string) (string, error) {
	passwordBytes := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(passwordBytes, 7)

	return string(passwordHash), err
}

func (hash *HashStorage) CheckHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
