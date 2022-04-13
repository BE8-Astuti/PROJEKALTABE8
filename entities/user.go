package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Address      string
	Phone        string
	Books        []Book        `gorm:"foreignKey:user_id"`
	Rented_BookS []Rented_Book `gorm:"foreignKey:user_id"`
}
