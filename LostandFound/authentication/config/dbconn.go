package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbhostname := os.Getenv("DBHOST")
	// dsn := fmt.Sprintf("%s:%s@/lostandfound?charset=utf8&parseTime=True&loc=Local", username, password)

	//For Docker
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
