package user

import (
	"log"
	"errors"

	"go-bioreg/server/database/config"
	"go-bioreg/server/models/db"
)

// SignUp - creates a new user entry in the database
func SignUp(user db.User) error {
	db := config.GetConn()
	result := db.Create(&user)
	if result.Error != nil {
		log.Println("could not insert user")
		return errors.New("could not insert user")
	}
	return nil
}

// GetAllUsers - creates a new user entry in the database
func GetAllUsers() ([]db.User, error) {
	var users []db.User
	db := config.GetConn()
	result := db.Find(&users)
	if result.Error != nil {
		log.Println("could not get users")
		return nil, errors.New("could not get users")
	}
	return users, nil
}