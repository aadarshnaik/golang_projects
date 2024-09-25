package service

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(user *models.User, db *gorm.DB) error {

	var wg sync.WaitGroup
	ch := make(chan bool)
	wg.Add(1)
	existingUser := &models.User{}
	go func() {
		err := db.Where("username = ?", user.Username).First(existingUser).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ch <- true
				log.Printf("User with username '%s' not found.", user.Username)
			} else {
				ch <- false
				log.Println("User already exists")
			}
		}
		wg.Done()
	}()

	passwordBytes := user.Passwordhash
	salt := utils.GenerateSalt(8)
	passwordBytes = passwordBytes + string(salt)
	password := []byte(passwordBytes)
	encodedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	user.Salt = string(salt)
	user.Passwordhash = string(encodedPassword)
	if <-ch {
		err = db.Create(user).Error
		if err != nil {
			log.Println("Error creating user:", err)
			return fmt.Errorf("error creating user")
		}
	}
	log.Printf("New User Created at %v with username: %v and password: %v ", time.Now(), user.Username, string(passwordBytes))
	wg.Wait()
	return nil
}
