package controllers

import (
	"auth/utils"
	"net/http"
)

type error interface {
	Error() string
}

var db = utils.ConnectDB()

// Login function
func Login(w http.ResponseWriter, r *http.Request) {

}

// FindOne function
func FindOne(email, password string) map[string]interface{} {

}

//CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

//FetchUsers function
func FetchUsers(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser function
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser function
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// GetUser function
func GetUser(w http.ResponseWriter, r *http.Request) {

}
