package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/config"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/service"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/utils"
)

var secretKey = []byte("mysecretstring")

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
	// testFromDB := interface{}
	errr := db.Where("username = ?", userData.Username).Limit(1).Find(&dataFromDB).Error
	log.Println("CreatedAt: ", dataFromDB.CreatedAt)
	log.Println("CreatedAt: ", dataFromDB.UpdatedAt)
	if errr != nil {
		log.Println("Error fetching data:", errr)
		return
	}
	expiryTime := time.Now().Add(time.Hour * 24).Unix()

	if service.ValidateCredentials(db, userData, &dataFromDB) {
		jwtToken := service.GenJWT(&dataFromDB, expiryTime, secretKey)
		w.Header().Set("Authorization", "Bearer "+jwtToken)

		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}

	// JWT Token Generation
	// var user models.User

	// log.Println("User Authorised")

	// type res struct {
	// 	Token      string `json:"token"`
	// 	ExpiryTime int64  `json:"expiryTime"`
	// }

	// t := &res{Token: jwtToken, ExpiryTime: expiryTime}

	// if out, err := json.MarshalIndent(t, "", " "); err != nil {
	// 	http.Error(w, "Error marshalling response", http.StatusInternalServerError)
	// 	return
	// } else {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(out)
	// }

}

func ValidateUser(w http.ResponseWriter, r *http.Request) {

	tokenheader := r.Header.Get("Authorization")
	jwtToken := strings.Fields(tokenheader)[1]

	if service.ValidateToken(jwtToken, string(secretKey)) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This token is Validated"))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("This token is Invalid ! Login to get a new Token"))
	}

}
