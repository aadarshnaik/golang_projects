package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/config"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"gorm.io/gorm"
)

func userExists(db *gorm.DB, user *models.User) bool {
	err := db.Where("username = ?", user.Username).First(user).Error
	// log.Println("err: ", err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("User with username '%s' not found.", user.Username)
			return false
		}
	}
	return true
}

func CreateUser(user *models.User) error {
	// createUser := &models.User{}
	db := config.InitializeDB()
	db.AutoMigrate(&models.User{})
	// log.Println(userExists(user))
	if userExists(db, user) {
		log.Println("User already exists")
		return fmt.Errorf("user with the same username or pincode already exists")
	} else if user.Username == "" || user.Passwordhash == "" || user.Pincode == 0 {
		return fmt.Errorf("some necessary field missing")
	}
	passwordBytes := []byte(user.Passwordhash)
	encodedPassword := base64.StdEncoding.EncodeToString(passwordBytes)
	user.Passwordhash = encodedPassword
	err := db.Create(user).Error
	if err != nil {
		log.Println("Error creating user:", err)
		return fmt.Errorf("error creating user")
	} else {
		log.Printf("New User Created at %v with username: %v and password: %v ", time.Now(), user.Username, string(passwordBytes))
	}

	return nil
}
