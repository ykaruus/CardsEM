package storage

import (
	"context"
	"errors"
	"fmt"
	"kairusService/internal/domain/entities"
	"kairusService/internal/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(coll *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{
		collection: coll,
	}
}

var _ repository.UserStorageRepository = (*MongoUserRepository)(nil)

func (r *MongoUserRepository) GetUser(id string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var result entities.User

	idObj, idObjErr := bson.ObjectIDFromHex(id)

	if idObjErr != nil {
		return nil, idObjErr
	}

	opts := options.FindOne().SetProjection(bson.M{"password": 0})
	err := r.collection.FindOne(ctx, bson.M{"_id": idObj}, opts).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNilDocument) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil

}
func (r *MongoUserRepository) GetAllUser() ([]entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	opts := options.Find().SetProjection(bson.M{"password": 0})
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)

	if err != nil {
		return []entities.User{}, err
	}
	defer cursor.Close(ctx)

	var results []entities.User

	err = cursor.All(ctx, &results)

	if err != nil {
		return nil, err
	}

	return results, nil

}

func (r *MongoUserRepository) CreateUser(u entities.UserRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	now := time.Now()

	result, err := r.collection.InsertOne(ctx, bson.M{
		"username":     u.Name,
		"role":         u.Role,
		"passwordHash": u.Password,
		"createdAt":    now,
		"updatedAt":    now,
	})

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return "", fmt.Errorf("duplicated key error")
		}
		return "", err
	}

	return result.InsertedID.(bson.ObjectID).Hex(), nil
}

func (r *MongoUserRepository) UpdateUser(id string, u entities.UserRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	idObj, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return "", err
	}

	now := time.Now()

	update := bson.D{{Key: "$set", Value: bson.M{
		"username":     u.Name,
		"role":         u.Role,
		"passwordHash": u.Password,
		"updatedAt":    now,
	}}}
	var updatedUser entities.User

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	updated_err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": idObj}, update, opts).Decode(&updatedUser)

	if updated_err != nil {
		if mongo.IsDuplicateKeyError(updated_err) {
			return "", fmt.Errorf("duplicated key error")
		}
		return "", err
	}

	return updatedUser.ID.Hex(), nil

}

func (r *MongoUserRepository) GetUserFrom(name string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"username": name}
	var foundUser *entities.User
	err := r.collection.FindOne(ctx, filter).Decode(&foundUser)

	if err != nil {
		if errors.Is(err, mongo.ErrNilDocument) {
			return nil, nil
		}
		return nil, err
	}

	return foundUser, nil
}

func (r *MongoUserRepository) DeleteUser(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	idObj, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return "", err
	}

	filter := bson.M{"_id": idObj}
	var deletedUser entities.User
	errResult := r.collection.FindOneAndDelete(ctx, filter).Decode(&deletedUser)

	if errResult != nil {
		if errors.Is(err, mongo.ErrNilDocument) {
			return "", nil
		}
	}

	return deletedUser.ID.Hex(), nil

}

func (r *MongoUserRepository) CheckIdString(id string) error {
	_, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	return nil
}
