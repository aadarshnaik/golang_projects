package controller

import (
	"net/http"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/service"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/utils"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusBadRequest)
		return
	}

	createUser := &models.User{}
	utils.ParseJSONBody(r, createUser)

	err := service.CreateUser(createUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
