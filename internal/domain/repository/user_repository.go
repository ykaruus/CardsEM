package repository

import (
	"kairusService/internal/domain/entities"
)

type UserStorageRepository interface {
	// CreateUser(u User) error
	GetUser(id string) (*entities.User, error)
	GetAllUser() ([]entities.User, error)
	CreateUser(u entities.UserRequest) (string, error)
	UpdateUser(id string, u entities.UserRequest) (string, error)
	GetUserFrom(name string) (*entities.User, error)

	// UpdateUser(id string) (*User, error)
	DeleteUser(id string) (string, error)

	CheckIdString(id string) error
}

type UserServiceRepository interface {
	// CreateUser(u User) error
	GetUser(id string) (*entities.User, error)
	GetAllUser() ([]entities.User, error)
	CreateUser(u entities.UserRequest) (string, error)
	UpdateUser(id string, u entities.UserRequest) (string, error)
	GetUserFrom(name string) (*entities.User, error)

	// UpdateUser(id string) (*User, error)
	DeleteUser(id string) (string, error)

	CheckIdString(id string) error
}
