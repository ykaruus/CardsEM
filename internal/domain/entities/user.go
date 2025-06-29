package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Name         string        `bson:"username" json:"username"`
	Role         string        `bson:"role" json:"role"`
	PasswordHash string        `bson:"passwordHash,omitempty" json:"passwordHash,omitempty"`
	ID           bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt" json:"updatedAt"`
	V            int           `bson:"__v" json:"__v"`
}

type UserRequest struct {
	Name     string `json:"username" binding:"required,min=3,max=20"`
	Role     string `json:"role" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
