package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct declaration
type User struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn string             `json:"isbn,omitempty" bson:"isbn,omitempty"`

	UserName string `json:"userName,omitempty" bson:"userName,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`

	Profile *Profile `json:"profile,omitempty" bson:"profile,omitempty"`
	UserLog *UserLog `json:"userLog,omitempty" bson:"userLog,omitempty"`
}
