package models

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID    uint
	FirstName string
	LastName  string
	Email     string
	*jwt.StandardClaims
}