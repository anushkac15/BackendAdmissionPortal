package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	ZipCode string `bson:"zipCode" json:"zipCode"`
	Country string `bson:"country" json:"country"`
}

type Student struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email       string             `bson:"email" json:"email" binding:"required,email"`
	Password    string             `bson:"password" json:"password" binding:"required,min=6"`
	Name        string             `bson:"name" json:"name" binding:"required"`
	Phone       string             `bson:"phone" json:"phone"`
	DateOfBirth string             `bson:"dateOfBirth" json:"dateOfBirth"`
	Gender      string             `bson:"gender" json:"gender"`
	Address     Address            `bson:"address" json:"address"`
	Role        string             `bson:"role" json:"role"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
