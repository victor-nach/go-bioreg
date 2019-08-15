package routes

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"go-bioreg/server/models/app"
	user "go-bioreg/server/controllers/user"
)

// MyRouter - contains all the routes in the app
func MyRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/signup", user.SignUp).Methods("POST")
	router.HandleFunc("/user", user.GetAllUsers).Methods("GET")
	router.HandleFunc("/", welcome)
	return router
}


func welcome(w http.ResponseWriter, r *http.Request) {
	response := app.StatusResponse{
		Status: 200,
		Message: "Welcome to bioreg backend v1",
	}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}