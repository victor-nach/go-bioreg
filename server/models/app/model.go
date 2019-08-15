package app

import (
	"go-bioreg/server/models/db"
)

// StatusResponse - standard status response
type StatusResponse struct {
	Status int `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// UserResponse - standard response for users
type UserResponse struct {
	Status 			int 		`json:"status,omitempty"`
	Message 		string 		`json:"message,omitempty"`
	Data 			db.User		`json:"data,omitempty"`
}

// AllUsersResponse - Standard user response for all users
type AllUsersResponse struct {
	Status 			int 		`json:"status,omitempty"`
	Message 		string 		`json:"message,omitempty"`
	Data 			[]db.User		`json:"data,omitempty"`
}

