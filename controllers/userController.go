package controllers

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/mendezandrewm/monGo-UserDB/connection"
)

var UserCollection *mongo.Collection = connection.OpenCollection(connection.ConnectDB(), "user")
var Validate = validator.New()

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		check = false
	}

	return check, msg
}

//FetchUser(s) function
func FetchUsers() {
}

// UpdateUser function
func UpdateUser() {
}

// DeleteUser function
func DeleteUser() {
}
