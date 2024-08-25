package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	username := "root"
	password := "password"
	dsn := fmt.Sprintf("%s:%s@/lostandfound?charset=utf8&parseTime=True&loc=Local", username, password)
	// dsn := "root:Aadarsh98@/lostandfound?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
