package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/register", CreateNewUser).Methods("POST")

	log.Printf("Server starting at 9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
