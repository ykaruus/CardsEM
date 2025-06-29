package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongo(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}
	fmt.Println("Banco conectado com sucesso.")

	return client, nil
}

func LoadCollection(client *mongo.Client, dbName string, collName string) *mongo.Collection {
	coll := client.Database(dbName).Collection(collName)

	return coll
}
