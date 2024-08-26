package models

import (
	"time"

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
