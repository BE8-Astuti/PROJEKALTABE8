package entities

import (
	"gorm.io/gorm"
)

type Rented_Book struct {
	gorm.Model
	User_id      uint
	Book_id      uint
	Judul        string
	Pengarang    string
	Tahun_terbit uint
}
