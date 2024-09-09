package service

import (
	"fmt"
	"log"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ValidateCredentials(db *gorm.DB, user *models.User, user_from_db *models.User) bool {

	password := user_from_db.Passwordhash
	passwordBytes := []byte(password) //This is becrypted password hash
	saved_salt := user_from_db.Salt
	// base64.StdEncoding.DecodeString(password_hash)
	userpass := user.Passwordhash + saved_salt
	err := bcrypt.CompareHashAndPassword(passwordBytes, []byte(userpass))
	if err != nil {
		log.Println("Password does not match!")
		return false
	} else {
		log.Println("Password match!")
		return true
	}
}

func GenJWT(user *models.User, expiry int64, secretkey []byte) string {
	claims := jwt.MapClaims{
		"expiryTime": expiry,
		"username":   user.Username,
		"pincode":    user.Pincode,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretkey)
	if err != nil {
		log.Println("Error generating JWT token:", err)
		return ""
	}
	return signedToken
}

func ValidateToken(jwttoken string, secret string, req_username string) bool {

	token, err := jwt.Parse(jwttoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("Error parsing JWT token:", err)
		return false
	}
	if !token.Valid {
		log.Println("Token is invalid")
		return false
	}
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	// pincode := claims["pincode"].(int)
	expiryTime := claims["expiryTime"].(float64)

	// log.Printf("time.Now().Unix(): %T : %v ", time.Now().Unix(), time.Now().Unix())

	if time.Now().Unix() > int64(expiryTime) {
		fmt.Println("Token has expired")
		return false
	}

	if req_username != username {
		log.Println("Username mismatch")
		return false
	}
	// log.Println("Username: ", username)
	log.Println("Token is valid")
	if username != "" {
		return true
	} else {
		return false
	}
}
