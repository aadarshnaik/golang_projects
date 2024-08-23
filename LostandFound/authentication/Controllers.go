package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	Id           int       `json:"id"`
	Username     string    `json:"username"`
	Passwordhash string    `json:"passwordhash"`
	CreatedAt    time.Time `json:"createdAt"`
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type. Expected application/json", http.StatusBadRequest)
		return
	}

	createUser := &user{}
	ReadRequestBody(r, createUser)

	passwordBytes := []byte(createUser.Passwordhash)
	encodedPassword := base64.StdEncoding.EncodeToString(passwordBytes)
	createUser.Passwordhash = encodedPassword

	db := initializeDB("root", "password")
	db.AutoMigrate(&user{})
	err := db.Create(createUser).Error
	if err != nil {
		fmt.Println("Error creating user:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	log.Printf("New User Created at %v with username: %v and password: %v ", time.Now(), createUser.Username, string(passwordBytes))

}
func ReadRequestBody(r *http.Request, a interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, a)
	if err != nil {
		return
	}
}

func initializeDB(username string, password string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@/lostandfound?charset=utf8&parseTime=True&loc=Local", username, password)
	// dsn := "root:Aadarsh98@/lostandfound?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
