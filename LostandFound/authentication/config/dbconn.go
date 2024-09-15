package config

import (
	"fmt"
	"log"
	"os"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {

	// dsn := fmt.Sprintf("%s:%s@/lostandfound?charset=utf8&parseTime=True&loc=Local", username, password)
	// dsn := "root:password@/lostandfound?charset=utf8&parseTime=True&loc=Local"

	//For Docker
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbhostname := os.Getenv("DBHOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/lostandfound?charset=utf8&parseTime=True&loc=Local", username, password, dbhostname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("connection to Db couldn't be established !")
		os.Exit(1)
	}
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalln("Failed to get sql.DB from gorm.DB:", err)
	// }
	// // Close the connection when you're done
	// defer sqlDB.Close()
	return db
}

func MigrateDB() error {
	db := InitializeDB()
	db.AutoMigrate(&models.User{})
	return nil
}
