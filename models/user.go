package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string             `bson:"firstName" json:"first_name"`
	LastName  string             `bson:"lastName" json:"last_name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Status    bool               `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"created_at"`
	OTP       OTP                `bson:"otp" json:"otp"`
}

func NewUser(id primitive.ObjectID, firstName, lastName, email, password string, status bool) *User {

	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Status:    status,
		CreatedAt: time.Now(),
	}
}
