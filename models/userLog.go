package models

// UserLog ...
type UserLog struct {
	Created   string `json:"created,omitempty" bson:"created,omitempty"`
	LastLogin string `json:"lastLogin,omitempty" bson:"lastLogin,omitempty"`
}
