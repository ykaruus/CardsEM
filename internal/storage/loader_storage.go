package storage

import "go.mongodb.org/mongo-driver/v2/mongo"

type LoaderStorage struct{}

func NewLoaderStorage() *LoaderStorage {
	return &LoaderStorage{}
}
func (loader *LoaderStorage) LoadTokenAndHashStorage(secretkey string) (*HashStorage, *TokenRepository) {
	tokenStorage := NewTokenRepository(secretkey)

	return &HashStorage{}, tokenStorage
}

func (loader *LoaderStorage) LoadUserStorage(coll *mongo.Collection) *MongoUserRepository {
	return NewMongoRepository(coll)
}

func (loader *LoaderStorage) LoadApiResponseStorage() *ApiResponseStorage {
	return &ApiResponseStorage{}
}
