package datastore

import (
	"be8/gorm/entities"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateBorrow(db *gorm.DB, book_id entities.Rented_Book) error {
	result := db.Create(&book_id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("borrow failed")
	}
	return nil
}

func GetRentedBook(db *gorm.DB) []entities.Rented_Book {
	var books []entities.Rented_Book
	tx := db.Find(&books)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return books
}

func GetRentedBookbyUserid(db *gorm.DB, user_id uint) []entities.Rented_Book {
	var books []entities.Rented_Book
	tx := db.Find(&books, "user_id= ?", user_id)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return books
}

func GetRentedBookbyBookid(db *gorm.DB, book_id uint) []entities.Rented_Book {
	var books []entities.Rented_Book
	tx := db.Find(&books, "book_id= ?", book_id)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return books
}

func UpdateRentedBook(db *gorm.DB, book_id, user_id uint) error {
	result := db.Model(&entities.Rented_Book{}).Where("book_id = ?", book_id).Update("user_id", user_id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}
func UpdatelendingBook(db *gorm.DB, book_id, user_id uint) error {
	result := db.Model(&entities.Rented_Book{}).Where("book_id = ?", book_id).Update("user_id", user_id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}
