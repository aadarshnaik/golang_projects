package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/config"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/service"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/utils"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusBadRequest)
		return
	}

	db := config.InitializeDB()

	createUser := &models.User{}
	utils.ParseJSONBody(r, createUser)

	err := service.CreateUser(createUser, db)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		endtime := time.Since(start)
		log.Println("Time taken:", endtime)
		w.WriteHeader(http.StatusOK)
	}
}
