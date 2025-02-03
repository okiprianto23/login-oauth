package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex"`
	Password     string
	RefreshToken string `gorm:"type:text"`
}
