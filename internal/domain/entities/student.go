package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Card struct {
	Name      string        `bson:"name" json:"name"`
	Content   string        `bson:"content" json:"content"`
	Subject   string        `bson:"subject" json:"subject"`
	Card_ID   bson.ObjectID `bson:"__id" json:"card_id"`
	User_ID   bson.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
	V         int           `bson:"__v" json:"__v"`
}

type CardRequest struct {
	Name      string        `bson:"name" json:"name"`
	Content   string        `bson:"content" json:"content"`
	Subject   string        `bson:"subject" json:"subject"`
	Card_ID   bson.ObjectID `bson:"_id" json:"card_id"`
	User_ID   bson.ObjectID `bson:"user_id" json:"user_id"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
	V         int           `bson:"__v" json:"__v"`
}
