package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string `gorm:"unique" json:"phone" form:"phone"`
	Address      string
	Phone        string
	Rented_BookS []Rented_Book `gorm:"foreignKey:user_id"`
}
