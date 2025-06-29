package services

import (
	"kairusService/internal/domain/entities"
	"kairusService/internal/domain/repository"
	"kairusService/internal/storage"
)

type UserService struct {
	mongo *storage.MongoUserRepository
	hashS *HashService
}

func NewUserService(r *storage.MongoUserRepository, hash *HashService) *UserService {
	return &UserService{
		mongo: r,
		hashS: hash,
	}
}

var _ repository.UserServiceRepository = (*UserService)(nil)

func (u *UserService) GetUser(id string) (*entities.User, error) {

	user, err := u.mongo.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *UserService) GetAllUser() ([]entities.User, error) {
	users, err := u.mongo.GetAllUser()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserService) CreateUser(user entities.UserRequest) (string, error) {

	passwordHash, err := u.hashS.CreateHash(user.Password)

	if err != nil {
		return "", err
	}

	newUser := entities.UserRequest{
		Password: passwordHash,
		Role:     user.Role,
		Name:     user.Name,
	}
	newUserId, err := u.mongo.CreateUser(newUser)

	if err != nil {
		return "", err
	}

	return newUserId, nil
}

func (u *UserService) UpdateUser(id string, user entities.UserRequest) (string, error) {
	newPasswordHash, err := u.hashS.CreateHash(user.Password)

	if err != nil {
		return "", err
	}
	updatedUserId, err := u.mongo.UpdateUser(id, entities.UserRequest{
		Name:     user.Name,
		Role:     user.Role,
		Password: newPasswordHash,
	})
	if err != nil {
		return "", err
	}

	return updatedUserId, nil
}

func (u *UserService) GetUserFrom(name string) (*entities.User, error) {
	FoundUser, err := u.mongo.GetUserFrom(name)

	if err != nil {
		return nil, err
	}

	if FoundUser == nil {
		return nil, nil
	}

	return FoundUser, nil

}

func (u *UserService) DeleteUser(id string) (string, error) {
	deletedUserId, err := u.mongo.DeleteUser(id)

	if err != nil {
		return "", nil
	}

	return deletedUserId, nil
}

func (u *UserService) CheckIdString(id string) error {
	err := u.mongo.CheckIdString(id)

	if err != nil {
		return err
	}

	return nil
}
