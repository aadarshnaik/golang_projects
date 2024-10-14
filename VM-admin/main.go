package main

import (
	"log"
	"net/http"

	"github.com/aadarshnaik/vm-admin/controllers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/heartbeat", controllers.GetHeartbeat).Methods("POST")

	log.Println("Waiting for Heartbeat on port 4040")
	log.Fatal(http.ListenAndServe(":4040", r))
}
