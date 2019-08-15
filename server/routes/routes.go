package routes

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"go-bioreg/server/models/app"
)

// MyRouter - contains all the routes in the app
func MyRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcome)
	return router
}


func welcome(w http.ResponseWriter, r *http.Request) {
	response := app.StatusResponse{200, "Welcome to bioreg backend v1"}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}