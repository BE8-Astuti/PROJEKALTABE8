package datastore

import (
	"be8/gorm/entities"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func AddBook(db *gorm.DB, newBook entities.Book) error {
	result := db.Create(&newBook)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func AddBookRet(db *gorm.DB, newBook entities.Book) error {
	result := db.Create(&newBook)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func GetBook(db *gorm.DB) []entities.Book {
	var books []entities.Book
	tx := db.Find(&books)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return books
}
func UpdateListBook(db *gorm.DB, id uint, judul string) error {
	result := db.Model(&entities.User{}).Where("id = ?", id).Update("judul", judul)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}

func GetBookbyid(db *gorm.DB, id uint) []entities.Book {
	var books []entities.Book
	tx := db.Find(&books, "ID= ?", id)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return books
}

func UpdateBookJudul(db *gorm.DB, id uint, judul string) error {
	result := db.Model(&entities.Book{}).Where("id = ?", id).Update("judul", judul)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}

func UpdateBookPengarang(db *gorm.DB, id uint, pengarang string) error {
	result := db.Model(&entities.User{}).Where("id = ?", id).Update("pengarang", pengarang)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}
func DeleteAvBook(db *gorm.DB, id uint) error {
	var book []entities.Book
	result := db.Delete(&book, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}
