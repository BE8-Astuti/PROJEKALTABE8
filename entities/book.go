package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul        string
	Pengarang    string
	Tahun_terbit string
	Rented_Books []Rented_Book `gorm:"foreignKey:book_id"`
}
