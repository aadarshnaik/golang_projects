package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/utils"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(user *models.User, db *gorm.DB) error {
	if userExists(db, user) {
		log.Println("User already exists")
	} else if user.Username == "" || user.Passwordhash == "" || user.Pincode == 0 {
		return fmt.Errorf("some necessary field missing")
	}
	passwordBytes := user.Passwordhash
	salt := utils.GenerateSalt(8)
	passwordBytes = passwordBytes + string(salt)
	password := []byte(passwordBytes)
	log.Println("PasswordWithSalt: ", passwordBytes)

	// encodedPassword := base64.StdEncoding.EncodeToString(passwordBytes)
	encodedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	user.Salt = string(salt)
	user.Passwordhash = string(encodedPassword)
	err = db.Create(user).Error
	if err != nil {
		log.Println("Error creating user:", err)
		return fmt.Errorf("error creating user")
	}
	log.Printf("New User Created at %v with username: %v and password: %v ", time.Now(), user.Username, string(passwordBytes))

	return nil
}
