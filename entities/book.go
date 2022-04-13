package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	User_id      uint
	Judul        string
	Pengarang    string
	Tahun_terbit uint
	Rented_Books []Rented_Book `gorm:"foreignKey:book_id"`
}
