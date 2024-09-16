package controller

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/config"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/service"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/utils"
)

var env_secretKey = os.Getenv("SECRET_KEY")
var secretKey = []byte(env_secretKey)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusBadRequest)
		return
	}
	userData := &models.User{}
	utils.ParseJSONBody(r, userData)

	//One DB Conn
	db := config.InitializeDB()
	var dataFromDB models.User //Read DB data from here
	errr := db.Where("username = ?", userData.Username).Limit(1).Find(&dataFromDB).Error
	if errr != nil {
		log.Println("Error fetching data:", errr)
		return
	}
	expiryTime := time.Now().Add(time.Minute * 60).Unix()

	if service.ValidateCredentials(db, userData, &dataFromDB) {
		jwtToken, err := service.GenJWT(&dataFromDB, expiryTime, secretKey)
		if err != nil {
			log.Println("Generate token error !")
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Authorization", "Bearer "+jwtToken)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

}

func ValidateUser(w http.ResponseWriter, r *http.Request) {

	tokenheader := r.Header.Get("Authorization")
	jwtToken := strings.Fields(tokenheader)[1]
	userData := &models.ValidationResponse{}
	utils.ParseJSONBody(r, userData)

	req_username := userData.Username
	if req_username == "" {
		log.Println("Error: Missing required fields in request body")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if service.ValidateToken(jwtToken, string(secretKey), req_username) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This token is Validated"))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("This token is Invalid ! Login to get a new Token"))
	}

}
