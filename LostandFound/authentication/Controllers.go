package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/config"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	Username     string         `json:"username" gorm:"primaryKey; not null; unique"`
	Passwordhash string         `json:"passwordhash" gorm:"not null"`
	Pincode      int            `json:"pincode" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func userExists(user *User) bool {
	db := config.InitializeDB()
	err := db.Where("username = ?", user.Username).First(&user).Error
	return err == nil
	// res, _ := json.MarshalIndent(user, "", " ")
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusBadRequest)
		return
	}

	createUser := &User{}
	utils.ReadRequestBody(r, createUser)

	passwordBytes := []byte(createUser.Passwordhash)
	encodedPassword := base64.StdEncoding.EncodeToString(passwordBytes)
	createUser.Passwordhash = encodedPassword

	db := config.InitializeDB()
	db.AutoMigrate(&User{})
	if userExists(createUser) {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	err := db.Create(&createUser).Error
	if err != nil {
		fmt.Println("Error creating user:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	log.Printf("New User Created at %v with username: %v and password: %v ", time.Now(), createUser.Username, string(passwordBytes))
	w.WriteHeader(http.StatusOK)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusBadRequest)
		return
	}
	userData := &User{}
	utils.ReadRequestBody(r, userData)

	db := config.InitializeDB()
	var user User
	err := db.Where("username = ?", userData.Username).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	passwordBytes, err := base64.StdEncoding.DecodeString(user.Passwordhash)
	if err != nil {
		http.Error(w, "Invalid password format", http.StatusBadRequest)
		return
	}
	if userData.Passwordhash != string(passwordBytes) {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		log.Println("User Unauthorized")
		return
	}
	// JWT Token Generation

	var secretKey = []byte("mysecretstring")
	expiryTime := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"expiryTime": expiryTime,
		"username":   user.Username,
		"pincode":    user.Pincode,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}
	// w.Header().Set("Authorization", "Bearer "+signedToken)

	w.WriteHeader(http.StatusOK)
	log.Println("User Authorised")

	type res struct {
		Token      string `json:"token"`
		ExpiryTime int64  `json:"expiryTime"`
	}

	t := &res{Token: signedToken, ExpiryTime: expiryTime}

	if out, err := json.MarshalIndent(t, "", " "); err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}

}
