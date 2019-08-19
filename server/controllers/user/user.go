package usercontroller

import (
	"net/http"
	"log"
	"encoding/json"
	"go-bioreg/server/models/db"
	"go-bioreg/server/models/app"
	userModel "go-bioreg/server/database/dao/user"

)

//SignUp controller to sign up a user
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(user)
		response := app.UserResponse {
			Status: 200,
			Message: err.Error(),
	
			}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		log.Println(err.Error())	
	}
	err = userModel.SignUp(user)
	if err != nil {
		response := app.UserResponse {
			Status: 200,
			Message: err.Error(),
	
			}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		log.Println("could not rereive users")		
	}
	response := app.UserResponse {
		Status: 200,
		Message: "success",
		Data: user,

	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllUsers - returns all the users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userModel.GetAllUsers()
	if err != nil {
		response := app.UserResponse {
		Status: 200,
		Message: err.Error(),

		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		log.Println("could not rereive users")
	}
	// on success
	response := app.AllUsersResponse {
		Status: 200,
		Message: "success",
		Data: users,

	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SignIn - checks a user's credentials and signs the user in
func SignIn(w http.ResponseWriter, r http.Request) {
	
}
