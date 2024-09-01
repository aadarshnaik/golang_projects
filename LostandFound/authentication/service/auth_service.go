package service

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func ValidateCredentials(db *gorm.DB, user *models.User, user_from_db *models.User) bool {

	passwordBytes, err := base64.StdEncoding.DecodeString(user_from_db.Passwordhash)
	// base64.StdEncoding.DecodeString(password_hash)
	if err != nil {
		log.Println("Error decoding password hash:", err)
		return false
	}
	if user.Passwordhash != string(passwordBytes) {
		log.Println("User Unauthorized")
		return false
	} else {
		log.Println("User Authorized")
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
