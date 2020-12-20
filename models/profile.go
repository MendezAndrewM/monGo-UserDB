package models

// Profile ...
type Profile struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	DOB       string `json:"DOB,omitempty" bson:"DOB,omitempty"`
}
