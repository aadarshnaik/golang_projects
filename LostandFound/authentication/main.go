package main

import (
	"log"
	"net/http"

	controller "github.com/aadarshnaik/golang_projects/LostandFound/authentication/controllers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/register", controller.CreateNewUser).Methods("POST")
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")
	r.HandleFunc("/validate", controller.ValidateUser).Methods("GET")

	log.Printf("Server starting at 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
